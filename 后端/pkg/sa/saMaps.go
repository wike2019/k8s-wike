package sa

import (
	"fmt"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/common"
	"k8sapi/pkg/helper"
	"sort"
	"sync"
)

type MapStruct struct {
	data sync.Map         // [ns string] []*v1.Service
	Helper *helper.Helper `inject:"-"` //帮助函数 用于分页
}


//获取单个Service
func(this *MapStruct) Get(ns string,name string) *corev1.ServiceAccount{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*corev1.ServiceAccount){
			if item.Name==name{
				return  item
			}
		}
	}
	return nil
}
func(this *MapStruct) Add(item *corev1.ServiceAccount){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*corev1.ServiceAccount),item)
		this.data.Store(item.Namespace,list)
		common.AddAutoComplete(fmt.Sprintf("资源类型:%s 命名空间:%s 资源名称: %s","ServiceAccount",item.Namespace,item.Name),item.Name,0)
	}else{
		this.data.Store(item.Namespace,[]*corev1.ServiceAccount{item})
	}
}
func(this *MapStruct) Update(item *corev1.ServiceAccount) error {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,listItem:=range list.([]*corev1.ServiceAccount){
			if listItem.Name==item.Name{
				list.([]*corev1.ServiceAccount)[i]=item
			}
		}
		return nil
	}
	return fmt.Errorf("service-%s not found",item.Name)
}
func(this *MapStruct) Delete(item *corev1.ServiceAccount){
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,listItem:=range list.([]*corev1.ServiceAccount){
			if listItem.Name==item.Name{
				newList:= append(list.([]*corev1.ServiceAccount)[:i], list.([]*corev1.ServiceAccount)[i+1:]...)
				this.data.Store(item.Namespace,newList)
				common.DeleteAutoComplete(
					fmt.Sprintf("资源类型:%s 命名空间:%s 资源名称: %s","ServiceAccount",item.Namespace,item.Name))
				break
			}
		}
	}
}
func(this *MapStruct) ListAll(ns string)[]*ListModel{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*corev1.ServiceAccount)
		sort.Sort(CoreV1ServiceAccount(newList))//  按时间倒排序
		ret:=make([]*ListModel,len(newList))
		for i,item:=range newList{
			ret[i]=&ListModel{
				Name:item.Name,
				CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
				Namespace:item.Namespace,
				Secrets:item.Secrets,
			}
		}
		return ret
	}
	return []*ListModel{} //返回空列表
}

func(this *MapStruct) PageDeps(ns string,page int) *helper.ItemsPage {
	List:=this.ListAll(ns)
	Covert:=make([]interface{},len(List))
	for i,dep:=range List{
		Covert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		10, //尺寸默认5
		Covert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":ns, //命名空间
	})
}
package Sa

import (
	"fmt"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/Common"
	"sort"
	"sync"
)

type MapStruct struct {
	data sync.Map   // [ns string] []*v1.Service
	Helper *Common.Helper `inject:"-"`//帮助函数 用于分页
}


//获取单个Service
func(this *MapStruct) Get(ns string,name string) *Model{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*corev1.ServiceAccount){
			if item.Name==name{
				return &Model{
					Name:item.Name,
					CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
					Namespace:item.Namespace,
					Secrets:item.Secrets,
					ImagePullSecrets:item.ImagePullSecrets,
					Labels:item.Labels,
					Annotations:item.Annotations,
				}
			}
		}
	}
	return nil
}
func(this *MapStruct) Add(sa *corev1.ServiceAccount){
	if list,ok:=this.data.Load(sa.Namespace);ok{
		list=append(list.([]*corev1.ServiceAccount),sa)
		this.data.Store(sa.Namespace,list)
	}else{
		this.data.Store(sa.Namespace,[]*corev1.ServiceAccount{sa})
	}
}
func(this *MapStruct) Update(sa *corev1.ServiceAccount) error {
	if list,ok:=this.data.Load(sa.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.ServiceAccount){
			if range_pod.Name==sa.Name{
				list.([]*corev1.ServiceAccount)[i]=sa
			}
		}
		return nil
	}
	return fmt.Errorf("service-%s not found",sa.Name)
}
func(this *MapStruct) Delete(sa *corev1.ServiceAccount){
	if list,ok:=this.data.Load(sa.Namespace);ok{
		for i,range_svc:=range list.([]*corev1.ServiceAccount){
			if range_svc.Name==sa.Name{
				newList:= append(list.([]*corev1.ServiceAccount)[:i], list.([]*corev1.ServiceAccount)[i+1:]...)
				this.data.Store(sa.Namespace,newList)
				break
			}
		}
	}
}
func(this *MapStruct) ListAll(ns string)[]*Model{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*corev1.ServiceAccount)
		sort.Sort(CoreV1ServiceAccount(newList))//  按时间倒排序
		ret:=make([]*Model,len(newList))
		for i,item:=range newList{
			ret[i]=&Model{
				Name:item.Name,
				CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
				Namespace:item.Namespace,
				Secrets:item.Secrets,
				ImagePullSecrets:item.ImagePullSecrets,
				Labels:item.Labels,
				Annotations:item.Annotations,
			}
		}
		return ret
	}
	return []*Model{} //返回空列表
}

func(this *MapStruct) PageDeps(ns string,page int) *Common.ItemsPage{
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
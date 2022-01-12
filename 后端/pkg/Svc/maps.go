package Svc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/Common"
	"sort"
	"sync"
)

type ServiceMapStruct struct {
	data sync.Map   // [ns string] []*v1.Service
	Helper *Common.Helper `inject:"-"`//帮助函数 用于分页
}


//获取单个Service
func(this *ServiceMapStruct) Get(ns string,name string) *corev1.Service{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*corev1.Service){
			if item.Name==name{
				return item
			}
		}
	}
	return nil
}
func(this *ServiceMapStruct) Add(svc *corev1.Service){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		list=append(list.([]*corev1.Service),svc)
		this.data.Store(svc.Namespace,list)
	}else{
		this.data.Store(svc.Namespace,[]*corev1.Service{svc})
	}
}
func(this *ServiceMapStruct) Update(svc *corev1.Service) error {
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.Service){
			if range_pod.Name==svc.Name{
				list.([]*corev1.Service)[i]=svc
			}
		}
		return nil
	}
	return fmt.Errorf("service-%s not found",svc.Name)
}
func(this *ServiceMapStruct) Delete(svc *corev1.Service){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_svc:=range list.([]*corev1.Service){
			if range_svc.Name==svc.Name{
				newList:= append(list.([]*corev1.Service)[:i], list.([]*corev1.Service)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				break
			}
		}
	}
}
func(this *ServiceMapStruct) ListAll(ns string)[]*ServiceModel{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*corev1.Service)
		sort.Sort(CoreV1Service(newList))//  按时间倒排序
		ret:=make([]*ServiceModel,len(newList))
		for i,item:=range newList{
			ret[i]=&ServiceModel{
				Name:item.Name,
				CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
				NameSpace:item.Namespace,
				ServicePort: item.Spec.Ports,
			}
		}
		return ret
	}
	return []*ServiceModel{} //返回空列表
}


func(this *ServiceMapStruct) PageDeps(ns string,page int) *Common.ItemsPage{
	List:=this.ListAll(ns)
	Covert:=make([]interface{},len(List))
	for i,dep:=range List{
		Covert[i]=dep
	}
	return this.Helper.PageResource(
		page, //当前页
		10, //尺寸默认5
		Covert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":ns, //命名空间
	})
}
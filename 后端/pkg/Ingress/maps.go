package Ingress

import (
	"fmt"
	netv1 "k8s.io/api/networking/v1"
	Other "k8sapi/pkg/other"
	"sort"
	"sync"
)




type IngressMapStruct struct {
	data sync.Map   // [ns string] []*v1beta1.Ingress
}
//获取单个Ingress
func(this *IngressMapStruct) Get(ns string,name string) *netv1.Ingress{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*netv1.Ingress){
			if item.Name==name{
				return item
			}
		}
	}
	return nil
}
func(this *IngressMapStruct) Add(ingress *netv1.Ingress){
	if list,ok:=this.data.Load(ingress.Namespace);ok{
		list=append(list.([]*netv1.Ingress),ingress)
		this.data.Store(ingress.Namespace,list)
		Other.AddAutoComplete(fmt.Sprintf("资源类型：%s 命名空间：%s 资源名称 %s",ingress.Kind,ingress.Namespace,ingress.Name),ingress.Name,0)

	}else{
		this.data.Store(ingress.Namespace,[]*netv1.Ingress{ingress})
	}
}
func(this *IngressMapStruct) Update(ingress *netv1.Ingress) error {
	if list,ok:=this.data.Load(ingress.Namespace);ok{
		for i,range_pod:=range list.([]*netv1.Ingress){
			if range_pod.Name==ingress.Name{
				list.([]*netv1.Ingress)[i]=ingress
			}
		}
		return nil
	}
	return fmt.Errorf("ingress-%s not found",ingress.Name)
}
func(this *IngressMapStruct) Delete(ingress *netv1.Ingress){
	if list,ok:=this.data.Load(ingress.Namespace);ok{
		for i,range_ingress:=range list.([]*netv1.Ingress){
			if range_ingress.Name==ingress.Name{
				newList:= append(list.([]*netv1.Ingress)[:i], list.([]*netv1.Ingress)[i+1:]...)
				this.data.Store(ingress.Namespace,newList)
				Other.DeleteAutoComplete(
					fmt.Sprintf("资源类型：%s 命名空间：%s 资源名称 %s",
						range_ingress.Kind,range_ingress.Namespace,range_ingress.Name))
				break
			}
		}
	}
}

func(this *IngressMapStruct) ListAll(ns string)[]*IngressModel{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*netv1.Ingress)
		sort.Sort(Netv1Ingress(newList))//  按时间倒排序
		ret:=make([]*IngressModel,len(newList))
		for i,item:=range newList{
		    host:=""
			if len(item.Spec.Rules)>0 {
				host=item.Spec.Rules[0].Host
			}

			ret[i]=&IngressModel{
				Name:item.Name,
				CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
				NameSpace:item.Namespace,
				Host:host,
				IsHttps: len(item.Spec.TLS)>0,
				Annotations: item.Annotations,
				IsReady:len(item.Status.LoadBalancer.Ingress)>0,
				TLS: item.Spec.TLS,
				Rules: item.Spec.Rules,
			}
		}
		return ret
	}
	return []*IngressModel{} //返回空列表
}

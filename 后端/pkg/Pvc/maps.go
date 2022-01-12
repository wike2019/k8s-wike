package Pvc

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

type MapStruct struct {
	data sync.Map   // [ns string] []* *corev1.PersistentVolumeClaim
}


//获取单个Service
func(this *MapStruct) Get(ns string,name string) *corev1.PersistentVolumeClaim{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*corev1.PersistentVolumeClaim){
			if item.Name==name{
				return item
			}
		}
	}
	return nil
}
func(this *MapStruct) Add(pvc *corev1.PersistentVolumeClaim){
	if list,ok:=this.data.Load(pvc.Namespace);ok{
		list=append(list.([]*corev1.PersistentVolumeClaim),pvc)
		this.data.Store(pvc.Namespace,list)
	}else{
		this.data.Store(pvc.Namespace,[]*corev1.PersistentVolumeClaim{pvc})
	}
}
func(this *MapStruct) Update(pvc *corev1.PersistentVolumeClaim) error {
	if list,ok:=this.data.Load(pvc.Namespace);ok{
		for i,range_pod:=range list.([]*corev1.PersistentVolumeClaim){
			if range_pod.Name==pvc.Name{
				list.([]*corev1.PersistentVolumeClaim)[i]=pvc
			}
		}
		return nil
	}
	return fmt.Errorf("pvc-%s not found",pvc.Name)
}
func(this *MapStruct) Delete(pvc *corev1.PersistentVolumeClaim){
	if list,ok:=this.data.Load(pvc.Namespace);ok{
		for i,range_svc:=range list.([]*corev1.PersistentVolumeClaim){
			if range_svc.Name==pvc.Name{
				newList:= append(list.([]*corev1.PersistentVolumeClaim)[:i], list.([]*corev1.PersistentVolumeClaim)[i+1:]...)
				this.data.Store(pvc.Namespace,newList)
				break
			}
		}
	}
}
func(this *MapStruct) ListAll(ns string)[]*corev1.PersistentVolumeClaim{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*corev1.PersistentVolumeClaim)
		sort.Sort(CoreV1PersistentVolumeClaim(newList))//  按时间倒排序
		return newList
	}
	return []*corev1.PersistentVolumeClaim{} //返回空列表
}


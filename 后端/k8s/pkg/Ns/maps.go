package Ns

import (
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)

// namespace相关
type MapStruct struct {
	data sync.Map
}
func(this *MapStruct) Get(ns string) *corev1.Namespace{  //根据字符串获取命名空间
	if item,ok:=this.data.Load(ns);ok{
		return item.(*corev1.Namespace)
	}
	return nil
}
func(this *MapStruct) Add(ns *corev1.Namespace){
	this.data.Store(ns.Name,ns)  //添加命名空间
}
func(this *MapStruct) Update(ns *corev1.Namespace) {
	this.data.Store(ns.Name,ns) //更新命名空间
}
func(this *MapStruct) Delete(ns *corev1.Namespace){
	this.data.Delete(ns.Name)  //删除命名空间
}
//显示所有命名空间
func(this *MapStruct) ListAll()[]*NsModel{
	ret:=make([]*NsModel,0)
	this.data.Range(func(key, value interface{}) bool {
		ret=append(ret,&NsModel{
			Name:key.(string),
			CreateTime:value.(*corev1.Namespace).CreationTimestamp.Format("2006-01-02 15:04:05"),
			})
		return true
	})
	sort.Sort(NsModelSort(ret))
	return ret
}


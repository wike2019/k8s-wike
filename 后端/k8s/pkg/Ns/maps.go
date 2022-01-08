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
func(this *MapStruct) ListAll()[]*Model{
	ret:=make([]*Model,0)
	this.data.Range(func(key, value interface{}) bool {
		ret=append(ret,&Model{
			Name:key.(string),
			CreateTime:value.(*corev1.Namespace).CreationTimestamp.Format("2006-01-02 15:04:05"),
			})
		return true
	})
	sort.Sort(Sort(ret))
	return ret
}


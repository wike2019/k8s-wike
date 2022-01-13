package ns

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/common"
	"sort"
	"sync"
)

// namespace相关
type MapStruct struct {
	data sync.Map
}

func(this *MapStruct) Add(ns *corev1.Namespace){
	this.data.Store(ns.Name,ns)  //添加命名空间
	common.AddAutoComplete(fmt.Sprintf("资源类型：%s  资源名称 %s","Namespace",ns.Name),ns.Name,0)

}
func(this *MapStruct) Update(ns *corev1.Namespace) {
	this.data.Store(ns.Name,ns) //更新命名空间
}
func(this *MapStruct) Delete(ns *corev1.Namespace){
	this.data.Delete(ns.Name)  //删除命名空间
	common.DeleteAutoComplete(fmt.Sprintf("资源类型：%s  资源名称 %s","Namespace",ns.Name))
}
//显示所有命名空间
func(this *MapStruct) ListAll()[]*ListModel{
	ret:=make([]*ListModel,0)
	this.data.Range(func(key, value interface{}) bool {
		ret=append(ret,&ListModel{
			Name:key.(string),
			CreateTime:value.(*corev1.Namespace).CreationTimestamp.Format("2006-01-02 15:04:05"),
			})
		return true
	})
	sort.Sort(Sort(ret))//排序
	return ret
}


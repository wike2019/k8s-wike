package ns

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/common"
	"sort"
	"sync"
)

// namespace相关
type NsMapStruct struct {
	data sync.Map
}

func(this *NsMapStruct) Add(item *corev1.Namespace){
	this.data.Store(item.Name,item)  //添加命名空间
	common.AddAutoComplete(fmt.Sprintf("资源类型：%s  资源名称 %s","Namespace",item.Name),item.Name,0)

}
func(this *NsMapStruct) Update(item *corev1.Namespace) {
	this.data.Store(item.Name,item) //更新命名空间
}
func(this *NsMapStruct) Delete(item *corev1.Namespace){
	this.data.Delete(item.Name)  //删除命名空间
	common.DeleteAutoComplete(fmt.Sprintf("资源类型：%s  资源名称 %s","Namespace",item.Name))
}
//显示所有命名空间
func(this *NsMapStruct) ListAll()[]*ListModel{
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


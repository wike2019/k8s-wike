package Node

import (
	corev1 "k8s.io/api/core/v1"
	"sort"
	"sync"
)



//node map
type NodeMapStruct struct {
	data sync.Map   // [nodename string] *v1.Node   注意里面不是切片
}
func(this *NodeMapStruct) Get(name string) *corev1.Node{
	if node,ok:=this.data.Load(name);ok{
		return node.(*corev1.Node)
	}
	return nil
}
func(this *NodeMapStruct) Add(item *corev1.Node){
	//直接覆盖
	this.data.Store(item.Name,item)
}

func(this *NodeMapStruct) Update(item *corev1.Node) bool {
	this.data.Store(item.Name,item)
	return true
}
func(this *NodeMapStruct) Delete(node *corev1.Node){
	this.data.Delete(node.Name)
}
func(this *NodeMapStruct) ListAll()[]*corev1.Node{
	ret:=make([]*corev1.Node,0)
	this.data.Range(func(key, value interface{}) bool {
		ret=append(ret,value.(*corev1.Node))
		return true
	})
	sort.Sort(Nodelist(ret))//  按时间倒排序
	return ret//返回空列表
}
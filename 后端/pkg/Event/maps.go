package Event

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sync"
)




// event 事件map 相关
// EventSet 集合 用来保存事件, 只保存最新的一条
type EventMapStruct struct {
	data sync.Map   // [key string] *v1.Event
	// key=>namespace+"_"+kind+"_"+name 这里的name 不一定是pod ,这样确保唯一
}
func(this *EventMapStruct) GetMessage(ns string,kind string,name string) []string{
	key:=fmt.Sprintf("%s_%s_%s",ns,kind,name)
	result:=make([]string,0)
	if v,ok:=this.data.Load(key);ok{
		for _,value:=range v.([]*corev1.Event) {
			result=append(result,value.Message)
		}
	}
	return result
}


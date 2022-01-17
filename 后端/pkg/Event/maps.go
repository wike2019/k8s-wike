package event

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
func(this *EventMapStruct) GetMessage(ns string,kind string,name string,uid string)[]*corev1.Event{
	key:=fmt.Sprintf("%s_%s_%s_%s",ns,kind,name,uid)
	if v,ok:=this.data.Load(key);ok{
		return v.([]*corev1.Event)
	}
	return []*corev1.Event{}
}


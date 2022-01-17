package event

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)


// event 事件相关的handler
type EventHandler struct {
	EventMap *EventMapStruct  `inject:"-"`
}

func(this *EventHandler) storeData(obj interface{},isdelete bool){
	if event,ok:=obj.(*corev1.Event);ok{
		key:=fmt.Sprintf("%s_%s_%s_%s",event.Namespace,event.InvolvedObject.Kind,event.InvolvedObject.Name,event.InvolvedObject.UID)
		if !isdelete {
			if v, ok := this.EventMap.data.Load(key); ok {

				store := v.([]*corev1.Event)
				if len(store) < 10 {
					store = append(store, event)
				} else {
					store = store[1:]
					store = append(store, event)
				}
				this.EventMap.data.Store(key, store)
			} else {
				store:=make([]*corev1.Event,0)
				store = append(store, event)
				this.EventMap.data.Store(key, store)
			}
		}else{
			if v, ok := this.EventMap.data.Load(key); ok {
				store := v.([]*corev1.Event)
				for index,value:=range store {
					if value==obj {
						store = append(store[:index], store[index+1:]...)
						this.EventMap.data.Store(key, store)
						return
					}
				}
			}
		}

		}

}
func(this *EventHandler) OnAdd(obj interface{}){
	this.storeData(obj,false)
}
func(this *EventHandler) OnUpdate(oldObj, newObj interface{}){
	this.storeData(newObj,false)
}
func(this *EventHandler) OnDelete(obj interface{}){
	this.storeData(obj,true)
}



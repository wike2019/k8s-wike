package Rs

import (
	v1 "k8s.io/api/apps/v1"
)

// namespace 相关的回调handler
type Handler struct {
	Map *MapStruct `inject:"-"`
}
func(this *Handler) OnAdd(obj interface{}){
	this.Map.Add(obj.(*v1.ReplicaSet))

}
func(this *Handler) OnUpdate(oldObj, newObj interface{}){
	this.Map.Update(newObj.(*v1.ReplicaSet))

}
func(this *Handler) OnDelete(obj interface{}){
	if d,ok:=obj.(*v1.ReplicaSet);ok{
		this.Map.Delete(d)
	}

}


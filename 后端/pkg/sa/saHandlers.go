package sa

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
	"log"
)



// Service 相关handler
type Handler struct {
	Map *MapStruct  `inject:"-"`
}
func(this *Handler) OnAdd(obj interface{}){
	this.Map.Add(obj.(*corev1.ServiceAccount))
	ns:=obj.(*corev1.ServiceAccount).Namespace
	ret:=map[string]interface{}{"result":this.Map.PageDeps(ns,1),"type":"sa","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *Handler) OnUpdate(oldObj, newObj interface{}){
	err:=this.Map.Update(newObj.(*corev1.ServiceAccount))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*corev1.ServiceAccount).Namespace
	ret:=map[string]interface{}{"result":this.Map.PageDeps(ns,1),"type":"sa","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *Handler) OnDelete(obj interface{}){
	this.Map.Delete(obj.(*corev1.ServiceAccount))
	ns:=obj.(*corev1.ServiceAccount).Namespace
	ret:=map[string]interface{}{"result":this.Map.PageDeps(ns,1),"type":"sa","ns":ns}
	wscore.ClientMap.SendAll(ret)
}

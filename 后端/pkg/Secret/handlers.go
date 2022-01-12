package Secret

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
	"log"
)

//Secret相关的handler
type Handler struct {
	Map *MapStruct  `inject:"-"`
	Service *Service  `inject:"-"`
}
func(this *Handler) OnAdd(obj interface{}){
	this.Map.Add(obj.(*corev1.Secret))
	ns:=obj.(*corev1.Secret).Namespace
	ret:=map[string]interface{}{"result":this.Service.PageDeps(ns,1,10),"type":"secret","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *Handler) OnUpdate(oldObj, newObj interface{}){
	err:=this.Map.Update(newObj.(*corev1.Secret))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*corev1.Secret).Namespace
	ret:=map[string]interface{}{"result":this.Service.PageDeps(ns,1,10),"type":"secret","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *Handler) OnDelete(obj interface{}){
	this.Map.Delete(obj.(*corev1.Secret))
	ns:=obj.(*corev1.Secret).Namespace
	ret:=map[string]interface{}{"result":this.Service.PageDeps(ns,1,10),"type":"secret","ns":ns}
	wscore.ClientMap.SendAll(ret)
}



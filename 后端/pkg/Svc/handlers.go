 package Svc

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
	"log"
)



// Service 相关handler
type ServiceHandler struct {
	SvcMap *ServiceMapStruct  `inject:"-"`
}
func(this *ServiceHandler) OnAdd(obj interface{}){
	this.SvcMap.Add(obj.(*corev1.Service))
	ns:=obj.(*corev1.Service).Namespace
	ret:=map[string]interface{}{"result":this.SvcMap.PageDeps(ns,1),"type":" ","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *ServiceHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.SvcMap.Update(newObj.(*corev1.Service))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*corev1.Service).Namespace
	ret:=map[string]interface{}{"result":this.SvcMap.PageDeps(ns,1),"type":"service","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *ServiceHandler) OnDelete(obj interface{}){
	this.SvcMap.Delete(obj.(*corev1.Service))
	ns:=obj.(*corev1.Service).Namespace
	ret:=map[string]interface{}{"result":this.SvcMap.PageDeps(ns,1),"type":"service","ns":ns}
	wscore.ClientMap.SendAll(ret)
}

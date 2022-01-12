package Ingress

import (
	netv1 "k8s.io/api/networking/v1"
	"k8sapi/src/wscore"
	"log"
)

// ingress相关handler
type IngressHandler struct {
	IngressMap *IngressMapStruct  `inject:"-"`
	IngressService *IngressService `inject:"-"`
}
func(this *IngressHandler) OnAdd(obj interface{}){
	this.IngressMap.Add(obj.(*netv1.Ingress))
	ns:=obj.(*netv1.Ingress).Namespace

	ret:=map[string]interface{}{"result":this.IngressService.PageDeps(ns,1,10),"type":"ingress","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *IngressHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.IngressMap.Update(newObj.(*netv1.Ingress))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*netv1.Ingress).Namespace
	ret:=map[string]interface{}{"result":this.IngressService.PageDeps(ns,1,10),"type":"ingress","ns":ns}
	wscore.ClientMap.SendAll(ret)

}
func(this *IngressHandler) OnDelete(obj interface{}){
	this.IngressMap.Delete(obj.(*netv1.Ingress))
	ns:=obj.(*netv1.Ingress).Namespace
	ret:=map[string]interface{}{"result":this.IngressService.PageDeps(ns,1,10),"type":"ingress","ns":ns}
	wscore.ClientMap.SendAll(ret)
}



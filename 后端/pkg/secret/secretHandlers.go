package secret

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
	"log"
)

type SecretHandler struct {
	Map *SecretMapStruct`inject:"-"`
	Service *SecretService  `inject:"-"`
}
func(this *SecretHandler) OnAdd(obj interface{}){
	this.Map.Add(obj.(*corev1.Secret))
	ns:=obj.(*corev1.Secret).Namespace
	ret:=map[string]interface{}{"result":this.Service.PageDeps(ns,1,10),"type":"secret","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *SecretHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.Map.Update(newObj.(*corev1.Secret))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*corev1.Secret).Namespace
	ret:=map[string]interface{}{"result":this.Service.PageDeps(ns,1,10),"type":"secret","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *SecretHandler) OnDelete(obj interface{}){
	this.Map.Delete(obj.(*corev1.Secret))
	ns:=obj.(*corev1.Secret).Namespace
	ret:=map[string]interface{}{"result":this.Service.PageDeps(ns,1,10),"type":"secret","ns":ns}
	wscore.ClientMap.SendAll(ret)
}



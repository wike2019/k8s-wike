package Pvc

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
	"log"
)



// Service 相关handler
type Handler struct {
	PvcMap *MapStruct  `inject:"-"`
}
func(this *Handler) OnAdd(obj interface{}){
	this.PvcMap.Add(obj.(*corev1.PersistentVolumeClaim))
	ns:=obj.(*corev1.PersistentVolumeClaim).Namespace
	ret:=map[string]interface{}{"result":this.PvcMap.ListAll(ns),"type":"pvc","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *Handler) OnUpdate(oldObj, newObj interface{}){
	err:=this.PvcMap.Update(newObj.(*corev1.PersistentVolumeClaim))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*corev1.ServiceAccount).Namespace
	ret:=map[string]interface{}{"result":this.PvcMap.ListAll(ns),"type":"pvc","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *Handler) OnDelete(obj interface{}){
	this.PvcMap.Delete(obj.(*corev1.PersistentVolumeClaim))
	ns:=obj.(*corev1.PersistentVolumeClaim).Namespace
	ret:=map[string]interface{}{"result":this.PvcMap.ListAll(ns),"type":"pvc","ns":ns}
	wscore.ClientMap.SendAll(ret)
}

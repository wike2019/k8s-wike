package ConfigMap

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
)



//CofigMap相关的handler
type ConfigMapHandler struct {
	ConfigMap *ConfigMapStruct  `inject:"-"`
	ConfigMapService *ConfigMapService  `inject:"-"`
}
func(this *ConfigMapHandler) OnAdd(obj interface{}){
	this.ConfigMap.Add(obj.(*corev1.ConfigMap))
	ns:=obj.(*corev1.ConfigMap).Namespace
	ret:=map[string]interface{}{"result":this.ConfigMapService.PageDeps(ns,1,10),"type":"cm","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *ConfigMapHandler) OnUpdate(oldObj, newObj interface{}){
	//重点： 只要update返回true 才会发送 。否则不发送
	if this.ConfigMap.Update(newObj.(*corev1.ConfigMap)){
		ns:=newObj.(*corev1.ConfigMap).Namespace
		ret:=map[string]interface{}{"result":this.ConfigMapService.PageDeps(ns,1,10),"type":"cm","ns":ns}
		wscore.ClientMap.SendAll(ret)
	}
}
func(this *ConfigMapHandler) OnDelete(obj interface{}){
	this.ConfigMap.Delete(obj.(*corev1.ConfigMap))
	ns:=obj.(*corev1.ConfigMap).Namespace
	ret:=map[string]interface{}{"result":this.ConfigMapService.PageDeps(ns,1,10),"type":"cm","ns":ns}
	wscore.ClientMap.SendAll(ret)
}


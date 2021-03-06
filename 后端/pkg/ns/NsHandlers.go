package ns

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
)

// namespace 相关的回调handler
type NsHandler struct {
	Map *NsMapStruct `inject:"-"`
}
func(this *NsHandler) OnAdd(obj interface{}){
	this.Map.Add(obj.(*corev1.Namespace))
	ret:=map[string]interface{}{"result":this.Map.ListAll(),"type":"namespace"}
	wscore.ClientMap.SendAll(ret) //像ws发送数据
}
func(this *NsHandler) OnUpdate(oldObj, newObj interface{}){
	this.Map.Update(newObj.(*corev1.Namespace))
	ret:=map[string]interface{}{"result":this.Map.ListAll(),"type":"namespace"}
	wscore.ClientMap.SendAll(ret)
}
func(this *NsHandler) OnDelete(obj interface{}){
	if d,ok:=obj.(*corev1.Namespace);ok{
		this.Map.Delete(d)
	}
	ret:=map[string]interface{}{"result":this.Map.ListAll(),"type":"namespace"}
	wscore.ClientMap.SendAll(ret)
}


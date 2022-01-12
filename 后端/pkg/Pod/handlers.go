package Pod

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/wscore"
	"log"
)


// pod相关的回调handler
type PodHandler struct {
	PodMap *PodMapStruct `inject:"-"`
	PodService *PodService `inject:"-"`
}
func(this *PodHandler) OnAdd(obj interface{}){
	this.PodMap.Add(obj.(*corev1.Pod))
	namespace:=obj.(*corev1.Pod).Namespace
	//ws操作 强制更新为第一页
	ret:=map[string]interface{}{"result":this.PodService.PagePods(namespace,1,5),"type":"pod","ns":namespace}
	wscore.ClientMap.SendAll(ret)
	depName,ns,	depErr:=this.PodService.FatherReplicaSetData(namespace,obj.(*corev1.Pod).Name,"Deployment")
	if depErr!=nil{
		deppod:=map[string]interface{}{"result":this.PodService.GetDepPod(ns,depName),"type":"dep-pod","ns":namespace}
		wscore.ClientMap.SendAll(deppod)
	}
}
func(this *PodHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.PodMap.Update(newObj.(*corev1.Pod))
	if err!=nil{
		log.Println(err)
	}
	namespace:=newObj.(*corev1.Pod).Namespace
	//ws操作 强制更新为第一页
	ret:=map[string]interface{}{"result":this.PodService.PagePods(namespace,1,5),"type":"pod","ns":namespace}
	wscore.ClientMap.SendAll(ret)
	depName,ns,	depErr:=this.PodService.FatherReplicaSetData(namespace,newObj.(*corev1.Pod).Name,"Deployment")
	if depErr!=nil{
		ret:=map[string]interface{}{"result":this.PodService.GetDepPod(ns,depName),"type":"dep-pod","ns":namespace}
		wscore.ClientMap.SendAll(ret)
	}
}
func(this *PodHandler)	OnDelete(obj interface{}){
	if d,ok:=obj.(*corev1.Pod);ok{
		this.PodMap.Delete(d)
	}
	namespace:=obj.(*corev1.Pod).Namespace
	//ws操作 强制更新为第一页
	ret:=map[string]interface{}{"result":this.PodService.PagePods(namespace,1,5),"type":"pod","ns":namespace}
	wscore.ClientMap.SendAll(ret)
	depName,ns,	depErr:=this.PodService.FatherReplicaSetData(namespace,obj.(*corev1.Pod).Name,"Deployment")
	if depErr!=nil{
		ret:=map[string]interface{}{"result":this.PodService.GetDepPod(ns,depName),"type":"dep-pod","ns":namespace}
		wscore.ClientMap.SendAll(ret)
	}
}


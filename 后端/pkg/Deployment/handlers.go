package Deployment

import (
	v1 "k8s.io/api/apps/v1"
	"k8sapi/src/wscore"
	"log"
)

//处理deployment 回调的handler
type DepHandler struct {
	DepMap *DeploymentMap `inject:"-"`
	DepService *DeploymentService `inject:"-"`

}
func(this *DepHandler) OnAdd(obj interface{}){
	this.DepMap.Add(obj.(*v1.Deployment))
	namespace:=obj.(*v1.Deployment).Namespace
	//ws操作 强制更新为第一页
	ret:=map[string]interface{}{"result":this.DepService.PageDeps(namespace,1,5),"type":"dep","ns":namespace}
	wscore.ClientMap.SendAll(ret)
 }
func(this *DepHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.DepMap.Update(newObj.(*v1.Deployment))
	if err!=nil{
		log.Println(err)
	}else{
		namespace:=newObj.(*v1.Deployment).Namespace
		//ws操作 强制更新为第一页
		ret:=map[string]interface{}{"result":this.DepService.PageDeps(namespace,1,5),"type":"dep","ns":namespace}
		wscore.ClientMap.SendAll(ret)
	}
}
func(this *DepHandler)	OnDelete(obj interface{}){
	if d,ok:=obj.(*v1.Deployment);ok{
		this.DepMap.Delete(d)
		namespace:=obj.(*v1.Deployment).Namespace
		//ws操作 强制更新为第一页
		ret:=map[string]interface{}{"result":this.DepService.PageDeps(namespace,1,5),"type":"dep","ns":namespace}
		wscore.ClientMap.SendAll(ret)
	}
}

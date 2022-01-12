package Rbac

import (
	"k8s.io/api/rbac/v1"
	"k8sapi/src/wscore"
	"log"
)

type RoleHander struct {
	RoleMap *RoleMapStruct  `inject:"-"`
	RoleService *RoleService  `inject:"-"`
}
func(this *RoleHander) OnAdd(obj interface{}){
	this.RoleMap.Add(obj.(*v1.Role))
	ns:=obj.(*v1.Role).Namespace
	ret:=map[string]interface{}{"result":this.RoleService.PageDeps(ns,1,5),"type":"role","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *RoleHander) OnUpdate(oldObj, newObj interface{}){
	err:=this.RoleMap.Update(newObj.(*v1.Role))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*v1.Role).Namespace
	ret:=map[string]interface{}{"result":this.RoleService.PageDeps(ns,1,5),"type":"role","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *RoleHander) OnDelete(obj interface{}){
	this.RoleMap.Delete(obj.(*v1.Role))
	ns:=obj.(*v1.Role).Namespace
	ret:=map[string]interface{}{"result":this.RoleService.PageDeps(ns,1,5),"type":"role","ns":ns}
	wscore.ClientMap.SendAll(ret)
}

type RoleBindingHander struct {
	RoleBindingMap *RoleBindingMapStruct  `inject:"-"`
	RoleService *RoleService  `inject:"-"`
}
func(this *RoleBindingHander) OnAdd(obj interface{}){
	this.RoleBindingMap.Add(obj.(*v1.RoleBinding))
	ns:=obj.(*v1.RoleBinding).Namespace
	ret:=map[string]interface{}{"result":this.RoleService.BindPageDeps(ns,1,5),"type":"roleBinding","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *RoleBindingHander) OnUpdate(oldObj, newObj interface{}){
	err:=this.RoleBindingMap.Update(newObj.(*v1.RoleBinding))
	if err!=nil{
		log.Println(err)
		return
	}
	ns:=newObj.(*v1.RoleBinding).Namespace
	ret:=map[string]interface{}{"result":this.RoleService.BindPageDeps(ns,1,5),"type":"roleBinding","ns":ns}
	wscore.ClientMap.SendAll(ret)
}
func(this *RoleBindingHander) OnDelete(obj interface{}){
	this.RoleBindingMap.Delete(obj.(*v1.RoleBinding))
	ns:=obj.(*v1.RoleBinding).Namespace
	ret:=map[string]interface{}{"result":this.RoleService.BindPageDeps(ns,1,5),"type":"roleBinding","ns":ns}
	wscore.ClientMap.SendAll(ret)
}



type ClusterRoleHandler struct {
	ClusterRoleMap *ClusterRoleMapStruct  `inject:"-"`
	RoleService *RoleService `inject:"-"`
}
func(this *ClusterRoleHandler) OnAdd(obj interface{}){
	this.ClusterRoleMap.Add(obj.(*v1.ClusterRole))
	ret:=map[string]interface{}{"result":this.RoleService.ListClusterRoles(1,5),"type":"__ClusterRole__","ns":"__ClusterRole__"}
	wscore.ClientMap.SendAll(ret)
}
func(this *ClusterRoleHandler) OnUpdate(oldObj, newObj interface{}){
	err:=this.ClusterRoleMap.Update(newObj.(*v1.ClusterRole))
	if err!=nil{
		return
	}
	ret:=map[string]interface{}{"result":this.RoleService.ListClusterRoles(1,5),"type":"__ClusterRole__","ns":"__ClusterRole__"}
	wscore.ClientMap.SendAll(ret)

}
func(this *ClusterRoleHandler) OnDelete(obj interface{}){
	this.ClusterRoleMap.Delete(obj.(*v1.ClusterRole))
	ret:=map[string]interface{}{"result":this.RoleService.ListClusterRoles(1,5),"type":"__ClusterRole__","ns":"__ClusterRole__"}
	wscore.ClientMap.SendAll(ret)
}



type ClusterRoleBindingHander struct {
	ClusterRoleBindingMap *ClusterRoleBindingMapStruct  `inject:"-"`
	RoleService *RoleService  `inject:"-"`
}
func(this *ClusterRoleBindingHander) OnAdd(obj interface{}){
	this.ClusterRoleBindingMap.Add(obj.(*v1.ClusterRoleBinding))
	ret:=map[string]interface{}{"result":this.RoleService.ListClusterRoleBindings(1,5),"type":"clusterRoleBinding","ns":"__clusterRoleBinding__"}
	wscore.ClientMap.SendAll(ret)

}
func(this *ClusterRoleBindingHander) OnUpdate(oldObj, newObj interface{}){
	this.ClusterRoleBindingMap.Update(newObj.(*v1.ClusterRoleBinding))
	ret:=map[string]interface{}{"result":this.RoleService.ListClusterRoleBindings(1,5),"type":"clusterRoleBinding","ns":"__clusterRoleBinding__"}
	wscore.ClientMap.SendAll(ret)
}
func(this *ClusterRoleBindingHander) OnDelete(obj interface{}){
	this.ClusterRoleBindingMap.Delete(obj.(*v1.ClusterRoleBinding))
	ret:=map[string]interface{}{"result":this.RoleService.ListClusterRoleBindings(1,5),"type":"clusterRoleBinding","ns":"__clusterRoleBinding__"}
	wscore.ClientMap.SendAll(ret)
}
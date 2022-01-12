package Rbac

import (
	"fmt"
	"k8s.io/api/rbac/v1"
	"sort"
	"sync"
)

type RoleMapStruct struct {
	data sync.Map   // [ns string] []*v1.Role
}
func(this *RoleMapStruct) Get(ns string,name string) *v1.Role{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*v1.Role){
			if item.Name==name{
				return item
			}
		}
	}
	return nil
}
func(this *RoleMapStruct) Add(item *v1.Role){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*v1.Role),item)
		this.data.Store(item.Namespace,list)
	}else{
		this.data.Store(item.Namespace,[]*v1.Role{item})
	}
}
func(this *RoleMapStruct) Update(item *v1.Role) error {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,range_item:=range list.([]*v1.Role){
			if range_item.Name==item.Name{
				list.([]*v1.Role)[i]=item
			}
		}
		return nil
	}
	return fmt.Errorf("Role-%s not found",item.Name)
}
func(this *RoleMapStruct) Delete(svc *v1.Role){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_item:=range list.([]*v1.Role){
			if range_item.Name==svc.Name{
				newList:= append(list.([]*v1.Role)[:i], list.([]*v1.Role)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				break
			}
		}
	}
}
func(this *RoleMapStruct) ListAll(ns string)[]*v1.Role{

	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*v1.Role)
		sort.Sort(V1Role(newList))//  按时间倒排序
		return newList
	}
	return []*v1.Role{} //返回空列表
}


type RoleBindingMapStruct struct {
data sync.Map   // [ns string] []*v1.RoleBinding
}
func(this *RoleBindingMapStruct) Get(ns string,name string) *v1.RoleBinding{
	if items,ok:=this.data.Load(ns);ok{
		for _,item:=range items.([]*v1.RoleBinding){
			if item.Name==name{
				return item
			}
		}
	}
	return nil
}
func(this *RoleBindingMapStruct) Add(item *v1.RoleBinding){
	if list,ok:=this.data.Load(item.Namespace);ok{
		list=append(list.([]*v1.RoleBinding),item)
		this.data.Store(item.Namespace,list)
	}else{
		this.data.Store(item.Namespace,[]*v1.RoleBinding{item})
	}
}
func(this *RoleBindingMapStruct) Update(item *v1.RoleBinding) error {
	if list,ok:=this.data.Load(item.Namespace);ok{
		for i,range_item:=range list.([]*v1.RoleBinding){
			if range_item.Name==item.Name{
				list.([]*v1.RoleBinding)[i]=item
			}
		}
		return nil
	}
	return fmt.Errorf("Role-%s not found",item.Name)
}
func(this *RoleBindingMapStruct) Delete(svc *v1.RoleBinding){
	if list,ok:=this.data.Load(svc.Namespace);ok{
		for i,range_item:=range list.([]*v1.RoleBinding){
			if range_item.Name==svc.Name{
				newList:= append(list.([]*v1.RoleBinding)[:i], list.([]*v1.RoleBinding)[i+1:]...)
				this.data.Store(svc.Namespace,newList)
				break
			}
		}
	}
}
func(this *RoleBindingMapStruct) ListAll(ns string)[]*v1.RoleBinding{
	if list,ok:=this.data.Load(ns);ok{
		newList:=list.([]*v1.RoleBinding)
		sort.Sort(V1RoleBinding(newList))//  按时间倒排序
		return newList
	}
	return []*v1.RoleBinding{} //返回空列表
}



type ClusterRoleMapStruct struct {
	data sync.Map   // [name string] *v1.ClusterRole      之前的Role是 [ns name] []*v1.Role
	//因此下面的方法和Role是不一样的
}
func(this *ClusterRoleMapStruct) Get(name string) *v1.ClusterRole{
	if item,ok:=this.data.Load(name);ok{
		return item.(*v1.ClusterRole)
	}
	return nil
}
func(this *ClusterRoleMapStruct) Add(item *v1.ClusterRole){
	this.data.Store(item.Name,item)
}
func(this *ClusterRoleMapStruct) Update(item *v1.ClusterRole) error {
	this.data.Store(item.Name,item)
	return nil
}
func(this *ClusterRoleMapStruct) Delete(svc *v1.ClusterRole){
	this.data.Delete(svc.Name)

}
//这里不需要填写ns参数
func(this *ClusterRoleMapStruct) ListAll()[]*v1.ClusterRole{
	list:=[]*v1.ClusterRole{}
	this.data.Range(func(key, value interface{}) bool {
		list=append(list,value.(*v1.ClusterRole))
		return true
	})
	sort.Sort(V1ClusterRole(list))
	return list
}



type ClusterRoleBindingMapStruct struct {
	data sync.Map   // [name string] *v1.ClusterRoleBinding
}
func(this *ClusterRoleBindingMapStruct) Get(name string) *v1.ClusterRoleBinding{
	if v,ok:=this.data.Load(name);ok{
		return v.(*v1.ClusterRoleBinding)
	}
	return nil

}
func(this *ClusterRoleBindingMapStruct) Add(item *v1.ClusterRoleBinding){

	this.data.Store(item.Name,item)

}
func(this *ClusterRoleBindingMapStruct) Update(item *v1.ClusterRoleBinding) error {
	this.data.Store(item.Name,item)
	return nil
}
func(this *ClusterRoleBindingMapStruct) Delete(svc *v1.ClusterRoleBinding){
	this.data.Delete(svc.Name)
}
func(this *ClusterRoleBindingMapStruct) ListAll()[]*v1.ClusterRoleBinding{
	list:=[]*v1.ClusterRoleBinding{}
	this.data.Range(func(key, value interface{}) bool {
		list=append(list,value.(*v1.ClusterRoleBinding))
		return true
	})
	sort.Sort(V1ClusterRoleBinding(list))
	return list
}

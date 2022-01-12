package Rbac

import 	"k8s.io/api/rbac/v1"
type V1Role []*v1.Role
func(this V1Role) Len() int{
	return len(this)
}
func(this V1Role) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1Role) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}


type V1RoleBinding []*v1.RoleBinding
func(this V1RoleBinding) Len() int{
	return len(this)
}
func(this V1RoleBinding) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1RoleBinding) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}

type V1ClusterRole []*v1.ClusterRole
func(this V1ClusterRole) Len() int{
	return len(this)
}
func(this V1ClusterRole) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1ClusterRole) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}


type V1ClusterRoleBinding []*v1.ClusterRoleBinding
func(this V1ClusterRoleBinding) Len() int{
	return len(this)
}
func(this V1ClusterRoleBinding) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this V1ClusterRoleBinding) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}
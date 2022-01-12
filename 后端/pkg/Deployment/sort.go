package Deployment

import (
	v1 "k8s.io/api/apps/v1"
)
//实现dep按照时间排序
type V1Deployment []*v1.Deployment //排序
func(this V1Deployment) Len() int{
	return len(this)  //实现排序接口
}
func(this V1Deployment) Less(i, j int) bool{
	//根据时间排序    正排序  //实现排序接口
	return this[i].CreationTimestamp.Time.Before(this[j].CreationTimestamp.Time)
}
func(this V1Deployment) Swap(i, j int){
	//实现排序接口
	this[i],this[j]=this[j],this[i]
}


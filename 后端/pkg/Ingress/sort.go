package Ingress

import (
	netv1 "k8s.io/api/networking/v1"
)



type Netv1Ingress []*netv1.Ingress
func(this Netv1Ingress) Len() int{
	return len(this)
}
func(this Netv1Ingress) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this Netv1Ingress) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}


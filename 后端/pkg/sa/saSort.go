package sa

import corev1 "k8s.io/api/core/v1"



type CoreV1ServiceAccount []*corev1.ServiceAccount
func(this CoreV1ServiceAccount) Len() int{
	return len(this)
}
func(this CoreV1ServiceAccount) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this CoreV1ServiceAccount) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}
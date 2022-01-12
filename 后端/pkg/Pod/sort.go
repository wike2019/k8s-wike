package Pod

import corev1 "k8s.io/api/core/v1"

//实现pod按照时间排序
type CoreV1Pods []*corev1.Pod //排序
func(this CoreV1Pods) Len() int{
	return len(this)  //实现排序接口
}
func(this CoreV1Pods) Less(i, j int) bool{
	//根据时间排序    正排序  //实现排序接口
	return this[i].CreationTimestamp.Time.Before(this[j].CreationTimestamp.Time)
}
func(this CoreV1Pods) Swap(i, j int){
	//实现排序接口
	this[i],this[j]=this[j],this[i]
}

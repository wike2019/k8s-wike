package Node

import corev1 "k8s.io/api/core/v1"

type Nodelist []*corev1.Node
func(this Nodelist) Len() int{
	return len(this)
}
func(this Nodelist) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this Nodelist) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
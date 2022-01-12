package Secret

import corev1 "k8s.io/api/core/v1"

type CoreV1Secret []*corev1.Secret
func(this CoreV1Secret) Len() int{
	return len(this)
}
func(this CoreV1Secret) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func(this CoreV1Secret) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

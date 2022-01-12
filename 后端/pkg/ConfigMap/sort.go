package ConfigMap





//ConfigMapMap

type CoreV1ConfigMap []*cm
func(this CoreV1ConfigMap) Len() int{
	return len(this)
}
func(this CoreV1ConfigMap) Less(i, j int) bool{
	//根据时间排序    倒排序
	return this[i].cmdata.CreationTimestamp.Time.After(this[j].cmdata.CreationTimestamp.Time)
}
func(this CoreV1ConfigMap) Swap(i, j int){
	this[i],this[j]=this[j],this[i]
}



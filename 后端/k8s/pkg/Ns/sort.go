package Ns


//实现命名空间排序
type NsModelSort []*NsModel //排序
func(this NsModelSort) Len() int{
	return len(this)  //实现排序接口
}
func(this NsModelSort) Less(i, j int) bool{
	//根据时间排序    正排序  //实现排序接口
	return this[i].Name<this[j].Name
}
func(this NsModelSort) Swap(i, j int){
	//实现排序接口
	this[i],this[j]=this[j],this[i]
}



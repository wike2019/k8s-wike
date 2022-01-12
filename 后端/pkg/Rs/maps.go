package Rs

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"sync"
)

// namespace相关
type MapStruct struct {
	data sync.Map
}

func(this *MapStruct) Add(ReplicaSet *v1.ReplicaSet){

	if list,ok:=this.data.Load(ReplicaSet.Namespace);ok{
		list=append(list.([]*v1.ReplicaSet),ReplicaSet) //取出列表添加再存回去
		this.data.Store(ReplicaSet.Namespace,list)
	}else{
		this.data.Store(ReplicaSet.Namespace,[]*v1.ReplicaSet{ReplicaSet}) //第一次添加
	}

}
func(this *MapStruct) Update(ReplicaSet *v1.ReplicaSet) error{
	if list,ok:=this.data.Load(ReplicaSet.Namespace);ok{
		for i,range_ReplicaSet:=range list.([]*v1.ReplicaSet){ //更新
			if range_ReplicaSet.Name==ReplicaSet.Name{
				list.([]*v1.ReplicaSet)[i]=ReplicaSet
			}
		}
		return nil
	}
	return fmt.Errorf("Pod-%s not found",ReplicaSet.Name)
}
func(this *MapStruct) Delete(ReplicaSet *v1.ReplicaSet){
	if list,ok:=this.data.Load(ReplicaSet.Namespace);ok{
		for i,range_ReplicaSet:=range list.([]*v1.ReplicaSet){
			if range_ReplicaSet.Name==ReplicaSet.Name{
				newList:= append(list.([]*v1.ReplicaSet)[:i], list.([]*v1.ReplicaSet)[i+1:]...) //删除
				this.data.Store(ReplicaSet.Namespace,newList)
				break
			}
		}
	}
}

func(this *MapStruct) Get(name string,ns string ) *v1.ReplicaSet{

	if list,ok:=this.data.Load(ns);ok{  //根据命名空间和pod名称获取pod对象
		for _,range_ReplicaSet:=range list.([]*v1.ReplicaSet){
			if range_ReplicaSet.Name==name{
				return range_ReplicaSet
			}
		}
	}
	return nil
}

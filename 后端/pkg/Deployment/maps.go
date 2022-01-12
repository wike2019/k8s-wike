package Deployment

import (
	"fmt"
	"k8s.io/api/apps/v1"
	"sort"
	"sync"
)




//对deployments的集合进行定义
type DeploymentMap struct {
	data sync.Map  // [key string] []*v1.Deployment    key=>namespace
}
//添加
func(this *DeploymentMap) Add(dep *v1.Deployment){
	if list,ok:=this.data.Load(dep.Namespace);ok{
		list=append(list.([]*v1.Deployment),dep)
		this.data.Store(dep.Namespace,list)
	}else{
		this.data.Store(dep.Namespace,[]*v1.Deployment{dep})
	}
}
//更新
func(this *DeploymentMap) Update(dep *v1.Deployment) error {
	if list,ok:=this.data.Load(dep.Namespace);ok{
		for i,range_dep:=range list.([]*v1.Deployment){
			if range_dep.Name==dep.Name{
				list.([]*v1.Deployment)[i]=dep
			}
		}
		return nil
	}
	return fmt.Errorf("deployment-%s not found",dep.Name)
}
// 删除
func(this *DeploymentMap) Delete(dep *v1.Deployment){
	if list,ok:=this.data.Load(dep.Namespace);ok{
		for i,range_dep:=range list.([]*v1.Deployment){
			if range_dep.Name==dep.Name{
				newList:= append(list.([]*v1.Deployment)[:i], list.([]*v1.Deployment)[i+1:]...)
				this.data.Store(dep.Namespace,newList)
				break
			}
		}
	}
}

func(this *DeploymentMap) ListByNS(ns string) []*v1.Deployment{  //根据命名空间获得dep列表
	if list,ok:=this.data.Load(ns);ok{  //取数据
		ret:=list.([]*v1.Deployment) //断言
		sort.Sort(V1Deployment(ret))//排序
		return ret
	}
	return nil
}

func(this *DeploymentMap) GetDeployment(ns string,depname string) *v1.Deployment{ //根据命名空间和dep名称获得dep对象
	if list,ok:=this.data.Load(ns);ok {
		for _,item:=range list.([]*v1.Deployment){
			if item.Name==depname{
				return item
			}
		}
	}
	return nil
}


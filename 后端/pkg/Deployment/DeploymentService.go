package Deployment

import (
	"github.com/gin-gonic/gin"
	"k8s.io/api/apps/v1"
	"k8sapi/pkg/Common"

)
//@Service
type DeploymentService struct {
	DepMap *DeploymentMap `inject:"-"`
	Common *CommonService `inject:"-"`
	Helper *Common.Helper `inject:"-"`//帮助函数 用于分页
}
func NewDeploymentService() *DeploymentService {
	return &DeploymentService{}
}
//内部方法判断dep状态，显示消息
func(*DeploymentService) getDeploymentCondition(dep *v1.Deployment) string{
	for _,item:=range dep.Status.Conditions{
		if string(item.Type)=="Available" && string(item.Status)!="True"{
			return item.Message
		}
		}
		return ""
}
//判断dep是否就绪
func(*DeploymentService) getDeploymentIsComplete(dep *v1.Deployment) bool{
	return  dep.Status.Replicas==dep.Status.AvailableReplicas
}
//内部方法根据ns获取dep列表
func(this *DeploymentService) listByNs(namespace string ) (ret []*Deployment){
	depList:=this.DepMap.ListByNS(namespace)
	for _,item:=range depList{ //遍历所有deployment
		// 加入副本数
		ret=append(ret,&Deployment{Name:item.Name,
			NameSpace:item.Namespace,
			Replicas:[3]int32{item.Status.Replicas,item.Status.AvailableReplicas,item.Status.UnavailableReplicas},
			Images:this.Common.GetImages(*item),
			IsComplete:this.getDeploymentIsComplete(item),
			Message:this.getDeploymentCondition(item),
			CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	return
}
//对外方法，根据分页取dep列表
func(this *DeploymentService) PageDeps(ns string,page,size int ) *Common.ItemsPage{
	depList:=this.listByNs(ns)
	depCovert:=make([]interface{},len(depList))
	for i,dep:=range depList{
		depCovert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		size, //尺寸默认5
		depCovert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":ns, //命名空间
	})
}
package Deployment

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/pkg/Common"
)

type DeploymentCtl struct {
	Helper *Common.Helper `inject:"-"`
	DepService *DeploymentService  `inject:"-"`  //首字母一定要大写
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}
func(this *DeploymentCtl) getList(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","5")
	return gin.H{
		"code":200,
		"data":this.DepService.PageDeps(ns,this.Helper.StrToInt(page,1),this.Helper.StrToInt(size,5)),
	}
}
func(this *DeploymentCtl)  Build(goft *goft.Goft){
	//路由
	goft.Handle("GET","/deployments",this.getList)
}
func(*DeploymentCtl) Name() string{
	return "DeploymentCtl"
}
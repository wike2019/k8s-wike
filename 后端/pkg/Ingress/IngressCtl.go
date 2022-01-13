package Ingress

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	netv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/helper"
)

type IngressCtl struct{
	IngressService *IngressService `inject:"-"`
	Helper *helper.Helper          `inject:"-"` //帮助函数 用于分页
	Client *kubernetes.Clientset   `inject:"-"`
}
func NewIngressCtl() *IngressCtl{
  return &IngressCtl{}
}
func(*IngressCtl)  Name() string{
	 return "IngressCtl"
}
func(this *IngressCtl) ListAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","10")
	return gin.H{
		"code":200,
		"data":this.IngressService.PageDeps(ns,this.Helper.StrToInt(page,1),this.Helper.StrToInt(size,10)),
	}
}
func(this *IngressCtl) PostIngress(c *gin.Context) goft.Json{
	postModel:=&netv1.Ingress{}
	goft.Error(c.BindJSON(postModel))
	goft.Error(this.IngressService.PostIngress(postModel))
	return gin.H{
		"code":200,
		"data":"资源创建成功",
	}
}
func(this *IngressCtl) DelIngress(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	goft.Error(this.IngressService.DelIngress(ns,name))
	return gin.H{
		"code":200,
		"data":"资源删除成功",
	}
}
func(this *IngressCtl) PutIngresses(c *gin.Context) goft.Json{
	postModel:=&netv1.Ingress{}
	err:=c.ShouldBindJSON(postModel)
	goft.Error(err)
	_,err=this.Client.NetworkingV1().Ingresses(postModel.Namespace).Update(
		c,
		postModel,
		v1.UpdateOptions{},
	)
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"修改资源成功",
	}
}

func(this *IngressCtl) GetItem(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	return gin.H{
		"code":200,
		"data":this.IngressService.GetItem(ns,name),
	}
}

func(this *IngressCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/ingress/item",this.GetItem)
	goft.Handle("GET","/ingress",this.ListAll)
	goft.Handle("POST","/ingress",this.PostIngress)
	goft.Handle("PUT","/ingress",this.PutIngresses)
	goft.Handle("DELETE","/ingress",this.DelIngress)
}

package secret

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/helper"
)

type SecretCtl struct{
	SecretService *SecretService `inject:"-"`
	Helper *helper.Helper  `inject:"-"` //帮助函数 用于分页
}
func NewSecretCtl() *SecretCtl{
  return &SecretCtl{}
}
func(*SecretCtl)  Name() string{
	 return "SecretCtl"
}
func(this *SecretCtl) listPage(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	return gin.H{
		"code":200,
		"data":this.SecretService.PageDeps(ns,this.Helper.StrToInt(page,1)),
	}
}
func(this *SecretCtl) listAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":200,
		"data":this.SecretService.All(ns),
	}
}
//
func(this *SecretCtl) update(c *gin.Context) goft.Json{
	postModel:=&corev1.Secret{}
	goft.Error(c.BindJSON(postModel))
	goft.Error(this.SecretService.UpdateSecret(postModel))
	return gin.H{
		"code":200,
		"data":"资源创建成功",
	}
}
func(this *SecretCtl) create(c *gin.Context) goft.Json{
	postModel:=&corev1.Secret{}
	goft.Error(c.BindJSON(postModel))
	goft.Error(this.SecretService.PostSecret(postModel))
	return gin.H{
		"code":200,
		"data":"资源创建成功",
	}
}
func(this *SecretCtl) del(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	goft.Error(this.SecretService.DelSecret(ns,name))
	return gin.H{
		"code":200,
		"data":"资源删除成功",
	}
}

func(this *SecretCtl) getItem(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")

	return gin.H{
		"code":200,
		"data":	this.SecretService.GetSecret(ns,name),
	}
}

func(this *SecretCtl)  Build(goft *goft.Goft){

	goft.Handle("GET","/secret",this.listAll)
	goft.Handle("GET","/secret/page",this.listPage)
	goft.Handle("POST","/secret",this.create)
	goft.Handle("DELETE","/secret",this.del)
	goft.Handle("GET","/secret/item",this.getItem)
	goft.Handle("PUT","/secret",this.update)

}

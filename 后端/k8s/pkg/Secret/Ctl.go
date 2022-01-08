package Secret

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	"k8sapi/pkg/Common"
)

type Ctl struct{
	SecretService *Service `inject:"-"`
	Helper *Common.Helper `inject:"-"`//帮助函数 用于分页
}
func NewSecretCtl() *Ctl{
  return &Ctl{}
}
func(*Ctl)  Name() string{
	 return "SecretCtl"
}
func(this *Ctl) list(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	return gin.H{
		"code":200,
		"data":this.SecretService.PageDeps(ns,this.Helper.StrToInt(page,1),this.Helper.StrToInt("10",10)),
	}
}
func(this *Ctl) listAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":200,
		"data":this.SecretService.All(ns),
	}
}
//
func(this *Ctl) update(c *gin.Context) goft.Json{
	postModel:=&corev1.Secret{}
	goft.Error(c.BindJSON(postModel))
	goft.Error(this.SecretService.UpdateSecret(postModel))
	return gin.H{
		"code":200,
		"data":"资源创建成功",
	}
}
func(this *Ctl) create(c *gin.Context) goft.Json{
	postModel:=&corev1.Secret{}
	goft.Error(c.BindJSON(postModel))
	goft.Error(this.SecretService.PostSecret(postModel))
	return gin.H{
		"code":200,
		"data":"资源创建成功",
	}
}
func(this *Ctl) del(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	goft.Error(this.SecretService.DelSecret(ns,name))
	return gin.H{
		"code":200,
		"data":"资源删除成功",
	}
}

func(this *Ctl) getItem(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")

	return gin.H{
		"code":200,
		"data":	this.SecretService.GetSecret(ns,name),
	}
}

func(this *Ctl)  Build(goft *goft.Goft){
	goft.Handle("GET","/secret/all",this.listAll)
	goft.Handle("GET","/secret",this.list)
	goft.Handle("POST","/secret",this.create)
	goft.Handle("DELETE","/secret",this.del)
	goft.Handle("GET","/GetSecret",this.getItem)
	goft.Handle("POST","/secret/update",this.update)

}

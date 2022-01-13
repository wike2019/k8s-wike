package Svc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/helper"
)

type SvcCtl struct{
	SvcMap *ServiceMapStruct     `inject:"-"`
	Helper *helper.Helper        `inject:"-"`
	Client *kubernetes.Clientset `inject:"-"`
}
func NewSvcCtl() *SvcCtl{
	return &SvcCtl{}
}
func(this *SvcCtl) ListAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	fmt.Println(ns)
	return gin.H{
		"code":200,
		"data":this.SvcMap.ListAll(ns),
	}
}
func(this *SvcCtl) ListByPage(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("page","1")
	return gin.H{
		"code":200,
		"data":this.SvcMap.PageDeps(ns,this.Helper.StrToInt(page,1)),
	}
}
func(this *SvcCtl) GetItem(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	return gin.H{
		"code":200,
		"data":this.SvcMap.Get(ns,name),
	}
}
func(this *SvcCtl) Create(c *gin.Context) goft.Json{
	svc:=&v1.Service{}
	goft.Error(c.ShouldBindJSON(svc))
	_,err:=this.Client.CoreV1().Services(svc.Namespace).Create(c,svc,metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"创建成功",
	}
}
func(this *SvcCtl) Del(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	err:=this.Client.CoreV1().Services(ns).Delete(c,name,metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"创建成功",
	}
}

func(this *SvcCtl) Update(c *gin.Context) goft.Json{
	svc:=&v1.Service{}
	goft.Error(c.ShouldBindJSON(svc))
	_,err:=this.Client.CoreV1().Services(svc.Namespace).Update(c,svc,metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"修改成功",
	}
}
func(this *SvcCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/svc/all",this.ListAll)
	goft.Handle("DELETE","/svc",this.Del)
	goft.Handle("GET","/svc",this.ListByPage)
	goft.Handle("GET","/svc/item",this.GetItem)
	goft.Handle("POST","/svc",this.Create)
	goft.Handle("POST","/svc/update",this.Update)
}
func(*SvcCtl)  Name() string{
	return "SvcCtl"
}

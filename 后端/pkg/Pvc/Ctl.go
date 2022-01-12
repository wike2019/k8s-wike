package Pvc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
)

type PvcCtl struct{
	ServiceMapStruct *MapStruct `inject:"-"`
	Client *kubernetes.Clientset `inject:"-"`
}
func NewSaCtl() *PvcCtl{
	return &PvcCtl{}
}
func(this *PvcCtl) List(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	fmt.Println(ns)
	return gin.H{
		"code":200,
		"data":this.ServiceMapStruct.ListAll(ns),
	}
}
func(this *PvcCtl) del(c *gin.Context) goft.Json{

	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *PvcCtl) create(c *gin.Context) goft.Json{

	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *PvcCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/pvc/all",this.List)
	goft.Handle("DELETE","/pvc",this.del)
	goft.Handle("POST","/pvc",this.create)
}
func(*PvcCtl)  Name() string{
	return "PvcCtl"
}

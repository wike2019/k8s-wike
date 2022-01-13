package sa

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/helper"
)

type SaCtl struct{
	MapStruct *MapStruct         `inject:"-"`
	Helper *helper.Helper        `inject:"-"`
	Client *kubernetes.Clientset `inject:"-"`
}
func NewSaCtl() *SaCtl{
	return &SaCtl{}
}
//列表
func(this *SaCtl) listAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":200,
		"data":this.MapStruct.ListAll(ns),
	}
}

//列表带分页
func(this *SaCtl) listPage(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("page","1")
	return gin.H{
		"code":200,
		"data":this.MapStruct.PageDeps(ns,this.Helper.StrToInt(page,1)),
	}
}
//具体一项
func(this *SaCtl) getItem(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	result:=this.MapStruct.Get(ns,name)
	if result==nil{
		goft.Error(fmt.Errorf("资源不存在"))
	}
	return gin.H{
		"code":200,
		"data":result,
	}
}

//删除
func(this *SaCtl) del(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	err:=this.Client.CoreV1().ServiceAccounts(ns).Delete(c,name,metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
//创建
func(this *SaCtl) create(c *gin.Context) goft.Json{
	ServiceAccount:=&corev1.ServiceAccount{} //原生的k8s ServiceAccount 对象
	goft.Error(c.ShouldBindJSON(ServiceAccount))
	_,err:=this.Client.CoreV1().ServiceAccounts(ServiceAccount.Namespace).Create(c,ServiceAccount,metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
//更新
func(this *SaCtl) update(c *gin.Context) goft.Json{
	ServiceAccount:=&corev1.ServiceAccount{} //原生的k8s ServiceAccount 对象
	goft.Error(c.ShouldBindJSON(ServiceAccount))
	fmt.Println(ServiceAccount)
	_,err:=this.Client.CoreV1().ServiceAccounts(ServiceAccount.Namespace).Update(c,ServiceAccount,metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *SaCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/sa",this.listAll)
	goft.Handle("GET","/sa/page",this.listPage)
	goft.Handle("DELETE","/sa",this.del)
	goft.Handle("POST","/sa",this.create)
	goft.Handle("GET","/sa/item",this.getItem)
	goft.Handle("PUT","/sa",this.update)
}
func(*SaCtl)  Name() string{
	return "SaCtl"
}

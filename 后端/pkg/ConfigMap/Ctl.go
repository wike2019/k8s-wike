package ConfigMap

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/Common"
)


//@restcontroller
type ConfigMapCtl struct {
	ConfigService *ConfigMapService `inject:"-"`
	Client *kubernetes.Clientset  `inject:"-"`
	Helper *Common.Helper `inject:"-"`
}
func NewConfigMapCtl() *ConfigMapCtl{
	return &ConfigMapCtl{}
}
func(*ConfigMapCtl)  Name() string{
	return "ConfigMapCtl"
}
//提交 config map
func(this *ConfigMapCtl) PostConfigmap(c *gin.Context) goft.Json{
	postModel:=&corev1.ConfigMap{}
	err:=c.ShouldBindJSON(postModel)

	goft.Error(err)
	_,err=this.Client.CoreV1().ConfigMaps(postModel.Namespace).Create(
		c,
		postModel,
		v1.CreateOptions{},
	)
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"资源创建成功",
	}
}
// 列出config map
func(this *ConfigMapCtl) ListAll(c *gin.Context) goft.Json{

	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","5")
	return gin.H{
		"code":200,
		"data":this.ConfigService.PageDeps(ns,this.Helper.StrToInt(page,1),this.Helper.StrToInt(size,10)),
	}

}
// 列出config map
func(this *ConfigMapCtl) List(c *gin.Context) goft.Json{

	ns:=c.DefaultQuery("ns","default")

	return gin.H{
		"code":200,
		"data":this.ConfigService.All(ns),
	}

}
//DELETE /configmaps?ns=xx&name=xx
func(this *ConfigMapCtl) RmCm(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	goft.Error(this.Client.CoreV1().ConfigMaps(ns).
		Delete(c,name,v1.DeleteOptions{}))
	return gin.H{
		"code":200,
		"data":"资源删除成功",
	}
}

func(this *ConfigMapCtl) GetConfigMap(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")

	return gin.H{
		"code":200,
		"data":	this.ConfigService.GetConfigMap(ns,name),
	}
}

func(this *ConfigMapCtl) PutConfigmap(c *gin.Context) goft.Json{
	postModel:=&corev1.ConfigMap{}
	err:=c.ShouldBindJSON(postModel)

	goft.Error(err)
	_,err=this.Client.CoreV1().ConfigMaps(postModel.Namespace).Update(
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
func(this *ConfigMapCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/ConfigMap",this.ListAll)
	goft.Handle("GET","/ConfigMap/all",this.List)
	goft.Handle("POST","/ConfigMap",this.PostConfigmap)
	goft.Handle("PUT","/ConfigMap",this.PutConfigmap)
	goft.Handle("DELETE","/ConfigMap",this.RmCm)
	goft.Handle("GET","/GetConfigMap",this.GetConfigMap)
}

package ns

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Ctl struct{
	Map *MapStruct `inject:"-"`
	Client *kubernetes.Clientset `inject:"-"`
}
func NewNsCtl() *Ctl {
  return &Ctl{}
}
//获取全部 不带分页
func(this *Ctl) listAll(c *gin.Context) goft.Json{
	return gin.H{"code":200, "data":this.Map.ListAll()}
}

//删除
func(this *Ctl) delete(c *gin.Context) goft.Json{
	name:=c.DefaultQuery("name","")

	if name=="" {
		goft.Error(fmt.Errorf("名称不能为空"))
	}
	err:=this.Client.CoreV1().Namespaces().Delete(c, name, metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{"code":200, "data":"删除成功"}
}
//添加
func(this *Ctl) create(c *gin.Context) goft.Json{
	obj:=&CreateModel{}
	err:=c.ShouldBindJSON(obj)
	goft.Error(err)
	if obj.Name=="" {
		goft.Error(fmt.Errorf("名称不能为空"))
	}
	namespace := &corev1.Namespace{
		TypeMeta:   metav1.TypeMeta{APIVersion: corev1.SchemeGroupVersion.String(), Kind: "Namespace"},
		ObjectMeta: metav1.ObjectMeta{Name: obj.Name},
	}
	_,err=this.Client.CoreV1().Namespaces().Create(c, namespace, metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{"code":200, "data":"创建成功"}
}

//路由
func(this *Ctl)  Build(goft *goft.Goft){
	goft.Handle("GET","/ns",this.listAll)
	goft.Handle("POST","/ns",this.create)
	goft.Handle("DELETE","/ns",this.delete)
}
func(this *Ctl)  Name() string{
	 return "NsCtl"
}
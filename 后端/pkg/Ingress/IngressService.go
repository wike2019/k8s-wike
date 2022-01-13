package Ingress

import (
	"context"
	"github.com/gin-gonic/gin"
	netv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/helper"
)
//@Service
type IngressService struct {
	IngressMapStruct *IngressMapStruct `inject:"-"`
	Helper *helper.Helper              `inject:"-"` //帮助函数 用于分页
	Client *kubernetes.Clientset       `inject:"-"`
}
func NewIngressService() *IngressService {
	return &IngressService{}
}


//对外方法，根据分页取dep列表
func(this *IngressService) PageDeps(ns string,page,size int ) *helper.ItemsPage {
	ingressList:=this.IngressMapStruct.ListAll(ns)
	ingressCovert:=make([]interface{},len(ingressList))
	for i,dep:=range ingressList{
		ingressCovert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		size, //尺寸默认5
		ingressCovert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":ns, //命名空间
	})
}

func(this *IngressService) PostIngress(post *netv1.Ingress) error{
	_,err:= this.Client.NetworkingV1().Ingresses(post.Namespace).
		Create(context.Background(),post,v1.CreateOptions{})
	return err

}
func(this *IngressService) UpdateIngress(post *netv1.Ingress) error{
	_,err:= this.Client.NetworkingV1().Ingresses(post.Namespace).
		Update(context.Background(),post,v1.UpdateOptions{})
	return err
}

func(this *IngressService) GetItem(ns string,name string ) *netv1.Ingress{
	return  this.IngressMapStruct.Get(ns,name)
}
func(this *IngressService) DelIngress(namespace string,name string) error{
	return  this.Client.NetworkingV1().Ingresses(namespace).Delete(context.Background(),name,v1.DeleteOptions{})
}


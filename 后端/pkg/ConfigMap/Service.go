package ConfigMap

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/Common"
)

//@service
type ConfigMapService struct {
	Client *kubernetes.Clientset `inject:"-"`
	ConfigMap *ConfigMapStruct  `inject:"-"`
	Helper *Common.Helper `inject:"-"`//帮助函数 用于分页
}
func NewConfigMapService() *ConfigMapService {
	return &ConfigMapService{}
}
//前台用于显示Secret列表
func(this *ConfigMapService) PageDeps(ns string,page,size int ) *Common.ItemsPage{
	ConfigMapList:=this.ConfigMap.ListAll(ns)
	ConfigMapCovert:=make([]interface{},len(ConfigMapList))
	for i,dep:=range ConfigMapList{
		ConfigMapCovert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		size, //尺寸默认5
		ConfigMapCovert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":ns, //命名空间
	})
}

func(this *ConfigMapService) All(ns string)[]*ConfigMapModel{
	return this.ConfigMap.ListAll(ns)
}

func(this *ConfigMapService) GetConfigMap(ns,name string) ConfigMapModel{
	return this.ConfigMap.GetConfigMap(ns,name)
}

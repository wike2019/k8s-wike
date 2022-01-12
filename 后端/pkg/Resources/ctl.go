package Resources

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"strings"
)

type ResourcesCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewResourcesCtl() *ResourcesCtl {
	return &ResourcesCtl{}
}
func(this *ResourcesCtl) GetGroupVersion(str string) (group,version string){
	list:=strings.Split(str,"/")
	if len(list)==1{
		return "core",list[0]
	}else if len(list)==2{
		return list[0],list[1]
	}
	panic("error GroupVersion"+str)
}
//获取所有资源
func(this *ResourcesCtl) ListResources(c *gin.Context) goft.Json{
	_,res,err:=this.Client.ServerGroupsAndResources()
	set:=make(map[string]struct{})
	goft.Error(err)
	gRes:=make([]*GroupResources,0)
	for _,r:=range res{
		group,version:=this.GetGroupVersion(r.GroupVersion)
		if _, isOk := set[group]; isOk{
			continue
		}
		gr:=&GroupResources{Group:group,Version:version,Resources:make([]*Resources,0)}
		for _,rr:=range r.APIResources{
			res:=&Resources{Name:rr.Name,Verbs:rr.Verbs}
			gr.Resources=append(gr.Resources,res)
		}
		gRes=append(gRes,gr)
		set[group]= struct{}{}
	}
	return gin.H{
		"code":200,
		"data":gRes,
	}
}
func(*ResourcesCtl) Name() string{
	return "Resources"
}
func(this *ResourcesCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/resources",this.ListResources)
}
package Pod

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8sapi/pkg/Event"
	"k8sapi/pkg/Rs"
	"k8sapi/pkg/helper"
)

//@Service
type PodService struct {
	PodMap *PodMapStruct           `inject:"-"` //pod对象存储集合
	Common *CommonService          `inject:"-"` //公共服务 例如镜像，状态等处理
	EventMap *event.EventMapStruct `inject:"-"` //事件对象存储集合
	RsMapStruct *Rs.MapStruct      `inject:"-"` //rs对象
	Helper *helper.Helper          `inject:"-"` //帮助函数 用于分页
}

func NewPodService() *PodService {
	return &PodService{}
}
//内部方法获取命名空间下所有pod列表
func(this *PodService) listByNs(ns string) []*Pod{
	podList:= this.PodMap.ListByNs(ns)  //根据命名空间取得所有该命名空间下的pod列表
	ret:=make([]*Pod,0)
	//dto
	for _,pod:=range podList{
		ret=append(ret,&Pod{
			Name:pod.Name, //pod名称
			NameSpace:pod.Namespace, //pod的命名空间
			Images:this.Common.GetImagesByPod(pod.Spec.Containers), //取得镜像
			NodeName:pod.Spec.NodeName, //取得调度的主机名称
			Phase:string(pod.Status.Phase),// 阶段
			IsReady:this.Common.PosIsReady(pod), //是否就绪
			IP:[]string{pod.Status.PodIP,pod.Status.HostIP}, //取ip
			Message:this.EventMap.GetMessage(pod.Namespace,"Pod",pod.Name,string(pod.UID)), //取事件信息
			CreateTime:pod.CreationTimestamp.Format("2006-01-02 15:04:05"),//取时间戳
		})
	}
	return  ret
}
//对外方法，根据分页取pod
func(this *PodService) PagePods(ns string,page,size int ) *helper.ItemsPage {
	pods:=this.listByNs(ns)
	readyCount:=0 //就绪的pod数量
	allCount:=0 //总数量
	podsCovert:=make([]interface{},len(pods))
	for i,pod:=range pods{
		allCount++
		podsCovert[i]=pod //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
		if pod.IsReady{
			readyCount++
		}
	}
	return this.Helper.PageResource(
		page, //当前页
		size, //尺寸默认5
		podsCovert,//全部数据集合，会自动筛选出分页里面的内容
		).SetExt(gin.H{
		"ReadyNum":readyCount,//额外辅助数据 就绪数量
		"AllNum":allCount, //额外辅助数据 所有数量
		"ns":ns, //命名空间
		})
}
func(this *PodService) FatherReplicaSetData(ns,podName string,kind string)  (string,string,error)  {
	pod:=this.PodMap.Get(ns,podName)
	ReferencesList:=pod.GetOwnerReferences()
	for _,rs :=range ReferencesList {
		if rs.Kind=="ReplicaSet" {
			ReplicaSet:=this.RsMapStruct.Get(rs.Name,pod.Namespace)
			if ReplicaSet!=nil{
				for _,item :=range ReplicaSet.GetOwnerReferences(){
					if item.Kind==kind {
						return item.Name,pod.Namespace,nil
					}
				}
			}
		}
	}
	return "","",fmt.Errorf("not found %s",kind)
}
func(this *PodService) GetDepPod(ns ,name string) []*Pod{
	podList:= this.PodMap.ListByNs(ns)
	ret:=make([]*Pod,0)
	for _,pod:=range podList{
		depName,_,_:=this.FatherReplicaSetData(pod.Namespace,pod.Name,"Deployment")
		if depName==name{
			ret=append(ret,&Pod{
				Name:pod.Name, //pod名称
				NameSpace:pod.Namespace, //pod的命名空间
				Images:this.Common.GetImagesByPod(pod.Spec.Containers), //取得镜像
				NodeName:pod.Spec.NodeName, //取得调度的主机名称
				Phase:string(pod.Status.Phase),// 阶段
				IsReady:this.Common.PosIsReady(pod), //是否就绪
				IP:[]string{pod.Status.PodIP,pod.Status.HostIP}, //取ip
				Message:this.EventMap.GetMessage(pod.Namespace,"Pod",pod.Name,string(pod.UID)), //取事件信息
				CreateTime:pod.CreationTimestamp.Format("2006-01-02 15:04:05"),//取时间戳
			})
		}
	}
	return  ret
}
//获取Pod容器列表
func(this *PodService) GetPodContainer(ns,podName string) []*ContainerModel{
	ret:=make([]*ContainerModel,0)

	pod:=this.PodMap.Get(ns,podName)


	if pod!=nil{
		for _,c:=range pod.Spec.Containers{
			ret=append(ret,&ContainerModel{
				Name: c.Name,
			})
		}
	}
	return ret
}
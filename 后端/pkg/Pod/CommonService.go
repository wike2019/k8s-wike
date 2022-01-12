package Pod

import (
	"fmt"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

//@Service
type CommonService struct {

}
func NewCommonService() *CommonService {
	return &CommonService{}
}
//获取镜像名称
func(this *CommonService)  GetImages(dep v1.Deployment) string   {
	return this.GetImagesByPod(dep.Spec.Template.Spec.Containers)
}
//根据容器获取镜像名称
func(*CommonService) GetImagesByPod(containers []corev1.Container) string{
	images:=containers[0].Image
	if imgLen:=len(containers);imgLen>1{
		images+=fmt.Sprintf("+其他%d个镜像",imgLen-1)
	}
	return images
}
//判断是否就绪
func(*CommonService) PosIsReady(pod *corev1.Pod) bool{

	if pod.Status.Phase!="Running"{
		return false
	}
	for _,condition:=range pod.Status.Conditions{
		if  condition.Status!="True"{
			return false
		}
	}
	for _,rg:=range pod.Spec.ReadinessGates{
		for _,condition:=range pod.Status.Conditions{
			if condition.Type==rg.ConditionType && condition.Status!="True"{
				return false
			}
		}
	}
	return true
}
package configs

import (
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/ConfigMap"
	"k8sapi/pkg/Deployment"
	"k8sapi/pkg/Event"
	"k8sapi/pkg/Ingress"
	"k8sapi/pkg/Node"
	"k8sapi/pkg/Ns"
	"k8sapi/pkg/Pod"
	"k8sapi/pkg/Pvc"
	"k8sapi/pkg/Rbac"
	"k8sapi/pkg/Sa"
	"k8sapi/pkg/Secret"
	"k8sapi/pkg/Svc"
)

type K8sMaps struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}
func NewK8sMaps() *K8sMaps {
	return &K8sMaps{}
}
//初始化 deploymentmap
func(this *K8sMaps) InitDepMap() *Deployment.DeploymentMap{
	return &Deployment.DeploymentMap{}
}

//初始化 podmap
func(this *K8sMaps) InitPodMap() *Pod.PodMapStruct{
	return &Pod.PodMapStruct{}
}


//初始化 nsmap
func(this *K8sMaps) InitNsMap() *Ns.MapStruct{
	return &Ns.MapStruct{}
}

//初始化 eventmap
func(this *K8sMaps) InitEventMap() *Event.EventMapStruct{
	return &Event.EventMapStruct{}
}


//初始化 ingress map
func(this *K8sMaps) InitIngressMap() *Ingress.IngressMapStruct{
	return &Ingress.IngressMapStruct{}
}
//初始化 Service map
func(this *K8sMaps) InitServiceMap() *Svc.ServiceMapStruct{
	return &Svc.ServiceMapStruct{}
}

//初始化 Secret map
func(this *K8sMaps) InitSecretMap() *Secret.SecretMapStruct{
	return &Secret.SecretMapStruct{}
}


//初始化 ConfigMap map
func(this *K8sMaps) InitConfigMap() *ConfigMap.ConfigMapStruct{
	return &ConfigMap.ConfigMapStruct{}
}

//初始化NodeMap
func(this *K8sMaps) InitNodeMap() *Node.NodeMapStruct{
	return &Node.NodeMapStruct{}
}

//初始化RoleMap
func(this *K8sMaps) InitRoleMap() *Rbac.RoleMapStruct{
	return &Rbac.RoleMapStruct{}
}

//初始化RoleBindingMap
func(this *K8sMaps) InitRoleBindingMap() *Rbac.RoleBindingMapStruct{
	return &Rbac.RoleBindingMapStruct{}
}

//初始化RoleBindingMap
func(this *K8sMaps) InitSaMap() *Sa.MapStruct{
	return &Sa.MapStruct{}
}

//初始化ClusterRole
func(this *K8sMaps) InitClusterRoleMap() *Rbac.ClusterRoleMapStruct{
	return &Rbac.ClusterRoleMapStruct{}
}
//初始化ClusterRoleBinding
func(this *K8sMaps) InitClusterRoleBindingMap() *Rbac.ClusterRoleBindingMapStruct{
	return &Rbac.ClusterRoleBindingMapStruct{}
}

func(this *K8sMaps) PvcBindingMap() *Pvc.MapStruct{
	return &Pvc.MapStruct{}
}
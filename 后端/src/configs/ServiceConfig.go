package configs

import (
	"k8sapi/pkg/Common"
	"k8sapi/pkg/ConfigMap"
	"k8sapi/pkg/Deployment"
	"k8sapi/pkg/Ingress"
	"k8sapi/pkg/Node"
	"k8sapi/pkg/Pod"
	"k8sapi/pkg/Rbac"
	"k8sapi/pkg/Secret"
)

//@Config
type ServiceConfig struct {}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}
func(*ServiceConfig) PodCommonService() *Pod.CommonService{
	return Pod.NewCommonService()
}
func(*ServiceConfig) DeploymentCommonService() *Deployment.CommonService{
	return Deployment.NewCommonService()
}
func(*ServiceConfig) DeploymentService() *Deployment.DeploymentService{
	return Deployment.NewDeploymentService()
}
func(*ServiceConfig) PodService() *Pod.PodService{
	return Pod.NewPodService()
}
func(*ServiceConfig) CommonHelper() *Common.Helper{
	return Common.NewHelper()
}

func(*ServiceConfig) PodHelper() *Pod.Helper{
	return Pod.NewHelper()
}

func(*ServiceConfig) NodeHelper() *Node.Helper{
	return Node.NewHelper()
}


func(*ServiceConfig) IngressService() *Ingress.IngressService{
	return Ingress.NewIngressService()
}
func(*ServiceConfig) SecretService() *Secret.Service{
	return Secret.NewSecretService()
}
func(*ServiceConfig) ConfigMapService() *ConfigMap.ConfigMapService{
	return ConfigMap.NewConfigMapService()
}
func(*ServiceConfig) ConfigNodeService() *Node.NodeService{
	return Node.NewNodeService()
}
func(*ServiceConfig) ConfigRoleService() *Rbac.RoleService{
	return Rbac.NewRoleService()
}

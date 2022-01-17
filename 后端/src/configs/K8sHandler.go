package configs

import (
	"k8sapi/pkg/ConfigMap"
	"k8sapi/pkg/Deployment"
	"k8sapi/pkg/Event"
	"k8sapi/pkg/Ingress"
	"k8sapi/pkg/Node"
	"k8sapi/pkg/ns"
	"k8sapi/pkg/Pod"
	"k8sapi/pkg/Pvc"
	"k8sapi/pkg/Rbac"
	"k8sapi/pkg/Rs"
	"k8sapi/pkg/sa"
	"k8sapi/pkg/secret"
	"k8sapi/pkg/Svc"
)




//注入 回调handler
type K8sHandler struct {}

func NewK8sHandler() *K8sHandler {
	return &K8sHandler{}
}
// deployment handler
func(this *K8sHandler) DepHandlers() *Deployment.DepHandler{
	return &Deployment.DepHandler{}
}
// pod handler
func(this *K8sHandler) PodHandlers() *Pod.PodHandler{
	return &Pod.PodHandler{}
}
//ns handler
func(this *K8sHandler) NsHandlers() *ns.NsHandler{
	return &ns.NsHandler{}
}
// event handler
func(this *K8sHandler) EventHandlers() *event.EventHandler{
	return &event.EventHandler{}
}

// IngressHandler
func(this *K8sHandler) IngressHandler() *Ingress.IngressHandler{
	return &Ingress.IngressHandler{}
}
// ServiceHandler
func(this *K8sHandler) ServiceHandler() *Svc.ServiceHandler{
	return &Svc.ServiceHandler{}
}
// SecretHandler
func(this *K8sHandler) SecretHandler() *secret.SecretHandler{
	return &secret.SecretHandler{}
}
// ConfigMapHandler
func(this *K8sHandler) ConfigMapHandler() *ConfigMap.ConfigMapHandler{
	return &ConfigMap.ConfigMapHandler{}
}

// ConfigMapHandler
func(this *K8sHandler) ConfigNodeHandler() *Node.NodeMapHandler{
	return &Node.NodeMapHandler{}
}

// RoleHandler
func(this *K8sHandler) ConfigRoleHandler() *Rbac.RoleHander{
	return &Rbac.RoleHander{}
}
// RoleBindingHandler
func(this *K8sHandler) ConfigRoleBindingHandler() *Rbac.RoleBindingHander{
	return &Rbac.RoleBindingHander{}
}

// RoleBindingHandler
func(this *K8sHandler) ConfigSaHandler() *sa.Handler{
	return &sa.Handler{}
}

// ClusterRoleHandler
func(this *K8sHandler) ConfigClusterRoleHandler() *Rbac.ClusterRoleHandler{
	return &Rbac.ClusterRoleHandler{}
}

// ClusterRoleBindingHandler
func(this *K8sHandler) ConfigClusterRoleBindingHandler() *Rbac.ClusterRoleBindingHander{
	return &Rbac.ClusterRoleBindingHander{}
}

func(this *K8sHandler) ConfigPvcBindingHandler() *Pvc.Handler{
	return &Pvc.Handler{}
}

func(this *K8sHandler) RsBindingHandler() *Rs.Handler{
	return &Rs.Handler{}
}

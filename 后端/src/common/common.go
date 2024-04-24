package common

import (
	"github.com/gin-gonic/gin"
	coreHttp "github.com/wike2019/wike_go/http"
	"github.com/wike2019/wike_go/lib/controller"
	"k8s.io/client-go/kubernetes"
)

type CommonService struct {
	controller.Controller
	Client *kubernetes.Clientset
}

func NewService(client *kubernetes.Clientset) *CommonService {
	return &CommonService{
		Client: client,
	}
}
func (r *CommonService) Path() string {
	return "/api/v1"
}

func (this *CommonService) Build(r *gin.RouterGroup, GCore *coreHttp.GCore) {
	GCore.RoleCtl.AddRole("admin", "guest")
	GCore.RoleCtl.AddRole("wike", "admin")
	GCore.GetWithRbac(r, "guest", "/ws", controller.WsHandler)
}

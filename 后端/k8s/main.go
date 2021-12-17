package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/pkg/ConfigMap"
	"k8sapi/pkg/Deployment"
	"k8sapi/pkg/Ingress"
	"k8sapi/pkg/Node"
	"k8sapi/pkg/Ns"
	"k8sapi/pkg/Pod"
	"k8sapi/pkg/Pvc"
	"k8sapi/pkg/Rbac"
	"k8sapi/pkg/Resources"
	"k8sapi/pkg/Sa"
	"k8sapi/pkg/Secret"
	"k8sapi/pkg/Svc"
	"k8sapi/pkg/User"
	"k8sapi/pkg/Ws"
	"k8sapi/src/configs"
	"net/http"
)

//go:embed html
var dist  embed.FS

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			fmt.Println(c.Request.URL)
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func main() {

	web:=goft.Ignite(Cors()).Config(
		configs.NewK8sHandler(),  //1
		configs.NewK8sConfig(), //2
		configs.NewK8sMaps(), //3
		configs.NewServiceConfig(),
	).
		Mount("v1",
			Deployment.NewDeploymentCtl(),
			Pod.NewPodCtl(),
			Ws.NewWsCtl(),
			Ns.NewNsCtl(),
			User.NewUserCtl(),
			Ingress.NewIngressCtl(),
			Svc.NewSvcCtl(),
			Secret.NewSecretCtl(),
			ConfigMap.NewConfigMapCtl(),
			Node.NewNodeCtl(),
			Rbac.NewRBACCtl(),
			Resources.NewResourcesCtl(),
			Sa.NewSaCtl(),
			Pvc.NewSaCtl(),
		)
	web.Engine.GET("/", func(c *gin.Context) {
		c.FileFromFS("/html/"+c.Param("filepath"),http.FS(dist))
	})
	web.Launch()

}

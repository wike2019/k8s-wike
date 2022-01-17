package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/pkg/ConfigMap"
	"k8sapi/pkg/Deployment"
	event "k8sapi/pkg/Event"
	"k8sapi/pkg/Ingress"
	"k8sapi/pkg/Node"
	"k8sapi/pkg/Pod"
	"k8sapi/pkg/Pvc"
	"k8sapi/pkg/Rbac"
	"k8sapi/pkg/Resources"
	"k8sapi/pkg/secret"
	"k8sapi/pkg/Svc"
	"k8sapi/pkg/User"
	"k8sapi/pkg/Ws"
	"k8sapi/pkg/common"
	"k8sapi/pkg/ns"
	"k8sapi/pkg/sa"
	"k8sapi/src/configs"
	"net/http"
)

//go:embed html
var dist  embed.FS
var keywords=[]string{
	"apiVersion",
	"Kind",
	"metadata",
	"name",
	"namespace",
	"labels",
	"annotations",
	"type",
	"stringData",
	"data",
	"Opaque",
	"secrets",
	"imagePullSecrets",
}
var keywordsValue=[]string{
	"kubernetes.io/tls",
	"kubernetes.io/dockercfg",
	"Namespace",
	"v1",
}


func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}



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
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func main() {
	web:=goft.Ignite(Cors(), func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.AbortWithStatusJSON(400, gin.H{"error": err})
			}
		}()
		context.Next()
	}).Config(
		configs.NewK8sHandler(),  //1
		configs.NewK8sConfig(), //2
		configs.NewK8sMaps(), //3
		configs.NewServiceConfig(),
	).
		Mount("v1",
			Deployment.NewDeploymentCtl(),
			Pod.NewPodCtl(),
			Ws.NewWsCtl(),
			ns.NewNsCtl(),
			User.NewUserCtl(),
			Ingress.NewIngressCtl(),
			Svc.NewSvcCtl(),
			secret.NewSecretCtl(),
			ConfigMap.NewConfigMapCtl(),
			Node.NewNodeCtl(),
			Rbac.NewRBACCtl(),
			Resources.NewResourcesCtl(),
			sa.NewSaCtl(),
			Pvc.NewSaCtl(),
			common.NewOtherCtl(),
			event.NewEventCtl(),
		)
	web.Engine.GET("/", func(c *gin.Context) {
		c.FileFromFS("/html/"+c.Param("filepath"),http.FS(dist))
	})
	for _,v :=range removeDuplicateElement(keywords){
		common.AddSySAutoComplete(v)
	}
	for _,v :=range removeDuplicateElement(keywordsValue){
		common.AddValueAutoComplete(v)
	}
	web.Launch()

}

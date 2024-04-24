package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wike2019/k8s-wike/config"
	"github.com/wike2019/k8s-wike/src/common"
	"github.com/wike2019/k8s-wike/src/informer"
	"github.com/wike2019/k8s-wike/src/namespace"
	coreHttp "github.com/wike2019/wike_go/http"
	"k8s.io/client-go/informers"
)

func main() {

	// @title           k8s管理系统 API 文档
	// @version         2.0
	// @description     这个是k8s管理系统 API 文档
	// @contact.email  200569525@qq.com
	// @license.name  Apache 2.0
	// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
	// @host      k8s.ng2-oa.com
	// @BasePath  /api/v2/
	coreHttp.God().
		GlobalUse(func(context *gin.Context) {
			context.Set("role", "admin")
		}).
		Config(config.NewK8s()).
		Provide(config.Config).
		Provide(namespace.Ns).
		Provide(informer.InitInformer).
		MountWithEmpty(namespace.NewService).
		MountWithEmpty(common.NewService).

		//Cron("* * * * *", func() {
		//	fmt.Println("第一行")
		//}).

		//MountWithEmpty(router.GtRouterInit).
		//Mount(router.HomeRouterInit, coreHttp.CreateTag("redis")).
		//MountWithEmpty(router.ServerRouterInit).
		//Mount(router.NewsRouterInit, coreHttp.CreateTag("redis")).
		//Mount(router.GameRouterInit, router.CommonRouterInit).
		Invokes(func(r informers.SharedInformerFactory) {
			//	fmt.Println("执行验证码任务")
			//	//controllers.CheckBypassStatus()
		}).
		Run()
}

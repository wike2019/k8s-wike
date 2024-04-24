package router

import (
	"github.com/gin-gonic/gin"
	coreHttp "github.com/wike2019/wike_go/http"
	"github.com/wike2019/wike_go/lib/controller"
)

type StaticRouter struct {
	controller.Controller
}

func (r *StaticRouter) Path() string {
	return "/"
}
func StaticRouterInit() *StaticRouter {
	r := &StaticRouter{}
	return r
}
func (this *StaticRouter) Build(r *gin.RouterGroup, GCore *coreHttp.GCore) {
	r.Static("/assets", "./public/dist/assets")
	r.StaticFile("/", "./public/dist/index.html")
	r.Static("/newsuploads", "./public/newsuploads")
	r.Static("/ueditor", "./public/ueditor")
	r.Static("/ueditor1_4_3_3", "./public/ueditor1_4_3_3")
	r.Static("/uploads", "./public/uploads")
}

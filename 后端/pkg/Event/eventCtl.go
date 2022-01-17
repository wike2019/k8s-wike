package event

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type EventCtl struct{
	Map *EventMapStruct `inject:"-"`
}
func NewEventCtl() *EventCtl {
  return &EventCtl{}
}
//获取全部 不带分页
func(this *EventCtl) listAll(c *gin.Context) goft.Json{
	name:=c.DefaultQuery("name","")
	ns:=c.DefaultQuery("namespace","")
	kind:=c.DefaultQuery("kind","")
	uid:=c.DefaultQuery("uid","")
	return gin.H{"code":200, "data":this.Map.GetMessage(ns,kind,name,uid)}
}



//路由
func(this *EventCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/event",this.listAll)

}
func(this *EventCtl)  Name() string{
	 return "NsCtl"
}
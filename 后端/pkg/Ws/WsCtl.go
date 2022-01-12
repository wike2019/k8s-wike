package Ws

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/wscore"
	"log"
)

//@Controller
type WsCtl struct {}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func(this *WsCtl) Connect(c *gin.Context)  (v goft.Void)  {
	client,err:=wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)  //升级
	if err!=nil{
		log.Println(err)
		return
	}else{
		wscore.ClientMap.Store(client)
		return
	}

}
func(this *WsCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/ws",this.Connect)
}
func(this *WsCtl) Name() string{
	return "WsCtl"
}

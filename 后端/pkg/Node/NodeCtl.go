package Node

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"golang.org/x/crypto/ssh"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/Common"
	"k8sapi/src/wscore"
)

//@controller
type NodeCtl struct {
	NodeService *NodeService `inject:"-"`
	Client *kubernetes.Clientset  `inject:"-"`
	Helper *Common.Helper `inject:"-"`
	NodeHelper *Helper `inject:"-"`
}
var NodeShellModes  = ssh.TerminalModes{
	ssh.ECHO:          1,     // enable echoing
	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
	ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
}

func NewNodeCtl() *NodeCtl {
	return &NodeCtl{}
}
func(this *NodeCtl) ListAll(c *gin.Context) goft.Json{
	return gin.H{
		"code":200,
		"data":this.NodeService.ListAllNodes(),
	}

}
func(this *NodeCtl) LoadDetail(c *gin.Context) goft.Json{
	nodeName:=c.Query("node")
	return gin.H{
		"code":200,
		"data":this.NodeService.LoadNode(nodeName),
	}

}

//保存node
func(this *NodeCtl) SaveNode(c *gin.Context) goft.Json{
	nodeModel:=&PostNodeModel{}
	_=c.ShouldBindJSON(nodeModel)
	node:=this.NodeService.LoadOriginNode(nodeModel.Name) //取出原始node 信息
	if node==nil{
		goft.Error(fmt.Errorf("资源没找到"))
	}
	node.Labels=nodeModel.OrginLabels  //覆盖标签
	node.Spec.Taints=nodeModel.OrginTaints //覆盖 污点
	_,err:=this.Client.CoreV1().Nodes().Update(c,node,v1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(*NodeCtl) Name() string{
	return "NodeCtl"
}

func(this *NodeCtl) Shell(c *gin.Context) (v goft.Void){
	defer func() {
		recover()
		//fmt.Println(err)
		fmt.Println("客户端关闭")

	}()
	wsClient,err:=wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)
	ip:=c.Query("ip")
	name:=c.Query("name")
	password:=c.Query("password")
	goft.Error(err)
	shellClient:=wscore.NewWsShellClient(wsClient)
	//session, err := helpers.SSHConnect(helpers.TempSSHUser,  helpers.TempSSHPWD, helpers.TempSSHIP ,22 )
	session, err := this.NodeHelper.SSHConnect(name, password, ip ,22 )
	if err!=nil {
		wsClient.WriteMessage(1,[]byte(err.Error()))
		session.Close()
	}

	defer session.Close()
	session.Stdout = shellClient
	session.Stderr =shellClient
	session.Stdin=shellClient
	err= session.RequestPty("xterm-256color", 300, 500,NodeShellModes)
	goft.Error(err)

	err=session.Run("/bin/bash")
	goft.Error(err)
	return
}

func(this *NodeCtl)  Build(goft *goft.Goft) {
	goft.Handle("GET", "/nodes", this.ListAll)
	goft.Handle("GET", "/nodes/detail", this.LoadDetail) //加载详细
	goft.Handle("PUT", "/nodes", this.SaveNode)          //保存
	goft.Handle("GET","/ws/node/shell",this.Shell )
}
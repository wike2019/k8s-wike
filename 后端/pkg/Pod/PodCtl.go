package Pod

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"k8sapi/pkg/Common"
	"k8sapi/src/wscore"
	"os"
	"time"
)

type PodCtl struct {
	PodService *PodService `inject:"-"`
	Helper *Common.Helper `inject:"-"`
	Client *kubernetes.Clientset `inject:"-"`
	Config *rest.Config `inject:"-"`
	podHelper *Helper `inject:"-"`
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}
func(this *PodCtl) getList(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","5")
	return gin.H{
		"code":200,
		"data":this.PodService.PagePods(ns,this.Helper.StrToInt(page,1),
			this.Helper.StrToInt(size,5)),
	}


}


//获取 容器
func(this *PodCtl) Containers(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	podName:=c.DefaultQuery("name","")
	return gin.H{
		"code":200,
		"data":this.PodService.GetPodContainer(ns,podName),
	}
}
func(this *PodCtl) GetLogs(c *gin.Context) (v goft.Void){
	defer func() {
		recover()
		fmt.Println("客户端关闭")
	}()
	cc,_:=context.WithTimeout(c,time.Minute*30)
	ns:=c.DefaultQuery("ns","default")
	podname:=c.DefaultQuery("podname","")
	cname:=c.DefaultQuery("cname","")
	Client,_:=wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)  //升级
	count:=int64(100);
	req:=this.Client.CoreV1().Pods(ns).GetLogs(podname,&v1.PodLogOptions{Container:cname,Follow: true,TailLines:&count})
	reader,err:=req.Stream(cc)
	goft.Error(err)
	for{

		buf:=make([]byte,1024)
		n,err:=reader.Read(buf)
		if err!=nil && err!=io.EOF{
			fmt.Println(err)
			break
		}


		if n>0{
			fmt.Println("有内容2")
			err:=Client.WriteMessage(1,buf[0:n])
			if err!=nil {
				fmt.Println(err)
				break
			}
		}

	}
	fmt.Println("已结束")
	return

}
func(this *PodCtl) Shell(c *gin.Context) (v goft.Void){
	defer func() {
		recover()
		//fmt.Println(err)
		fmt.Println("客户端关闭")
	}()
	wsClient,err:=wscore.Upgrader.Upgrade(c.Writer,c.Request,nil)
	ns:=c.DefaultQuery("ns","default")
	podname:=c.DefaultQuery("podname","")
	cname:=c.DefaultQuery("cname","")
	goft.Error(err)
	shellClient:=wscore.NewWsShellClient(wsClient)
	option := &v1.PodExecOptions{
		Container:cname,
		Command: []string{"/bin/bash","-c","ls"},
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
	}
	req:=this.Client.CoreV1().RESTClient().Post().Resource("pods").
		Namespace(ns).
		Name(podname).
		SubResource("exec").VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	fmt.Println(req.URL())
	SHEll:="/bin/bash"

	exec,err:=remotecommand.NewSPDYExecutor(this.Config,"POST",
		req.URL())
	if err!=nil{
		fmt.Println("错误发生了")
		SHEll="/bin/sh"
	}
	err=exec.Stream(remotecommand.StreamOptions{
		Stdin:os.Stdin,
		Stdout:os.Stdout,
		Stderr:os.Stderr,
		Tty:true,
	})
	if err!=nil{
		fmt.Println("错误发生了")
		SHEll="/bin/sh"
	}
	err=this.podHelper.HandleCommand(this.Client,this.Config,[]string{SHEll},ns,podname,cname).
		Stream(remotecommand.StreamOptions{
			Stdin:shellClient,
			Stdout:shellClient,
			Stderr:shellClient,
			Tty:true,
		})


	goft.Error(err)
	fmt.Println("结束了222")
	return
}
//获取 容器
func(this *PodCtl) GetDepPod(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	depName:=c.DefaultQuery("name","")
	return gin.H{
		"code":200,
		"data":this.PodService.GetDepPod(ns,depName),
	}
}
func(this *PodCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/pods",this.getList)
	goft.Handle("GET","/dep/pods",this.GetDepPod)
	goft.Handle("GET","/pods/containers",this.Containers)
	goft.Handle("GET","/ws/pods/logs",this.GetLogs )
	goft.Handle("GET","/ws/pods/shell",this.Shell )

}
func(this *PodCtl) Name() string{
	return "PodCtl"
}
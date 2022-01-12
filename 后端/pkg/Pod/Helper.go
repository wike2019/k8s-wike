package Pod

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

//@Component
type Helper struct {
	//	PodMap *Pod.PodMapStruct `inject:"-"`
}
func NewHelper() *Helper {
	return &Helper{}
}

func(this *Helper)  HandleCommand(client *kubernetes.Clientset,config *rest.Config,command []string,ns string,podname string,cname string) remotecommand.Executor{
	option := &v1.PodExecOptions{
		Container:cname,
		Command: command,
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:true,
	}
	req:=client.CoreV1().RESTClient().Post().Resource("pods").
		Namespace(ns).
		Name(podname).
		SubResource("exec").
		Param("color","true").
		VersionedParams(
			option,
			scheme.ParameterCodec,
		)

	exec,err:=remotecommand.NewSPDYExecutor(config,"POST",
		req.URL())
	if err!=nil{
		panic(err)
	}
	return exec
}


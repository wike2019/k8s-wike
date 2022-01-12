package Node

import (
	"context"
	"fmt"
	"golang.org/x/crypto/ssh"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"k8sapi/pkg/Pod"
	"net"
	"regexp"
)


//@Component
type Helper struct {
	PodMap *Pod.PodMapStruct `inject:"-"`
}
func NewHelper() *Helper {
	return &Helper{}
}

func(this *Helper)  showLable(key string ) bool{
	const hostPattern="[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\\.?"
	return !regexp.MustCompile(hostPattern).MatchString(key)
}
//过滤 要显示的标签
func(this *Helper) FilterLables(labels map[string]string) []string{
	result:=make([]string,0)
	for k,v:=range labels{
		if this.showLable(k){
			result=append(result,fmt.Sprintf("%s=%s",k,v))
		}
	}
	return result
}
func(this *Helper) FilterTaints(taints []v1.Taint) []string{
	result:=make([]string,0)
	for _,taint:=range taints{
		result=append(result,fmt.Sprintf("%s=%s:%s",taint.Key,taint.Value,taint.Effect))
	}
	return result
}

//第一个是cpu使用 第二个是内存使用
func(this *Helper) GetNodeUsage(c *versioned.Clientset,node *v1.Node) *NodeUsage{
	nodeMetric, _ := c.MetricsV1beta1().
		NodeMetricses().Get(context.Background(),node.Name,metav1.GetOptions{})
	cpu:=float64(nodeMetric.Usage.Cpu().MilliValue())/float64(node.Status.Capacity.Cpu().MilliValue())
	memory:=float64(nodeMetric.Usage.Memory().MilliValue())/float64(node.Status.Capacity.Memory().MilliValue())
	return &NodeUsage{
		Cpu: cpu,
		Memory: memory,
		Pods: int64(this.PodMap.GetNum(node.Name)),
	}
}


func(this *Helper) SSHConnect( user, password, host string, port int ) ( *ssh.Session, error ) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User:               user,
		Auth:               auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback:    hostKeyCallbk,
	}
	// connet to ssh
	addr = fmt.Sprintf( "%s:%d", host, port )
	if client, err = ssh.Dial( "tcp", addr, clientConfig ); err != nil {
		return nil, err
	}
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}
package config

import (
	"context"
	"github.com/wike2019/wike_go/lib/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"log"
)

type K8s struct {
}

func NewK8s() *K8s {
	return &K8s{}
}

func (this *K8s) Ns() *Box[*corev1.Namespace] {
	return &Box[*corev1.Namespace]{
		Content: utils.NewMap[*corev1.Namespace](),
	}
}

func (this *K8s) InitConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", "k8sConfig")
	if err != nil {
		log.Fatal(err)
	}
	config.Insecure = false
	return config
}
func (this *K8s) InitClient() *kubernetes.Clientset {

	c, err := kubernetes.NewForConfig(this.InitConfig())
	if err != nil {
		log.Fatal(err)
	}
	namespace := &corev1.Namespace{
		TypeMeta:   metav1.TypeMeta{APIVersion: corev1.SchemeGroupVersion.String(), Kind: "Namespace"},
		ObjectMeta: metav1.ObjectMeta{Name: "wike-system"},
	}

	_, err = c.CoreV1().Namespaces().Create(context.Background(), namespace, metav1.CreateOptions{})
	return c
}
func (this *K8s) InitMetricClient() *versioned.Clientset {

	c, err := versioned.NewForConfig(this.InitConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

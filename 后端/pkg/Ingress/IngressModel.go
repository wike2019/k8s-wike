package Ingress

import v1 "k8s.io/api/networking/v1"

type IngressModel struct {
	Name string `json:"name"`
	NameSpace string `json:"namespace"`
	CreateTime string `json:"create_time"`
	Host string `json:"host"`
	IsHttps bool `json:"isHttps"`
	Annotations map[string]string `json:"annotations"`
	IsReady bool `json:"is_ready"`
	TLS []v1.IngressTLS `json:"tls"`
	Rules []v1.IngressRule `json:"rules"`
}


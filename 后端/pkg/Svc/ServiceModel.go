package Svc

import v1 "k8s.io/api/core/v1"

type ServiceModel struct {
	Name string `json:"name"`
	NameSpace string `json:"namespace"`
	CreateTime string `json:"create_time"`
	ServicePort []v1.ServicePort `json:"port"`
 }

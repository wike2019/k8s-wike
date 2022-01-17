package secret

import (
	"context"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
	"k8sapi/pkg/helper"
)

type SecretService struct {
	Client *kubernetes.Clientset `inject:"-"`
	Map *SecretMapStruct               `inject:"-"`
	Helper *helper.Helper        `inject:"-"` //帮助函数 用于分页
}
func NewSecretService() *SecretService {
	return &SecretService{}
}


func(this *SecretService) PageDeps(ns string,page int) *helper.ItemsPage {
	SecretList:=this.Map.ListAll(ns)
	SecretCovert:=make([]interface{},len(SecretList))
	for i,dep:=range SecretList{
		SecretCovert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		10,
		SecretCovert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":ns, //命名空间
	})
}
func(this *SecretService) All(ns string )[]*SecretModel{
	return this.Map.ListAll(ns)
}


func(this *SecretService) GetSecret(ns,name string) SecretModel{
	return this.Map.GetSecret(ns,name)
}


func(this *SecretService) DelSecret(ns,name string) error{
	return this.Client.CoreV1().Secrets(ns).
		Delete(context.Background(),name,v1.DeleteOptions{})
}
func (this *SecretService)  encodeDockerConfigFieldAuth(username, password string) string {
	fieldValue := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(fieldValue))
}

func (this *SecretService)  handleDockerCfgJSONContent(username, password, email, server string) (string, error) {
	dockerConfigAuth := DockerConfigEntry{
		Username: username,
		Password: password,
		Email:    email,
		Auth:     this.encodeDockerConfigFieldAuth(username, password),
	}
	dockerConfigJSON := DockerConfigJSON{
		Auths: map[string]DockerConfigEntry{server: dockerConfigAuth},
	}
	data,err:=json.Marshal(dockerConfigJSON)
	return string(data),err
}
func(this *SecretService) PostSecret(postModel *corev1.Secret) error{
	data :=make(map[string]string)
	if postModel.Type=="kubernetes.io/dockercfg"{
		username:=postModel.StringData["username"]
		server:=postModel.StringData["server"]
		password:=postModel.StringData["password"]
		email:=postModel.StringData["email"]
		json,err:=this.handleDockerCfgJSONContent(username,password,email,server)
		if err !=nil{
			return err
		}

		data[".dockercfg"]=json
		postModel.StringData=data
	}
	if postModel.Type=="kubernetes.io/dockerconfigjson"{
		username:=postModel.StringData["username"]
		server:=postModel.StringData["server"]
		password:=postModel.StringData["password"]
		email:=postModel.StringData["email"]
		json,err:=this.handleDockerCfgJSONContent(username,password,email,server)
		if err !=nil{
			return err
		}

		data[".dockerconfigjson"]=json
		postModel.StringData=data
	}
	_,err:=this.Client.CoreV1().Secrets(postModel.Namespace).Create(
		context.Background(),
		postModel,
		v1.CreateOptions{},
	)
	return err

}
func(this *SecretService) UpdateSecret(postModel *corev1.Secret) error{
	_,err:=this.Client.CoreV1().Secrets(postModel.Namespace).Update(
		context.Background(),
		postModel,
		v1.UpdateOptions{},
	)
	return err

}
package secret

var SECRET_TYPE map[string]string

func init() {
	SECRET_TYPE= map[string]string{
		"Opaque":"自定义类型",
		"kubernetes.io/service-account-token":"服务账号令牌",
		"kubernetes.io/dockercfg":"docker配置",
		"kubernetes.io/dockerconfigjson":"docker配置(JSON)",
		"kubernetes.io/basic-auth":"Basic认证凭据",
		"kubernetes.io/ssh-auth":" SSH凭据",
		"kubernetes.io/tls":"TLS凭据",
		"bootstrap.kubernetes.io/token":"启动引导令牌数据",
		"helm.sh/release.v1":"helm相关密文",
	}
}

type SecretModel struct {
	Name string `json:"name"`
	NameSpace string `json:"namespace"`
	CreateTime string `json:"create_time"`
	Type string  `json:"type"`
	Data map[string][]byte `json:"data"`
	StringData map[string]string `json:"string_data"`
	TlsData CertModel `json:"tls_data"`
	TypeRaw string `json:"typeRaw"`
	Annotations map[string]string `json:"annotations"`
	Labels  map[string]string`json:"labels"`
}


type CertModel struct {
	CN string  //域名
	Algorithm string //算法
	Issuer string //签发者
	BeginTime string //生效时间
	EndTime string //到期时间
}

type DockerConfigEntry struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty" datapolicy:"password"`
	Email    string `json:"email,omitempty"`
	Auth     string `json:"auth,omitempty" datapolicy:"token"`
}
type DockerConfigJSON struct {
	Auths DockerConfig `json:"auths" datapolicy:"token"`
	// +optional
	HttpHeaders map[string]string `json:"HttpHeaders,omitempty" datapolicy:"token"`
}
type DockerConfig map[string]DockerConfigEntry
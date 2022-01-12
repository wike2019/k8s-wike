package Rbac
import "k8s.io/api/rbac/v1"
type RoleModel struct {
	Name string `json:"name"`
	NameSpace string `json:"name_space"`
	CreateTime string `json:"create_time"`
}

type RoleBindingModel struct {
	Name string `json:"name"`
	NameSpace string `json:"name_space"`
	CreateTime string     `json:"create_time"`
	RoleRef v1.RoleRef
	Subject []v1.Subject `json:"subject"`
}

type ClusterRoleModel struct {
	Name string `json:"name"`
	CreateTime string `json:"create_time"`
}

//提交 用的
type PostRoleModel struct {
	Name string `json:"name"`
	NameSpace string `json:"name_space"`
}


//UserAccount 模型
type UAModel struct {
	Name string `json:"name"`
	CreateTime string `json:"create_time"`
}
//提交用户时的对象模型
type PostUAModel struct {
	CN string `json:"cn" binding:"required,min=2"`
	O string  `json:"o"`
}

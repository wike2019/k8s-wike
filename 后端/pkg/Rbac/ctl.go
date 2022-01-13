package Rbac

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8sapi/pkg/helper"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type RBACCtl struct {
	RoleService *RoleService     `inject:"-"`
	Helper  *helper.Helper       `inject:"-"`
	Client *kubernetes.Clientset `inject:"-"`
}


func NewRBACCtl() *RBACCtl {
	return &RBACCtl{}
}
func(this *RBACCtl) Roles(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","5")
	return gin.H{
		"code":200,
		"data":this.RoleService.PageDeps(ns,this.Helper.StrToInt(page,1),this.Helper.StrToInt(size,5)),
	}
}
func(this *RBACCtl) DeleteRole(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	err:=this.Client.RbacV1().Roles(ns).Delete(c,name,metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *RBACCtl) RoleBindingList(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","5")
	return gin.H{
		"code":200,
		"data":this.RoleService.BindPageDeps(ns,this.Helper.StrToInt(page,1),this.Helper.StrToInt(size,5)),
	}
}
func(this *RBACCtl) CreateRole(c *gin.Context) goft.Json{
	role:=&rbacv1.Role{} //原生的k8s role 对象
	goft.Error(c.ShouldBindJSON(role))
	role.APIVersion="rbac.authorization.k8s.io/v1"
	role.Kind="Role"
	_,err:=this.Client.RbacV1().Roles(role.Namespace).Create(c,role,metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *RBACCtl) CreateClusterRoles(c *gin.Context) goft.Json{
	ClusterRole:=&rbacv1.ClusterRole{} //原生的k8s role 对象
	goft.Error(c.ShouldBindJSON(ClusterRole))
	ClusterRole.APIVersion="rbac.authorization.k8s.io/v1"
	ClusterRole.Kind="ClusterRole"
	_,err:=this.Client.RbacV1().ClusterRoles().Create(c,ClusterRole,metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}

func(this *RBACCtl) RolesAll(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	return gin.H{
		"code":200,
		"data":this.RoleService.rolesAll(ns),
	}
}

func(this *RBACCtl) DeleteRoleBinding(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")
	err:=this.Client.RbacV1().RoleBindings(ns).Delete(c,name,metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}

func(this *RBACCtl) CreateRoleBinding(c *gin.Context) goft.Json{
	rb:=&rbacv1.RoleBinding{}
	goft.Error(c.ShouldBindJSON(rb))
	_,err:=this.Client.RbacV1().RoleBindings(rb.Namespace).Create(c,rb,metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}

//从rolebinding中 增加或删除用户
func(this *RBACCtl) AddUserToRoleBinding(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")//rolebinding 名称
	t:=c.DefaultQuery("type","") //如果没传值就是增加，传值（不管什么代表删除)
	subject:=rbacv1.Subject{}// 传过来
	goft.Error(c.ShouldBindJSON(&subject))
	if subject.Kind=="ServiceAccount"{
		subject.APIGroup=""
	}
	rb:=this.RoleService.GetRoleBinding(ns,name) //通过名称获取 rolebinding对象
	if t!=""{//代表删除


		for i,sub:=range rb.Subjects{
			if sub.Kind==subject.Kind && sub.Name==subject.Name{
				rb.Subjects=append(rb.Subjects[:i], rb.Subjects[i+1:]...)
				break //确保只删一个（哪怕有同名同kind用户)
			}
		}
		fmt.Println(rb.Subjects)
	}else{
		rb.Subjects=append(rb.Subjects,subject)
	}
	_,err:=this.Client.RbacV1().RoleBindings(ns).Update(c,rb,metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}

//获取角色详细
func(this *RBACCtl) RolesDetail(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")//rolebinding 名称
	return gin.H{
		"code":200,
		"data":this.RoleService.GetRole(ns,name),
	}
}
//更新角色
func(this *RBACCtl) UpdateRolesDetail(c *gin.Context) goft.Json{
	ns:=c.DefaultQuery("ns","default")
	name:=c.DefaultQuery("name","")//rolebinding 名称
	role:=this.RoleService.GetRole(ns,name)
	fmt.Println(name)
	fmt.Println(ns)
	postRole:=rbacv1.Role{}
	goft.Error(c.ShouldBindJSON(&postRole))  //获取提交过来的对象

	role.Rules=postRole.Rules //目前修改只允许修改 rules，其他不允许。大家可以自行扩展，如标签也允许修改
	_,err:=this.Client.RbacV1().Roles(role.Namespace).Update(c,role,metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *RBACCtl) UpdateClusterRolesDetail(c *gin.Context) goft.Json{
	name:=c.DefaultQuery("name","")//rolebinding 名称
	ClusterRoleNew:=this.RoleService.GetClusterRole(name)

	ClusterRole:=rbacv1.ClusterRole{}
	goft.Error(c.ShouldBindJSON(&ClusterRole))  //获取提交过来的对象

	ClusterRoleNew.Rules=ClusterRole.Rules //目前修改只允许修改 rules，其他不允许。大家可以自行扩展，如标签也允许修改
	_,err:=this.Client.RbacV1().ClusterRoles().Update(c,ClusterRoleNew,metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}

func(*RBACCtl) Name() string{
	return "RBACCtl"
}

func(this *RBACCtl) clusterRolesAll(c *gin.Context) goft.Json{
	return gin.H{
		"code":200,
		"data":this.RoleService.ALlListClusterRoles(),
	}
}
func(this *RBACCtl) clusterRoles(c *gin.Context) goft.Json{
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","5")
	return gin.H{
		"code":200,
		"data":this.RoleService.ListClusterRoles(this.Helper.StrToInt(page,1),this.Helper.StrToInt(size,5)),
	}
}
func(this *RBACCtl) DeleteClusterRoles(c *gin.Context) goft.Json{
	name:=c.DefaultQuery("name","")
	err:=this.Client.RbacV1().ClusterRoles().Delete(c,name,metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}


//获取角色详细
func(this *RBACCtl) ClusterRolesDetail(c *gin.Context) goft.Json{
	name:=c.DefaultQuery("name","")//rolebinding 名称
	return gin.H{
		"code":200,
		"data":this.RoleService.GetClusterRole(name),
	}
}

func(this *RBACCtl) DeleteClusterRoleBinding(c *gin.Context) goft.Json{

	name:=c.DefaultQuery("name","")
	err:=this.Client.RbacV1().ClusterRoleBindings().Delete(c,name,metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *RBACCtl) AddUserToClusterRoleBinding(c *gin.Context) goft.Json{

	name:=c.DefaultQuery("name","")//clusterrolebinding 名称
	t:=c.DefaultQuery("type","") //如果没传值就是增加，传值（不管什么代表删除)
	subject:=rbacv1.Subject{}// 传过来
	goft.Error(c.ShouldBindJSON(&subject))
	if subject.Kind=="ServiceAccount"{
		subject.APIGroup=""
	}
	rb:=this.RoleService.GetClusterRoleBinding(name) //通过名称获取 clusterrolebinding对象
	if t!=""{//代表删除
		for i,sub:=range rb.Subjects{
			if sub.Kind==subject.Kind && sub.Name==subject.Name&&sub.Namespace==subject.Namespace{
				rb.Subjects=append(rb.Subjects[:i], rb.Subjects[i+1:]...)
				break //确保只删一个（哪怕有同名同kind用户)
			}
		}
	}else{
		rb.Subjects=append(rb.Subjects,subject)
	}
	_,err:=this.Client.RbacV1().ClusterRoleBindings().Update(c,rb,metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *RBACCtl) CreateClusterRoleBinding(c *gin.Context) goft.Json{
	rb:=&rbacv1.ClusterRoleBinding{}
	goft.Error(c.ShouldBindJSON(rb))
	_,err:=this.Client.RbacV1().ClusterRoleBindings().Create(c,rb,metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":"success",
	}
}
func(this *RBACCtl) ClusterRoleBindingList(c *gin.Context) goft.Json{
	page:=c.DefaultQuery("current","1") //当前页
	size:=c.DefaultQuery("size","5")
	return gin.H{
		"code":200,
		"data":this.RoleService.ListClusterRoleBindings(this.Helper.StrToInt(page,1),this.Helper.StrToInt(size,5)),
	}
}


//UserAccount 列表
func(this *RBACCtl) UaList(c *gin.Context) goft.Json{
  //写死的路径存储证书
	keyReg:=regexp.MustCompile(".*_key.pem")
	users:=make([]*UAModel,0)
	suffix:=".pem"
	err:= filepath.Walk(helper.UserPath, func(p string, f os.FileInfo, err error) error {
		if f.IsDir() {return nil}
		if path.Ext(f.Name())==suffix{
			if !keyReg.MatchString(f.Name()){
				users=append(users,&UAModel{
					Name:strings.Replace(f.Name(),suffix,"",-1),
					CreateTime:f.ModTime().Format("2006-01-02 15:04:05"),
				})
			}
		}
		return nil
	})
	goft.Error(err)
	return gin.H{
		"code":200,
		"data":users,
	}

}
func(this *RBACCtl) PostUa(c *gin.Context) goft.Json{
	postModel:=&PostUAModel{}
	goft.Error(c.ShouldBindJSON(postModel))
	goft.Error(this.Helper.GenK8sUser(postModel.CN,postModel.O))
	return gin.H{
		"code":200,
		"data":"success",
	}

}
func(this *RBACCtl) DeleteUa(c *gin.Context) goft.Json{
	user:=c.DefaultQuery("user","")
	goft.Error(this.Helper.DeleteK8sUser(user))
	return gin.H{
		"code":200,
		"data":"success",
	}

}
//生成 config文件
func(this *RBACCtl) Clientconfig(c *gin.Context) goft.Json{
	user:=c.DefaultQuery("user","")
	API:=c.DefaultQuery("api","https://127.0.0.1:6443")
	if user==""{
		panic("no such user")
	}
	cfg:=api.NewConfig()
	clusterName:="kubernetes"
	cfg.Clusters[clusterName]=&api.Cluster{
		Server:API,
		CertificateAuthorityData:this.Helper.CertData(helper.CaCrtPath),
	}
	contextName:=fmt.Sprintf("%s@kubernetes",user)
	cfg.Contexts[contextName]=&api.Context{
		AuthInfo:user,
		Cluster:clusterName,
	}
	cfg.CurrentContext=contextName
	userCertFile:=fmt.Sprintf("%s/%s.pem", helper.UserPath,user)
	userCertKeyFile:=fmt.Sprintf("%s/%s_key.pem", helper.UserPath,user)
	cfg.AuthInfos[user]=&api.AuthInfo{
		ClientKeyData:this.Helper.CertData(userCertKeyFile),
		ClientCertificateData:this.Helper.CertData(userCertFile),
	}
	fileContent,err:=clientcmd.Write(*cfg)
	goft.Error(err)

	return gin.H{
		"code":200,
		"data":string(fileContent),
	}
}
func(this *RBACCtl)  Build(goft *goft.Goft){
	goft.Handle("GET","/rolesAll",this.RolesAll)
	goft.Handle("GET","/roles",this.Roles)
	goft.Handle("POST","/roles",this.CreateRole)
	goft.Handle("DELETE","/roles",this.DeleteRole)
	goft.Handle("PUT","/roles",this.UpdateRolesDetail) //修改角色
	goft.Handle("GET","/GetRoles",this.RolesDetail)

	goft.Handle("POST","/roleBindings",this.CreateRoleBinding)
	goft.Handle("DELETE","/roleBindings",this.DeleteRoleBinding)
	goft.Handle("GET","/roleBindings",this.RoleBindingList)


	goft.Handle("GET","/clusterRolesAll",this.clusterRolesAll)
	goft.Handle("GET","/clusterRoles",this.clusterRoles)
	goft.Handle("POST","/clusterRoles",this.CreateClusterRoles)
	goft.Handle("DELETE","/clusterRoles",this.DeleteClusterRoles)
	goft.Handle("GET","/GetClusterRoles",this.ClusterRolesDetail)
	goft.Handle("PUT","/clusterRoles",this.UpdateClusterRolesDetail) //修改集群角色

	goft.Handle("PUT","/roleBindings",this.AddUserToRoleBinding) //添加用户到binding



	goft.Handle("GET","/clusterRoleBindings",this.ClusterRoleBindingList)
	goft.Handle("POST","/clusterRoleBindings",this.CreateClusterRoleBinding)
	goft.Handle("PUT","/clusterRoleBindings",this.AddUserToClusterRoleBinding)
	goft.Handle("DELETE","/clusterRoleBindings",this.DeleteClusterRoleBinding)


	goft.Handle("GET","/ua",this.UaList)
	goft.Handle("POST","/ua",this.PostUa)
	goft.Handle("DELETE","/ua",this.DeleteUa)

	goft.Handle("GET","/clientConfig",this.Clientconfig)

}
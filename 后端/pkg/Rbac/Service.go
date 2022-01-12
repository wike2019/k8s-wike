package Rbac

import (
	"github.com/gin-gonic/gin"
	"k8s.io/api/rbac/v1"
	"k8sapi/pkg/Common"
)

//@Service
type RoleService struct {
   RoleMap *RoleMapStruct  `inject:"-"`
   RoleBindingMap *RoleBindingMapStruct  `inject:"-"`
   ClusterRoleMap *ClusterRoleMapStruct  `inject:"-"`
   Common.Helper
   ClusterRoleBindingMap *ClusterRoleBindingMapStruct  `inject:"-"`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

//对外方法，根据分页取dep列表
func(this *RoleService) PageDeps(ns string,page,size int ) *Common.ItemsPage{
	List:=this.RoleMap.ListAll(ns)

	ret:=make([]*RoleModel,len(List))
	for i,item:=range List{
		ret[i]=&RoleModel{
			Name:item.Name,
			CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:item.Namespace,
		}
	}
	Covert:=make([]interface{},len(List))
	for i,dep:=range ret{
		Covert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		size, //尺寸默认5
		Covert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":ns, //命名空间
	})
}

//对外方法，根据分页取dep列表
func(this *RoleService) rolesAll(ns string)[]*RoleModel{
	List:=this.RoleMap.ListAll(ns)

	ret:=make([]*RoleModel,len(List))
	for i,item:=range List{
		ret[i]=&RoleModel{
			Name:item.Name,
			CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:item.Namespace,

		}
	}
	return  ret
}


func(this *RoleService) BindPageDeps(ns string,page,size int) *Common.ItemsPage {
	List:=this.RoleBindingMap.ListAll(ns)
	ret:=make([]*RoleBindingModel,len(List))
	for i,item:=range List{
		ret[i]=&RoleBindingModel{
			Name:item.Name,
			CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			NameSpace:item.Namespace,
			Subject:item.Subjects,
			RoleRef:item.RoleRef,

		}
	}
	Covert:=make([]interface{},len(List))
	for i,dep:=range ret{
		Covert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		size, //尺寸默认5
		Covert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":ns, //命名空间
	})

}

func(this *RoleService) GetRoleBinding(ns ,name string) *v1.RoleBinding{
	rb:= this.RoleBindingMap.Get(ns,name)
	if rb==nil{
		panic("no such rolebinding")
	}
	return rb
}

func(this *RoleService) GetRole(ns ,name string) *v1.Role{
	rb:= this.RoleMap.Get(ns,name)
	if rb==nil{
		panic("no such role")
	}
	return rb
}

func(this *RoleService) GetClusterRole(name string) *v1.ClusterRole{
	rb:= this.ClusterRoleMap.Get(name)
	if rb==nil{
		panic("no such GetClusterRole")
	}
	return rb
}

func(this *RoleService) ListClusterRoles(page,size int) *Common.ItemsPage {
	List:=this.ClusterRoleMap.ListAll()
	ret:=make([]*ClusterRoleModel,len(List))
	for i,item:=range List{
		ret[i]=&ClusterRoleModel{
			Name:item.Name,
			CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
	}
	Covert:=make([]interface{},len(List))
	for i,dep:=range ret{
		Covert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		size, //尺寸默认5
		Covert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":"__ClusterRole__", //命名空间
	})
}

//对外方法，根据分页取dep列表
func(this *RoleService)  ALlListClusterRoles()[]*ClusterRoleModel{
	List:=this.ClusterRoleMap.ListAll()
	ret:=make([]*ClusterRoleModel,len(List))
	for i,item:=range List{
		ret[i]=&ClusterRoleModel{
			Name:item.Name,
			CreateTime:item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		}
	}
	return  ret
}

func(this *RoleService) GetClusterRoleBinding(name string) *v1.ClusterRoleBinding{
	crb:= this.ClusterRoleBindingMap.Get(name)
	if crb==nil{
		panic("no such clusterrolebinding")
	}
	return crb
}

func(this *RoleService) ListClusterRoleBindings(page,size int) *Common.ItemsPage {
	List:= this.ClusterRoleBindingMap.ListAll()
	Covert:=make([]interface{},len(List))
	for i,dep:=range List{
		Covert[i]=dep //这里主要作用就是转换类型 把[]*models.Pod 类型转化成 []interface{}
	}
	return this.Helper.PageResource(
		page, //当前页
		size, //尺寸默认5
		Covert,//全部数据集合，会自动筛选出分页里面的内容
	).SetExt(gin.H{
		"ns":"__clusterRoleBinding__", //命名空间
	})
}
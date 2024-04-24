package namespace

import (
	"github.com/gin-gonic/gin"
	coreHttp "github.com/wike2019/wike_go/http"
	"github.com/wike2019/wike_go/lib/controller"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"sort"
)

type NsService struct {
	controller.Controller
	Ctl    *NsCtl
	Client *kubernetes.Clientset
}

func NewService(ctl *NsCtl, client *kubernetes.Clientset) *NsService {
	return &NsService{
		Ctl:    ctl,
		Client: client,
	}
}
func (this *NsService) List() []*corev1.Namespace {
	list := this.Ctl.Content.Values()
	sort.Slice(list, func(i, j int) bool {
		if list[i].CreationTimestamp.Time.Equal(list[j].CreationTimestamp.Time) {
			return list[i].Name < list[j].Name
		}
		return list[i].CreationTimestamp.Time.After(list[j].CreationTimestamp.Time)
	})
	return list
}

func (this *NsService) get(c *gin.Context) {
	this.Success(c, "获取数据成功", this.List())
}

func (this *NsService) delete(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	if name == "" {
		controller.Error(400, "名称不能为空")
	}
	err := this.Client.CoreV1().Namespaces().Delete(c, name, metav1.DeleteOptions{})
	controller.Error(500, err.Error())
	this.Success(c, "删除成功", nil)
}

func (this *NsService) create(c *gin.Context) {
	obj := &corev1.Namespace{}
	err := c.ShouldBindJSON(obj)
	controller.Error(400, err.Error())
	if obj.Name == "" {
		controller.Error(400, "名称不能为空")
	}
	_, err = this.Client.CoreV1().Namespaces().Create(c, obj, metav1.CreateOptions{})
	controller.Error(500, err.Error())
	this.Success(c, "创建成功", nil)
}

func (r *NsService) Path() string {
	return "/api/v1"
}

func (this *NsService) Build(r *gin.RouterGroup, GCore *coreHttp.GCore) {
	GCore.DelWithRbac(r, "admin", "/ns", this.delete)
	GCore.GetWithRbac(r, "guest", "/ns", this.get)
	GCore.PostWithRbac(r, "guest", "/ns", this.create)
}

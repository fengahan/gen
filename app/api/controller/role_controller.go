package controller

import (
	"gen/service"
	"github.com/gin-gonic/gin"
)

//@Controller
//@RequestMap(path="/role",method="get")
type RoleController struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) *RoleController {
	return &RoleController{RoleService: roleService}
}

//@RequestMap(method="get",path="getList")
func (r *RoleController) GetList(c *gin.Context) {
	c.String(200, "111")
	return
}

//@RequestMap(method="get",path="hello")
func (r *RoleController) Hello(c *gin.Context) {

	//读取配置
	c.String(200, "hello")
	return
}

func (r *RoleController) Hello3(c *gin.Context) {
	c.String(200, c.Query("name"))
	return
}

//The default path is named UPDATE
//@RequestMap(method="post")
func (r *RoleController) Update(c *gin.Context) {

	return

}

//!!!!It will not be considered part of Route
func (r *RoleController) Create(c *gin.Context) {

	return

}

func Name(c *gin.Context) {

}

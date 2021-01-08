package controller

import "github.com/gin-gonic/gin"

//@Controller
//@RequestMap(path="/role",method="get")
type RoleController struct {
	Service int
}

//@RequestMap(method="get",path="getList")
func (r *RoleController) GetList(c *gin.Context) {
	c.JSON(200, "111")
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

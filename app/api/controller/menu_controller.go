package controller

import (
	"gen/service"
	"github.com/gin-gonic/gin"
)

//@Controller
//@RequestMap(path="/menu",method="get")
type MenuController struct {
	service service.MenuService
}

/***
Wire 容器提供者
*/
func NewMenuController(menuService service.MenuService) *MenuController {
	return &MenuController{service: menuService}
}

//@RequestMap(method="post",path="create")
func (m *MenuController) Create(c *gin.Context) {

	return

}

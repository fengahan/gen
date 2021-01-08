package controller

import "github.com/gin-gonic/gin"

//@Controller
//@RequestMap(path="/menu",method="get")
type MenuController struct {
	Service int
}

//@RequestMap(method="post",path="create")
func (m *MenuController) Create(c *gin.Context) {

	return

}

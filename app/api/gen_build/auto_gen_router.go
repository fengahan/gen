//!!!!!!This code is automatically generated for AmountRoute. Please do not change it
package gen_build

import (
	"gen/app/api/controller"

	"github.com/gin-gonic/gin"
)

func AmountRoute(router *gin.Engine) *gin.Engine {

	var menuController = controller.MenuController{}
	menuGroup := router.Group("/menu")
	{

		menuGroup.POST("create", menuController.Create)

	}

	var roleController = controller.RoleController{}
	roleGroup := router.Group("/role")
	{

		roleGroup.GET("getList", roleController.GetList)

		roleGroup.GET("hello", roleController.Hello)

		roleGroup.GET("hello2", roleController.Hello2)

		roleGroup.POST("Update", roleController.Update)

	}

	return router
}

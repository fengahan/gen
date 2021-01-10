//+build wireinject

package gen_build

import (
	"gen/app/api/controller"
	"gen/service"
	"github.com/google/wire"
)

var SetRoleController = wire.NewSet(controller.NewRoleController, service.NewRoleService)

func InitRoleController() *controller.RoleController {
	wire.Build(SetRoleController)
	return &controller.RoleController{}
}

var SetMenuController = wire.NewSet(controller.NewMenuController, service.NewMenuService)

func InitMenuController() *controller.MenuController {
	wire.Build(SetMenuController)
	return &controller.MenuController{}
}

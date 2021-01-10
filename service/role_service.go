package service

type RoleService struct {
	
}

func NewRoleService() RoleService  {
	return RoleService{}
}

func (rs RoleService)FindName()string {
	return "roleService"
}
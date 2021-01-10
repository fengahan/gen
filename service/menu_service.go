package service

type MenuService struct {

}

func NewMenuService() MenuService  {
	return MenuService{}
}

func (rs MenuService)FindName()string {
	return "roleService"
}
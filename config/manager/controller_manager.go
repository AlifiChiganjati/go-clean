package manager

import "github.com/AlifiChiganjati/go-clean/internal/user/controller"

type ControlManager interface {
	UserController() controller.UserController
}

type userControlManager struct {
	useCaseManager UseCaseManager
}

func NewControlManager(useCaseManager UseCaseManager) ControlManager {
	return &userControlManager{useCaseManager: useCaseManager}
}

func (c *userControlManager) UserController() controller.UserController {
	return *controller.NewUserController(c.useCaseManager.UserUseCase())
}

package v1

import (
	"github.com/AlifiChiganjati/go-clean/internal/user/controller"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	uc controller.UserController
	rg *gin.RouterGroup
}

func NewUserRoutes(ac controller.UserController, rg *gin.RouterGroup) *UserRoutes {
	return &UserRoutes{
		uc: ac,
		rg: rg,
	}
}

func (ur *UserRoutes) Route() {
	user := ur.rg.Group("/user")
	user.POST("/register", ur.uc.RegisterHandler)
	user.POST("/login", ur.uc.LoginHandler)
}

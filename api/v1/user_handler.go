package v1

import (
	"github.com/AlifiChiganjati/go-clean/internal/user/controller"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	uc controller.UserController
	rg *gin.RouterGroup
}

func NewUserRoutes(uc controller.UserController, rg *gin.RouterGroup) *UserRoutes {
	return &UserRoutes{
		uc: uc,
		rg: rg,
	}
}

func (ur *UserRoutes) Route() {
	user := ur.rg.Group("/user")
	user.POST("/register", ur.uc.RegisterNewUser)
}

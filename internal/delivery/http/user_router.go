package http

import (
	"github.com/AlifiChiganjati/go-clean/internal/delivery/middleware"
	"github.com/AlifiChiganjati/go-clean/internal/user/handler"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	uh             handler.UserHandler
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func NewUserRouter(uh handler.UserHandler, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *UserRouter {
	return &UserRouter{
		uh:             uh,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}

func (ur *UserRouter) Route() {
	user := ur.rg.Group("user")
	user.Use(ur.authMiddleware.RequireToken())
	user.GET("/profile", ur.uh.GetHandler)
	user.PUT("/profile", ur.uh.UpdateNameHandler)
}

package http

import (
	"github.com/AlifiChiganjati/go-clean/internal/auth/handler"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	ah handler.AuthHandler
	rg *gin.RouterGroup
}

func NewAuthRouter(ah handler.AuthHandler, rg *gin.RouterGroup) *AuthRouter {
	return &AuthRouter{
		ah: ah,
		rg: rg,
	}
}

func (ar *AuthRouter) Route() {
	auth := ar.rg.Group("/user")
	auth.POST("/register", ar.ah.RegisterHandler)
	auth.POST("/login", ar.ah.LoginHandler)
}

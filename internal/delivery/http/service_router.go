package http

import (
	"github.com/AlifiChiganjati/go-clean/internal/delivery/middleware"
	"github.com/AlifiChiganjati/go-clean/internal/services/handler"
	"github.com/gin-gonic/gin"
)

type (
	ServiceRouter struct {
		sh             handler.ServiceHandler
		rg             *gin.RouterGroup
		authMiddleware middleware.AuthMiddleware
	}
)

func NewServiceRouter(sh handler.ServiceHandler, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *ServiceRouter {
	return &ServiceRouter{
		sh:             sh,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}

func (sr *ServiceRouter) Route() {
	service := sr.rg.Group("/service")
	service.GET("/", sr.sh.GetAllServiceHandler)
}

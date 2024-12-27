package manager

import (
	"github.com/AlifiChiganjati/go-clean/internal/delivery/middleware"
	"github.com/AlifiChiganjati/go-clean/internal/user/handler"
	"github.com/gin-gonic/gin"
)

type (
	HandlerManager interface {
		UserHandler() handler.UserHandler
	}
	handlerManager struct {
		uc             UseCaseManager
		rg             *gin.RouterGroup
		authMiddleware middleware.AuthMiddleware
	}
)

func NewHandlerManager(uc UseCaseManager, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) HandlerManager {
	return &handlerManager{
		uc:             uc,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}

func (hm *handlerManager) UserHandler() handler.UserHandler {
	return handler.NewUserHanlder(hm.uc.UserUseCase(), hm.rg, hm.authMiddleware)
}

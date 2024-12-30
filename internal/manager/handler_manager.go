package manager

import (
	bannerHandler "github.com/AlifiChiganjati/go-clean/internal/banner/handler"
	"github.com/AlifiChiganjati/go-clean/internal/delivery/middleware"
	serviceHandler "github.com/AlifiChiganjati/go-clean/internal/services/handler"
	"github.com/AlifiChiganjati/go-clean/internal/user/handler"
	"github.com/gin-gonic/gin"
)

type (
	HandlerManager interface {
		UserHandler() handler.UserHandler
		ServiceHandler() serviceHandler.ServiceHandler
		BannerHandler() bannerHandler.BannerHandler
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

func (hm *handlerManager) ServiceHandler() serviceHandler.ServiceHandler {
	return serviceHandler.NewUserHanlder(hm.uc.ServiceUseCase(), hm.rg, hm.authMiddleware)
}

func (hm *handlerManager) BannerHandler() bannerHandler.BannerHandler {
	return bannerHandler.NewBannerHandler(hm.uc.BannerUseCase(), hm.rg)
}

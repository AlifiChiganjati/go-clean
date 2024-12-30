package http

import (
	"github.com/AlifiChiganjati/go-clean/internal/banner/handler"
	"github.com/gin-gonic/gin"
)

type (
	BannerRouter struct {
		bh handler.BannerHandler
		rg *gin.RouterGroup
	}
)

func NewBannerRouter(bh handler.BannerHandler, rg *gin.RouterGroup) *BannerRouter {
	return &BannerRouter{
		bh: bh,
		rg: rg,
	}
}

func (br *BannerRouter) Route() {
	banner := br.rg.Group("banner")
	banner.GET("/", br.bh.FindAllHandler)
}

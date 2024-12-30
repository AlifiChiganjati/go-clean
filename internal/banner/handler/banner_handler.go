package handler

import (
	"fmt"
	"net/http"

	"github.com/AlifiChiganjati/go-clean/internal/banner/domain"
	"github.com/AlifiChiganjati/go-clean/internal/banner/usecase"
	"github.com/AlifiChiganjati/go-clean/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	BannerHandler interface {
		FindAllHandler(c *gin.Context)
	}

	bannerHandler struct {
		bc usecase.BannerUseCase
		rg *gin.RouterGroup
	}
)

func NewBannerHandler(bc usecase.BannerUseCase, rg *gin.RouterGroup) BannerHandler {
	return &bannerHandler{
		bc: bc,
		rg: rg,
	}
}

func (bh *bannerHandler) FindAllHandler(c *gin.Context) {
	banners, err := bh.bc.FindAll()
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var rsp []domain.Banner
	for _, banner := range banners {
		rsp = append(rsp, domain.Banner{
			Id: banner.Id,
			BannerResponse: domain.BannerResponse{
				BannerName:  banner.BannerName,
				BannerImg:   banner.BannerImg,
				Description: banner.Description,
			},
		})
		fmt.Printf("Banner Image: %v\n", banner.BannerImg)
	}

	response.SendSingleResponse(c, "Ok", rsp)
}

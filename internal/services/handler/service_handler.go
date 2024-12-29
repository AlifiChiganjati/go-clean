package handler

import (
	"net/http"

	"github.com/AlifiChiganjati/go-clean/internal/delivery/middleware"
	"github.com/AlifiChiganjati/go-clean/internal/services/domain"
	"github.com/AlifiChiganjati/go-clean/internal/services/usecase"
	"github.com/AlifiChiganjati/go-clean/pkg/response"
	"github.com/gin-gonic/gin"
)

type (
	ServiceHandler interface {
		GetAllServiceHandler(c *gin.Context)
	}

	serviceHandler struct {
		sc             usecase.ServiceUseCase
		rg             *gin.RouterGroup
		authMiddleware middleware.AuthMiddleware
	}
)

func NewUserHanlder(sc usecase.ServiceUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) ServiceHandler {
	return &serviceHandler{
		sc:             sc,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}

func (sh *serviceHandler) GetAllServiceHandler(c *gin.Context) {
	services, err := sh.sc.FindAll()
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var rsp []domain.Service
	for _, service := range services {
		rsp = append(rsp, domain.Service{
			Id:           service.Id,
			ServiceCode:  service.ServiceCode,
			ServiceName:  service.ServiceName,
			ServiceIcon:  service.ServiceIcon,
			ServiceTarif: service.ServiceTarif,
			CreatedAt:    service.CreatedAt,
			UpdatedAt:    service.UpdatedAt,
		})
	}
	var newRsp []domain.ServiceResponse
	for _, service := range rsp {
		newRsp = append(newRsp, domain.ServiceResponse{
			ServiceCode:  service.ServiceCode,
			ServiceName:  service.ServiceName,
			ServiceIcon:  service.ServiceIcon,
			ServiceTarif: service.ServiceTarif,
			CreatedAt:    service.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    service.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response.SendSingleResponse(c, "ok", newRsp)
}

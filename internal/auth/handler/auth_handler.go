package handler

import (
	"net/http"

	"github.com/AlifiChiganjati/go-clean/internal/auth/dto"
	"github.com/AlifiChiganjati/go-clean/internal/auth/usecase"
	userDto "github.com/AlifiChiganjati/go-clean/internal/user/dto"
	"github.com/AlifiChiganjati/go-clean/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc usecase.AuthUseCase
}

func NewAuthHandler(uc usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (ah *AuthHandler) RegisterHandler(c *gin.Context) {
	var payload userDto.UserRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	Rsp, err := ah.uc.Register(payload)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	formattedCreatedAt := Rsp.CreatedAt.Format("2006-01-02 15:04:05")
	formattedUpdatedAt := Rsp.UpdatedAt.Format("2006-01-02 15:04:05")
	newRsp := userDto.UserResponseDto{
		Id:        Rsp.Id,
		Email:     Rsp.Email,
		FirstName: Rsp.FirstName,
		LastName:  Rsp.LastName,
		CreatedAt: formattedCreatedAt,
		UpdatedAt: formattedUpdatedAt,
	}

	response.SendCreateResponse(c, "Ok", newRsp)
}

func (ah *AuthHandler) LoginHandler(c *gin.Context) {
	var payload dto.AuthRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newRsp, err := ah.uc.Login(payload)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SendCreateResponse(c, "Ok", newRsp)
}

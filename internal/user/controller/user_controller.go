package controller

import (
	"net/http"

	"github.com/AlifiChiganjati/go-clean/internal/user/domain/dto"
	"github.com/AlifiChiganjati/go-clean/internal/user/usecase"
	"github.com/AlifiChiganjati/go-clean/pkg/helper"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUseCase
}

func NewUserController(uc usecase.UserUseCase) *UserController {
	return &UserController{uc: uc}
}

func (e *UserController) RegisterNewUser(ctx *gin.Context) {
	var payload dto.UserRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	createdUser, err := e.uc.RegisterNewUser(payload)
	if err != nil {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := dto.UserResponseDto{
		Id:        createdUser.Id,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
	}
	helper.SendCreateResponse(ctx, "Success", response)
}

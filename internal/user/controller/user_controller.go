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

func (e *UserController) RegisterHandler(ctx *gin.Context) {
	var payload dto.UserRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helper.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	createdUser, err := e.uc.CreatedUser(payload)
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

func (u *UserController) LoginHandler(c *gin.Context) {
	var payload dto.LoginRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		helper.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	loginData, err := u.uc.Login(payload)
	if err != nil {
		if err.Error() == "1" {
			helper.SendErrorResponse(c, http.StatusForbidden, "Password salah")
			return
		}
		helper.SendErrorResponse(c, http.StatusForbidden, err.Error())
		return
	}
	helper.SendSingleResponse(c, "success", loginData)
}

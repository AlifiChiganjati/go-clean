package controller

import (
	"fmt"
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

func (u *UserController) GetById(ctx *gin.Context) {
	id, exists := ctx.Get("id")
	if !exists {
		helper.SendErrorResponse(ctx, http.StatusUnauthorized, "User ID not found in token")
		return
	}

	userId, ok := id.(string)
	if !ok || userId == "" {
		helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Invalid or empty user ID")
		return
	}

	fmt.Println("qqq", userId)
	user, err := u.uc.GetById(userId)
	if err != nil {
		if err.Error() == fmt.Sprintf("user with ID %s not found", userId) {
			helper.SendErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("User with ID %s not found", userId))
		} else {
			helper.SendErrorResponse(ctx, http.StatusInternalServerError, "Error: "+err.Error())
		}
		return
	}
	response := dto.UserGetByIdResponseDto{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
	}
	helper.SendSingleResponse(ctx, "OK", response)
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/AlifiChiganjati/go-clean/internal/delivery/middleware"
	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"github.com/AlifiChiganjati/go-clean/internal/user/dto"
	"github.com/AlifiChiganjati/go-clean/internal/user/usecase"
	"github.com/AlifiChiganjati/go-clean/pkg/response"
	"github.com/AlifiChiganjati/go-clean/pkg/upload"
	"github.com/gin-gonic/gin"
)

type (
	UserHandler interface {
		GetHandler(c *gin.Context)
		UpdateNameHandler(c *gin.Context)
		UpdateProfileImgHandler(c *gin.Context)
		GetBalanceHandler(c *gin.Context)
	}

	userHandler struct {
		uc             usecase.UserUseCase
		rg             *gin.RouterGroup
		authMiddleware middleware.AuthMiddleware
	}
)

func NewUserHanlder(uc usecase.UserUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) UserHandler {
	return &userHandler{
		uc:             uc,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}

func (uh *userHandler) GetHandler(c *gin.Context) {
	userID, exists := c.Get("user")
	fmt.Println(userID)
	if !exists {
		response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: User not found in context")
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error: Invalid user ID type")
		return
	}

	rsp, err := uh.uc.FindById(userIDStr)
	if err != nil {
		response.SendErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	formattedCreatedAt := rsp.CreatedAt.Format("2006-01-02 15:04:05")
	formattedUpdatedAt := rsp.UpdatedAt.Format("2006-01-02 15:04:05")

	newRsp := domain.UserProfileResponse{
		Email:        rsp.Email,
		FirstName:    rsp.FirstName,
		LastName:     rsp.LastName,
		ProfileImage: rsp.ProfileImage,
		CreatedAt:    formattedCreatedAt,
		UpdatedAt:    formattedUpdatedAt,
	}
	response.SendSingleResponse(c, "ok", newRsp)
}

func (uh *userHandler) UpdateNameHandler(c *gin.Context) {
	userID, exists := c.Get("user")
	if !exists {
		response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: User not found in context")
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error: Invalid user ID type")
		return
	}

	var payload dto.UserUpdateNameDto
	if err := c.BindJSON(&payload); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateUser, err := uh.uc.UpdateNameUser(payload, userIDStr)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	formattedCreatedAt := updateUser.CreatedAt.Format("2006-01-02 15:04:05")
	formattedUpdatedAt := updateUser.UpdatedAt.Format("2006-01-02 15:04:05")

	newRsp := domain.UserProfileResponse{
		Email:        updateUser.Email,
		FirstName:    updateUser.FirstName,
		LastName:     updateUser.LastName,
		ProfileImage: updateUser.ProfileImage,
		CreatedAt:    formattedCreatedAt,
		UpdatedAt:    formattedUpdatedAt,
	}
	response.SendSingleResponse(c, "ok", newRsp)
}

func (uh *userHandler) UpdateProfileImgHandler(c *gin.Context) {
	userID, exists := c.Get("user")
	if !exists {
		response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: User not found in context")
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error: Invalid user ID type")
		return
	}

	userCredential, err := upload.FileImageHandler(c)
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if userCredential.Id != userIDStr {
		response.SendErrorResponse(c, http.StatusForbidden, "You can only update your own profile image")
		return
	}

	updatedUser, err := uh.uc.UpdateProfileImg(userCredential, userIDStr)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	formattedCreatedAt := updatedUser.CreatedAt.Format("2006-01-02 15:04:05")
	formattedUpdatedAt := updatedUser.UpdatedAt.Format("2006-01-02 15:04:05")

	newRsp := domain.UserProfileResponse{
		Email:        updatedUser.Email,
		FirstName:    updatedUser.FirstName,
		LastName:     updatedUser.LastName,
		ProfileImage: updatedUser.ProfileImage,
		CreatedAt:    formattedCreatedAt,
		UpdatedAt:    formattedUpdatedAt,
	}

	response.SendSingleResponse(c, "ok", newRsp)
}

func (uh *userHandler) GetBalanceHandler(c *gin.Context) {
	userID, exists := c.Get("user")
	fmt.Println(userID)
	if !exists {
		response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: User not found in context")
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		response.SendErrorResponse(c, http.StatusInternalServerError, "Internal Server Error: Invalid user ID type")
		return
	}

	rsp, err := uh.uc.FindById(userIDStr)
	if err != nil {
		response.SendErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	formattedCreatedAt := rsp.CreatedAt.Format("2006-01-02 15:04:05")
	formattedUpdatedAt := rsp.UpdatedAt.Format("2006-01-02 15:04:05")

	newRsp := domain.UserBalanceResponse{
		Saldo:     rsp.Saldo,
		CreatedAt: formattedCreatedAt,
		UpdatedAt: formattedUpdatedAt,
	}

	response.SendSingleResponse(c, "ok", newRsp)
}

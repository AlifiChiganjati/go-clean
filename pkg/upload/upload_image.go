package upload

import (
	"fmt"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"github.com/AlifiChiganjati/go-clean/pkg/response"
	"github.com/gin-gonic/gin"
)

func FileImageHandler(c *gin.Context) (domain.User, error) {
	userID, exists := c.Get("user")
	if !exists {
		response.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: User not found in context")
		return domain.User{}, fmt.Errorf("user not found in context")
	}

	userIDStr, ok := userID.(string)
	if !ok {
		return domain.User{}, fmt.Errorf("invalid user ID type")
	}

	file, header, err := c.Request.FormFile("profile")
	if err != nil {
		return domain.User{}, err
	}
	defer file.Close()

	ext := [3]string{".jpg", ".jpeg", ".png"}
	if !checkExtension(header.Filename, ext) {
		response.SendErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid file extension. Allowed extensions are %s", ext))
		return domain.User{}, fmt.Errorf("invalid file extension")
	}

	fileName := fmt.Sprintf("%v_user_%s", rand.New(rand.NewSource(time.Now().UnixNano())).Int(), filepath.Ext(header.Filename))
	fileLocation := filepath.Join("assets", "uploads", fileName)

	if err := c.SaveUploadedFile(header, fileLocation); err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, "Failed to save the uploaded file")
		return domain.User{}, err
	}

	var userCredential domain.User
	userCredential.Id = userIDStr
	userCredential.ProfileImage = fileName

	return userCredential, nil
}

func checkExtension(filename string, ext [3]string) bool {
	e := filepath.Ext(filename)
	for _, a := range ext {
		if a == e {
			return true
		}
	}
	return false
}

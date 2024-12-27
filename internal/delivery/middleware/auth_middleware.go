package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AlifiChiganjati/go-clean/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type (
	AuthMiddleware interface {
		RequireToken(id string) gin.HandlerFunc
	}

	authMiddleware struct {
		jwtService jwt.JwtToken
	}

	authHeader struct {
		AuthorizationHeader string `json:"Authorization"`
	}
)

func NewAuthMiddleware(jwtService jwt.JwtToken) AuthMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
	}
}

func (a *authMiddleware) RequireToken(id string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aH authHeader
		if err := c.ShouldBindHeader(&aH); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		tokenString := strings.Replace(aH.AuthorizationHeader, "Bearer ", "", -1)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		claims, err := a.jwtService.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		fmt.Println("ini claims", claims)
		c.Set("user", claims["user_id"])

		c.Next()
	}
}

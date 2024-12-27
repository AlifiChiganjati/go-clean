package middleware

import (
	"net/http"
	"strings"

	"github.com/AlifiChiganjati/go-clean/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type (
	AuthMiddleware interface {
		RequireToken() gin.HandlerFunc
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

func (a *authMiddleware) RequireToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing Authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization token format"})
			return
		}

		claims, err := a.jwtService.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok || userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID in token claims"})
			return
		}

		c.Set("user", userID)
		c.Next()
	}
}

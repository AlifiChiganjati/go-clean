package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"github.com/dgrijalva/jwt-go"
)

type JwtClaim struct {
	jwt.StandardClaims
	DataClaims domain.JwtClaims `json:"data"`
}

var (
	appName          = os.Getenv("APP_NAME")
	jwtSigningMethod = jwt.SigningMethodHS256
	jwtSignatureKey  = []byte(os.Getenv("TOKEN_KEY"))
)

// GenerateTokenJwt creates a new JWT token with claims
func GenerateTokenJwt(id string, expiredAt int64) (string, error) {
	claims := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    appName,
			ExpiresAt: expiredAt,
			IssuedAt:  time.Now().Unix(),
		},
		DataClaims: domain.JwtClaims{
			Id: id,
		},
	}

	token := jwt.NewWithClaims(jwtSigningMethod, claims)
	signedToken, err := token.SignedString(jwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return jwtSignatureKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token parse error: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

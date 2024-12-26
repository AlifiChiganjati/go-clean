package jwt

import (
	"errors"
	"time"

	"github.com/AlifiChiganjati/go-clean/config"
	"github.com/AlifiChiganjati/go-clean/internal/auth/dto"
	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	JwtToken interface {
		GenerateToken(payload domain.User) (dto.AuthResponseDto, error)
		VerifyToken(tokenString string) (jwt.MapClaims, error)
		RefreshToken(oldTokenString string) (dto.AuthResponseDto, error)
	}

	JwtTokenClaims struct {
		jwt.RegisteredClaims
		UserId uuid.UUID `json:"user_id"`
	}

	jwtToken struct {
		cfg config.TokenConfig
	}
)

func NewJwtToken(cfg config.TokenConfig) JwtToken {
	return &jwtToken{cfg: cfg}
}

func (j *jwtToken) GenerateToken(payload domain.User) (dto.AuthResponseDto, error) {
	claims := JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtLifeTime)),
		},
		UserId: payload.Id,
	}

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtNewClaims.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, errors.New("failed to generate token")
	}
	return dto.AuthResponseDto{Token: token}, nil
}

func (j *jwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})
	if err != nil {
		return nil, errors.New("failed to verify token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return nil, errors.New("invalid claim token")
	}

	return claims, nil
}

func (j *jwtToken) RefreshToken(oldTokenString string) (dto.AuthResponseDto, error) {
	token, err := jwt.Parse(oldTokenString, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKey, nil
	})
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return dto.AuthResponseDto{}, errors.New("invalid claim token")
	}

	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return dto.AuthResponseDto{Token: newTokenString}, nil
}

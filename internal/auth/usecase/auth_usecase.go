package usecase

import (
	"github.com/AlifiChiganjati/go-clean/internal/auth/dto"
	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	userDto "github.com/AlifiChiganjati/go-clean/internal/user/dto"
	"github.com/AlifiChiganjati/go-clean/internal/user/usecase"
	"github.com/AlifiChiganjati/go-clean/pkg/jwt"
)

type (
	AuthUseCase interface {
		Register(payload userDto.UserRequestDto) (domain.User, error)
		Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	}

	authUseCase struct {
		uc       usecase.UserUseCase
		jwtToken jwt.JwtToken
	}
)

func NewAuthUseCase(uc usecase.UserUseCase, jwtToken jwt.JwtToken) AuthUseCase {
	return &authUseCase{
		uc:       uc,
		jwtToken: jwtToken,
	}
}

func (a *authUseCase) Register(payload userDto.UserRequestDto) (domain.User, error) {
	return a.uc.RegisterNewUser(payload)
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.uc.FindByEmailPassword(payload.Email, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := a.jwtToken.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

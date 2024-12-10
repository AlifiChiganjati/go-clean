package usecase

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"github.com/AlifiChiganjati/go-clean/internal/user/domain/dto"
	"github.com/AlifiChiganjati/go-clean/internal/user/repository"
	"github.com/AlifiChiganjati/go-clean/pkg/helper"
)

type (
	UserUseCase interface {
		GetById(id string) (domain.User, error)
		CreatedUser(payload dto.UserRequestDto) (domain.User, error)
		Login(payload dto.LoginRequestDto) (dto.LoginResponseDto, error)
	}

	userUseCase struct {
		repo repository.UserRepository
	}
)

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) GetById(id string) (domain.User, error) {
	user, err := u.repo.Get(id)
	fmt.Println("zzz", user)
	if err != nil {
		return domain.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	return user, nil
}

func (u *userUseCase) CreatedUser(payload dto.UserRequestDto) (domain.User, error) {
	hashPassword, err := helper.HashPassword(payload.Password)
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  hashPassword,
		Email:     payload.Email,
	}

	createdUser, err := u.repo.Create(user)
	if err != nil {
		return domain.User{}, err
	}

	return createdUser, nil
}

func (u *userUseCase) Login(payload dto.LoginRequestDto) (dto.LoginResponseDto, error) {
	user, err := u.repo.GetByEmail(payload.Email)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	isValid := helper.CheckPasswordHash(payload.Pass, user.Password)
	if !isValid {
		return dto.LoginResponseDto{}, errors.New("1")
	}

	loginExpDuration, _ := strconv.Atoi(os.Getenv("TOKEN_LIFE_TIME"))
	expiredAt := time.Now().Add(time.Duration(loginExpDuration) * time.Minute).Unix()
	accessToken, err := helper.GenerateTokenJwt(user.Id, expiredAt)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	return dto.LoginResponseDto{
		Token:  accessToken,
		UserId: user.Id,
	}, nil
}

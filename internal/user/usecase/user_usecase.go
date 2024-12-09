package usecase

import (
	"fmt"

	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"github.com/AlifiChiganjati/go-clean/internal/user/domain/dto"
	"github.com/AlifiChiganjati/go-clean/internal/user/repository"
)

type (
	UserUseCase interface {
		GetById(id string) (domain.User, error)
		RegisterNewUser(payload dto.UserRequestDto) (domain.User, error)
	}

	userUseCase struct {
		repo repository.UserRepository
	}
)

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) GetById(id string) (domain.User, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return domain.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	return user, nil
}

func (u *userUseCase) RegisterNewUser(payload dto.UserRequestDto) (domain.User, error) {
	user := domain.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  payload.Password,
		Email:     payload.Email,
	}

	createdUser, err := u.repo.Create(user)
	if err != nil {
		return domain.User{}, err
	}

	return createdUser, nil
}

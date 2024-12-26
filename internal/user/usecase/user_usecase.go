package usecase

import (
	"errors"
	"fmt"

	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"github.com/AlifiChiganjati/go-clean/internal/user/dto"
	"github.com/AlifiChiganjati/go-clean/internal/user/repository"
	"github.com/AlifiChiganjati/go-clean/pkg/encrypt"
)

type (
	UserUseCase interface {
		FindById(id string) (domain.User, error)
		FindByEmailPassword(email string, password string) (domain.User, error)
		RegisterNewUser(payload dto.UserRequestDto) (domain.User, error)
	}

	userUseCase struct {
		repo repository.UserRepository
	}
)

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) FindById(id string) (domain.User, error) {
	user, err := uc.repo.Get(id)
	if err != nil {
		return domain.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (uc *userUseCase) FindByEmailPassword(email string, password string) (domain.User, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return domain.User{}, errors.New("invalid username or password")
	}

	if err := encrypt.ComparePasswordHash(user.Password, password); err != nil {
		return domain.User{}, err
	}

	user.Password = ""
	return user, nil
}

func (uc *userUseCase) RegisterNewUser(payload dto.UserRequestDto) (domain.User, error) {
	newPassword, err := encrypt.GeneratePasswordHash(payload.Password)
	if err != nil {
		return domain.User{}, err
	}
	newUser := dto.UserRequestDto{
		Id:        payload.Id,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  newPassword,
	}

	return uc.repo.Create(newUser)
}

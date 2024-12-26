package manager

import (
	"github.com/AlifiChiganjati/go-clean/internal/user/usecase"
)

type (
	UseCaseManager interface {
		UserUseCase() usecase.UserUseCase
	}

	useCaseManager struct {
		repo RepoManager
	}
)

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

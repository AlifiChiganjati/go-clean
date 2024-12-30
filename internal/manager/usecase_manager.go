package manager

import (
	bannerUseCase "github.com/AlifiChiganjati/go-clean/internal/banner/usecase"
	serviceUseCase "github.com/AlifiChiganjati/go-clean/internal/services/usecase"
	"github.com/AlifiChiganjati/go-clean/internal/user/usecase"
)

type (
	UseCaseManager interface {
		UserUseCase() usecase.UserUseCase
		ServiceUseCase() serviceUseCase.ServiceUseCase
		BannerUseCase() bannerUseCase.BannerUseCase
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

func (u *useCaseManager) ServiceUseCase() serviceUseCase.ServiceUseCase {
	return serviceUseCase.NewServiceUseCase(u.repo.ServiceRepo())
}

func (u *useCaseManager) BannerUseCase() bannerUseCase.BannerUseCase {
	return bannerUseCase.NewBannerRepository(u.repo.BannerRepo())
}

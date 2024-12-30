package manager

import (
	bannerRepo "github.com/AlifiChiganjati/go-clean/internal/banner/repository"
	servicesRepo "github.com/AlifiChiganjati/go-clean/internal/services/repository"
	"github.com/AlifiChiganjati/go-clean/internal/user/repository"
)

type (
	RepoManager interface {
		UserRepo() repository.UserRepository
		ServiceRepo() servicesRepo.ServiceRepository
		BannerRepo() bannerRepo.BannerRepository
	}

	repoManager struct {
		infra InfraManager
	}
)

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) ServiceRepo() servicesRepo.ServiceRepository {
	return servicesRepo.NewServiceRepository(r.infra.Conn())
}

func (r *repoManager) BannerRepo() bannerRepo.BannerRepository {
	return bannerRepo.NewBannerRepository(r.infra.Conn())
}

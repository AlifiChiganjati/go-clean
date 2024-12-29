package manager

import (
	servicesRepo "github.com/AlifiChiganjati/go-clean/internal/services/repository"
	"github.com/AlifiChiganjati/go-clean/internal/user/repository"
)

type (
	RepoManager interface {
		UserRepo() repository.UserRepository
		ServiceRepo() servicesRepo.ServiceRepository
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

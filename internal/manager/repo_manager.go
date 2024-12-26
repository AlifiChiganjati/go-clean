package manager

import "github.com/AlifiChiganjati/go-clean/internal/user/repository"

type (
	RepoManager interface {
		UserRepo() repository.UserRepository
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

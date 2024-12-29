package usecase

import (
	"github.com/AlifiChiganjati/go-clean/internal/services/domain"
	"github.com/AlifiChiganjati/go-clean/internal/services/repository"
)

type (
	ServiceUseCase interface {
		FindAll() ([]domain.Service, error)
	}

	serviceUseCase struct {
		repo repository.ServiceRepository
	}
)

func NewServiceUseCase(repo repository.ServiceRepository) ServiceUseCase {
	return &serviceUseCase{repo: repo}
}

func (sc *serviceUseCase) FindAll() ([]domain.Service, error) {
	services, err := sc.repo.GetAll()
	if err != nil {
		return []domain.Service{}, nil
	}
	return services, nil
}

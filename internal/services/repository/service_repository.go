package repository

import (
	"github.com/AlifiChiganjati/go-clean/internal/services/domain"
	"gorm.io/gorm"
)

type (
	ServiceRepository interface {
		GetAll() ([]domain.Service, error)
	}

	serviceRepository struct {
		db *gorm.DB
	}
)

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (s *serviceRepository) GetAll() ([]domain.Service, error) {
	var services []domain.Service
	if result := s.db.Find(&services); result.Error != nil {
		return []domain.Service{}, result.Error
	}

	return services, nil
}

package repository

import (
	"github.com/AlifiChiganjati/go-clean/internal/banner/domain"
	"gorm.io/gorm"
)

type (
	BannerRepository interface {
		GetAll() ([]domain.Banner, error)
	}

	bannerRepository struct {
		db *gorm.DB
	}
)

func NewBannerRepository(db *gorm.DB) BannerRepository {
	return &bannerRepository{db: db}
}

func (b *bannerRepository) GetAll() ([]domain.Banner, error) {
	var banners []domain.Banner
	if result := b.db.Find(&banners); result.Error != nil {
		return []domain.Banner{}, result.Error
	}
	return banners, nil
}

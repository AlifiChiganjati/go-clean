package usecase

import (
	"github.com/AlifiChiganjati/go-clean/internal/banner/domain"
	"github.com/AlifiChiganjati/go-clean/internal/banner/repository"
)

type (
	BannerUseCase interface {
		FindAll() ([]domain.Banner, error)
	}

	bannerUseCase struct {
		repo repository.BannerRepository
	}
)

func NewBannerRepository(repo repository.BannerRepository) BannerUseCase {
	return &bannerUseCase{repo: repo}
}

func (bc *bannerUseCase) FindAll() ([]domain.Banner, error) {
	banners, err := bc.repo.GetAll()
	if err != nil {
		return []domain.Banner{}, err
	}

	return banners, nil
}

package repository

import (
	"fmt"

	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Get(id string) (domain.User, error)
		Create(user domain.User) (domain.User, error)
		GetByEmail(email string) (domain.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(user domain.User) (domain.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *userRepository) Get(id string) (domain.User, error) {
	var user domain.User
	result := u.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	fmt.Println("ini id", result)
	return user, nil
}

func (u *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	result := u.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return user, nil
}

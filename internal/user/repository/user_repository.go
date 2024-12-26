package repository

import (
	"fmt"

	"github.com/AlifiChiganjati/go-clean/internal/user/domain"
	"github.com/AlifiChiganjati/go-clean/internal/user/dto"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Get(id string) (domain.User, error)
		GetByEmail(email string) (domain.User, error)
		Create(payload dto.UserRequestDto) (domain.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Get(id string) (domain.User, error) {
	var user domain.User
	if result := u.db.First(&user, "id= ?", id); result.Error != nil {
		return domain.User{}, result.Error
	}
	return user, nil
}

func (u *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	if result := u.db.First(&user, "email= ?", email); result.Error != nil {
		return domain.User{}, result.Error
	}

	return user, nil
}

func (u *userRepository) Create(payload dto.UserRequestDto) (domain.User, error) {
	user := domain.User{
		Id:        payload.Id,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  payload.Password,
	}
	if result := u.db.Create(&user); result.Error != nil {
		return domain.User{}, result.Error
	}
	fmt.Println("aaaa", user)
	return user, nil
}

package domain

import (
	"time"
)

type (
	User struct {
		// gorm.Model
		Id           string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
		FirstName    string    `json:"first_name"`
		LastName     string    `json:"last_name"`
		Email        string    `json:"email"`
		Password     string    `json:"password"`
		ProfileImage string    `json:"profile_image"`
		Saldo        float64   `json:"saldo"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	UserProfileResponse struct {
		Email        string `json:"email"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		ProfileImage string `json:"profile_image"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"UpdatedAt"`
	}
)

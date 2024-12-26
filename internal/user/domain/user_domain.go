package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	// gorm.Model
	Id           uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ProfileImage string    `json:"profile_image"`
	Saldo        float64   `json:"saldo"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

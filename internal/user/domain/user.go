package domain

import "time"

type User struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	ProfileImage string    `json:"profile_image"`
	Saldo        float64   `json:"saldo"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type JwtClaims struct {
	Id string `json:"id"`
}

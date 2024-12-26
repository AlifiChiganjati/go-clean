package dto

import (
	"github.com/google/uuid"
)

type (
	UserRequestDto struct {
		Id        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Password  string    `json:"password"`
	}
	UserResponseDto struct {
		Id        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
	}
)

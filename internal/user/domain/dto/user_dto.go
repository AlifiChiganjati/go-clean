package dto

type UserRequestDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserResponseDto struct {
	Id        string `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name`
	Email     string `json:"email"`
}

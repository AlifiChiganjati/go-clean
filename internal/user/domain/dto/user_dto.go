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
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type LoginRequestDto struct {
	Email string `json:"email" binding:"required"`
	Pass  string `json:"password" binding:"required"`
}

type LoginResponseDto struct {
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}

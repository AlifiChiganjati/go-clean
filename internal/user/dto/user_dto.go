package dto

type (
	UserRequestDto struct {
		Id        string `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}
	UserResponseDto struct {
		Id        string `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	UserUpdateNameDto struct {
		Id        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	UserUpdateNameResponseDto struct {
		Id        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		CreatedAt string `json:"created_at"`
		UpdateAt  string `json:"updated_at"`
	}

	UserUpdateProfileImageDto struct {
		ProfileImage string `json:"profile_image"`
	}

	TopUpRequestDto struct {
		TopUpAmount float64 `json:"top_up_amount" binding:"required"`
	}
)

package dto

type (
	AuthRequestDto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AuthResponseDto struct {
		Token string `json:"token"`
	}
)

package domain

type (
	Banner struct {
		Id string `json:"id"`
		BannerResponse
	}

	BannerResponse struct {
		BannerName  string `json:"banner_name"`
		BannerImg   string `json:"banner_image" gorm:"column:banner_image"`
		Description string `json:"description"`
	}
)

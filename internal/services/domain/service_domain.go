package domain

import "time"

type (
	Service struct {
		Id           string    `json:"id"`
		ServiceCode  string    `json:"service_code"`
		ServiceName  string    `json:"service_name"`
		ServiceIcon  string    `json:"service_icon"`
		ServiceTarif float64   `json:"service_tarif"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	ServiceResponse struct {
		ServiceCode  string  `json:"service_code"`
		ServiceName  string  `json:"service_name"`
		ServiceIcon  string  `json:"service_icon"`
		ServiceTarif float64 `json:"service_tarif"`
		CreatedAt    string  `json:"created_at"`
		UpdatedAt    string  `json:"updated_at"`
	}
)

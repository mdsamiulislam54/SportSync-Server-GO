package dto

import "time"

type ZoneResponse struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	TotalCapacity int       `json:"total_capacity"`
	PricePerHour  float64   `json:"price_per_hour"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateZoneResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    ZoneResponse `json:"data"`
}

type GetAllZoneResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []ZoneResponse `json:"data"`
}

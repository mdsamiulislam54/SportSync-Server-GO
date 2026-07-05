package dto

import "time"

type ReservationResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ZoneID       uint      `json:"zone_id"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateReservationResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    ReservationResponse `json:"data"`
}
type UpdateReservationResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    ReservationResponse `json:"data"`
}

type GetAllReservationResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Data    []MyReservationResponse `json:"data"`
}

type MyReservationResponse struct {
	ID           uint                      `json:"id"`
	LicensePlate string                    `json:"license_plate"`
	Status       string                    `json:"status"`
	Zone         MyReservationResponseZone `json:"zone"`
	CreatedAt    time.Time                 `json:"created_at"`
}

type MyReservationResponseZone struct {
	ID   uint      `json:"id"`
	Name string    `json:"name"`
	Type string `json:"type"`
}

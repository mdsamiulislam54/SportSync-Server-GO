package dto

type CreateReservationRequest struct {
	ZoneID       uint   `json:"zone_id" validate:"required,gt=0"`
	LicensePlate string `json:"license_plate" validate:"required,max=15"`
}

type UpdateReservationRequest struct {
	LicensePlate string `json:"license_plate" validate:"omitempty,max=15"`
	Status       string `json:"status" validate:"required,oneof=active completed cancelled"`
}

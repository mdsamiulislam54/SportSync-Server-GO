package reservations

import (
	"fmt"
	"sportsync/internal/reservations/dto"
)

type service struct {
	repo Repository

}

func NewService(repo Repository) *service {
	return &service{
		repo, 
	}
}

func (s *service) ReservationCreate(req *Reservation) (*dto.CreateReservationResponse, error) {
	reservation := &Reservation{
		UserID:       req.UserID,
		ZoneID:       req.ZoneID,
		LicensePlate: req.LicensePlate,
		Status:       "active",
	}
	fmt.Println("reservation service ...................",reservation)
	if err := s.repo.ReservationCreate(reservation); err != nil {
		return nil, err
	}

	return &dto.CreateReservationResponse{
		Success: true,
		Message: "Reservation confirmed successfully",
		Data: dto.ReservationResponse{
			ID:            reservation.ID,
			UserID:        reservation.UserID,
			ZoneID:        reservation.ZoneID,
			LicensePlate:  reservation.LicensePlate,
			Status:        reservation.Status,
			CreatedAt:     reservation.CreatedAt,
			UpdatedAt:     reservation.UpdatedAt,
		},
	}, nil
}

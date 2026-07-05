package reservations

import (
	"errors"
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
	fmt.Println("reservation service ...................", reservation)
	if err := s.repo.ReservationCreate(reservation); err != nil {
		return nil, err
	}

	return &dto.CreateReservationResponse{
		Success: true,
		Message: "Reservation confirmed successfully",
		Data: dto.ReservationResponse{
			ID:           reservation.ID,
			UserID:       reservation.UserID,
			ZoneID:       reservation.ZoneID,
			LicensePlate: reservation.LicensePlate,
			Status:       reservation.Status,
			CreatedAt:    reservation.CreatedAt,
			UpdatedAt:    reservation.UpdatedAt,
		},
	}, nil
}
func (s *service) MyReservation(id uint) (*dto.GetAllReservationResponse, error) {

	reservations, err := s.repo.MyReservations(id)
	if err != nil {
		return nil, err
	}
	response := make([]dto.MyReservationResponse, 0, len(reservations))

	for _, reservation := range reservations {
		response = append(response, dto.MyReservationResponse{
			ID:           reservation.ID,
			LicensePlate: reservation.LicensePlate,
			Status:       reservation.Status,
			CreatedAt:    reservation.CreatedAt,
			Zone: dto.MyReservationResponseZone{
				ID:   reservation.Zone.ID,
				Name: reservation.Zone.Name,
				Type: reservation.Zone.Type,
			},
		})
	}
	return &dto.GetAllReservationResponse{
		Success: true,
		Message: "My reservations retrieved successfully",
		Data:    response,
	}, nil
}

func (s *service) CancelReservation(reservationId uint, userID uint, role string) error {
	reservation, err := s.repo.GetReservationById(reservationId)
	if err != nil {
		return fmt.Errorf("failed to get reservation: %w", err)
	}

	if role == "driver" && reservation.UserID != userID {
		return errors.New("forbidden")
	}

	if err := s.repo.ReservationsCancel(reservation.ID); err != nil {
		return fmt.Errorf("failed to cancel reservation: %w", err)
	}

	return nil

}

func (s *service) GetAllReservation(role string) ([]Reservation, error) {

	if role != "admin" {
		return nil, errors.New("Access dialed")
	}

	allReservation, err := s.repo.GetAllReservation()
	if err != nil {
		return nil, errors.New("Reservation data not found")
	}

	return allReservation, nil
}

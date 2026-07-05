package reservations

import (
	"errors"
	"log"
	"sportsync/internal/parkingzones"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	ReservationCreate(Reservation *Reservation) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r repository) ReservationCreate(reservation *Reservation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		var zones parkingzones.Zone

		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).First(&zones, reservation.ZoneID).Error; err != nil {
			return err
		}

		var count int64
		if err := tx.Model(&Reservation{}).
			Where("zone_id = ? AND status = ?", reservation.ZoneID, "active").
			Count(&count).Error; err != nil {
			return err
		}

		log.Printf("Active Reservations = %d", count)

		if count >= int64(zones.TotalCapacity) {
			return errors.New("parking zone is full")
		}

		if err := tx.Create(reservation).Error; err != nil {
			log.Printf("Create failed: %v", err)
			return err
		}
		return nil
	})
}
func (r repository) myReservations(userId uint) ([]Reservation, error) {
	var reservation []Reservation
	err := r.db.Find(&reservation, userId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Reservation not found")
		}

		return nil, err
	}

	return reservation, nil
}
func (r repository) reservationsCancel(id uint) (*Reservation, error) {
	var reservation Reservation

	if err := r.db.First(&reservation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation not found")
		}
		return nil, err
	}

	reservation.Status = "cancelled"
	if err := r.db.Save(&reservation).Error; err != nil {
		return nil, err
	}

	return &reservation, nil
}

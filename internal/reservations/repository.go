package reservations

import (
	"errors"
	"log"
	"sportsync/internal/parkingzones"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetAllReservation() ([]Reservation, error)
	ReservationCreate(Reservation *Reservation) error
	MyReservations(userId uint) ([]Reservation, error)
	GetReservationById(id uint) (*Reservation, error)
	ReservationsCancel(id uint) error
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
func (r repository) MyReservations(userId uint) ([]Reservation, error) {
	var reservation []Reservation
	err := r.db.
		Where("user_id = ?", userId).
		Preload("Zone", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "type")
		}).
		Find(&reservation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Reservation not found")
		}

		return nil, err
	}

	return reservation, nil
}
func (r repository) ReservationsCancel(id uint) error {
	var reservation Reservation

	if err := r.db.First(&reservation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("reservation not found")
		}
		return err
	}

	reservation.Status = "cancelled"
	if err := r.db.Save(&reservation).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) GetReservationById(id uint) (*Reservation, error) {
	var reservation Reservation

	if err := r.db.First(&reservation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation not found")
		}
		return nil, err
	}

	return &reservation, nil
}

func (r repository) GetAllReservation() ([]Reservation, error) {
	var reservation []Reservation

	err := r.db.
		Preload("Zone").
		Preload("User").
		Find(&reservation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Reservation not found")
		}
	}

	return reservation, nil
}

package parkingzones

import (
	"errors"

	"gorm.io/gorm"
)

type ZoneWithAvailable struct {
	Zone
	AvailableSpots int `gorm:"column:available_spots"`
}
type Repository interface {
	CreateZone(zone *Zone) error
	GetAllZones() ([]ZoneWithAvailable, error)
	GetZoneById(id uint) (*Zone, error)
}

type repository struct {
	db *gorm.DB
}

func NewZoneRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) CreateZone(zone *Zone) error {
	return r.db.Create(&zone).Error
}
func (r *repository) GetAllZones() ([]ZoneWithAvailable, error) {
	var zones []ZoneWithAvailable
	err := r.db.Model(&Zone{}).
		Select(`
		zones.*,
		(total_capacity - (
		SELECT COUNT(*)
		FROM reservations
		WHERE reservations.zone_id = zones.id
		AND reservations.status ='active'
		)) As available_spots
	
	`).Find(&zones).Error

	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (r *repository) GetZoneById(id uint) (*Zone, error) {
	var zone Zone
	err := r.db.First(&zone, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Zone not found")
		}

		return nil, err
	}

	return &zone, nil
}

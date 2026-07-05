package parkingzones

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateZone(zone *Zone) error
	GetAllZones() ([]Zone, error)
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
func (r *repository) GetAllZones() ([]Zone, error) {
	var zone []Zone
	result := r.db.Find(&zone)
	if result.Error != nil {
		return nil, result.Error
	}
	return zone, nil
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

package reservations

import (
	"sportsync/internal/parkingzones"
	"sportsync/internal/user"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID       uint              `gorm:"not null;index" json:"user_id"`
	ZoneID       uint              `gorm:"not null;index" json:"zone_id"`
	LicensePlate string            `gorm:"size:15;not null" json:"license_plate"`
	Status       string            `gorm:"type:varchar(20);default:active;not null" json:"status"`
	Zone         parkingzones.Zone `gorm:"foreignKey:ZoneID"`
	User         user.User         `gorm:"foreignKey:UserID"`
}

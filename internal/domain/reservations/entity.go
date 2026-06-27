package reservations

import (
	"SpotSync/internal/domain/parking_zones"
	"SpotSync/internal/domain/user"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID        uint                      `gorm:"not null" json:"user_id"`
	ParkingZoneID uint                      `gorm:"not null" json:"zone_id"`
	LicensePlate  string                    `gorm:"type:varchar(20);not null" json:"license_plate"`
	Status        string                    `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	User          user.User                 `gorm:"foreignKey:UserID" json:"user"`
	Zone          parking_zones.ParkingZone `gorm:"foreignKey:ParkingZoneID" json:"zone"`
}

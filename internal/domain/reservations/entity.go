package reservations

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID        uint      `gorm:"not null" json:"user_id"`
	ParkingZoneID uint      `gorm:"not null" json:"parking_zone_id"`
	Status        string    `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	StartTime     time.Time `gorm:"not null" json:"start_time"`
	EndTime       time.Time `gorm:"not null" json:"end_time"`
	TotalCost     float64   `gorm:"type:decimal(10,2);not null" json:"total_cost"`
}
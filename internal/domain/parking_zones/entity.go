package parking_zones

import (
	"gorm.io/gorm"
)

type ParkingZone struct {
	gorm.Model
	Name          string  `gorm:"not null" json:"name"`
	Type          string  `gorm:"type:varchar(20);not null" json:"type"`
	TotalCapacity int     `gorm:"not null" json:"total_capacity"`
	PricePerHour  float64 `gorm:"type:decimal(10,2);not null" json:"price_per_hour"`
}

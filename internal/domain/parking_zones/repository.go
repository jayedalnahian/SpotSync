package parking_zones

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorUnknown = errors.New("Something went wrong.")

type Repository interface {
	CreateZone(zone *ParkingZone) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateZone(zone *ParkingZone) error {
	result := r.db.Create(zone)
	if result.Error != nil {
		return ErrorUnknown
	}
	return nil
}

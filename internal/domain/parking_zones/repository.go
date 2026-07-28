package parking_zones

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorUnknown = errors.New("Something went wrong.")

type ParkingZoneWithAvailableSpots struct {
	ParkingZone
	AvailableSpots int
}

type Repository interface {
	CreateZone(zone *ParkingZone) error
	GetAllZones() ([]ParkingZoneWithAvailableSpots, error)
	GetZoneByID(id uint) (*ParkingZoneWithAvailableSpots, error)
	GetRawZoneByID(id uint) (*ParkingZone, error)
	UpdateZone(zone *ParkingZone) error
	DeleteZone(id uint) error
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

func (r *repository) GetAllZones() ([]ParkingZoneWithAvailableSpots, error) {
	var zones []ParkingZoneWithAvailableSpots

	subQuery := r.db.Table("reservations").
		Select("COUNT(id)").
		Where("parking_zone_id = parking_zones.id").
		Where("status = ?", "active")

	result := r.db.Model(&ParkingZone{}).
		Select("parking_zones.*, (parking_zones.total_capacity - COALESCE((?), 0)) AS available_spots", subQuery).
		Find(&zones)

	if result.Error != nil {
		return nil, ErrorUnknown
	}
	return zones, nil
}

func (r *repository) GetZoneByID(id uint) (*ParkingZoneWithAvailableSpots, error) {
	var zone ParkingZoneWithAvailableSpots

	subQuery := r.db.Table("reservations").
		Select("COUNT(id)").
		Where("parking_zone_id = parking_zones.id").
		Where("status = ?", "active")

	result := r.db.Model(&ParkingZone{}).
		Select("parking_zones.*, (parking_zones.total_capacity - COALESCE((?), 0)) AS available_spots", subQuery).
		Where("parking_zones.id = ?", id).
		First(&zone)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("parking zone not found")
		}
		return nil, ErrorUnknown
	}

	return &zone, nil
}

func (r *repository) GetRawZoneByID(id uint) (*ParkingZone, error) {
	var zone ParkingZone
	result := r.db.First(&zone, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("parking zone not found")
		}
		return nil, ErrorUnknown
	}
	return &zone, nil
}

func (r *repository) UpdateZone(zone *ParkingZone) error {
	result := r.db.Save(zone)
	if result.Error != nil {
		return ErrorUnknown
	}
	return nil
}

func (r *repository) DeleteZone(id uint) error {
	result := r.db.Delete(&ParkingZone{}, id)
	if result.Error != nil {
		return ErrorUnknown
	}
	if result.RowsAffected == 0 {
		return errors.New("parking zone not found")
	}
	return nil
}

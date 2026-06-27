package reservations

import (
	"SpotSync/internal/domain/parking_zones"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrZoneFull = errors.New("parking zone is full")
var ErrReservationNotFound = errors.New("reservation not found")
var ErrForbiddenReservationAccess = errors.New("forbidden: you can only cancel your own reservation")

type Repository interface {
	ReserveSpot(userID, zoneID uint, licensePlate string) (*Reservation, error)
	GetMyReservations(userID uint) ([]Reservation, error)
	GetAllReservations() ([]Reservation, error)
	CancelReservation(userID uint, userRole string, reservationID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) ReserveSpot(userID, zoneID uint, licensePlate string) (*Reservation, error) {
	var reservation *Reservation

	err := r.db.Transaction(func(tx *gorm.DB) error {
		var zone parking_zones.ParkingZone

		// 1. Lock the row for update
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&zone, zoneID).Error; err != nil {
			return err
		}

		// 2. Count current 'active' reservations for this zone
		var activeCount int64
		if err := tx.Model(&Reservation{}).Where("parking_zone_id = ? AND status = ?", zoneID, "active").Count(&activeCount).Error; err != nil {
			return err
		}

		// 3. Check if active_count < zone.total_capacity
		if activeCount >= int64(zone.TotalCapacity) {
			return ErrZoneFull
		}

		// 4. Create reservation
		res := &Reservation{
			UserID:        userID,
			ParkingZoneID: zoneID,
			LicensePlate:  licensePlate,
			Status:        "active",
		}
		if err := tx.Create(res).Error; err != nil {
			return err
		}

		reservation = res
		return nil
	})

	if err != nil {
		return nil, err
	}

	return reservation, nil
}
func (r *repository) GetMyReservations(userID uint) ([]Reservation, error) {
	var reservations []Reservation
	if err := r.db.Preload("Zone").Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *repository) GetAllReservations() ([]Reservation, error) {
	var reservations []Reservation
	if err := r.db.Preload("User").Preload("Zone").Order("created_at desc").Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *repository) CancelReservation(userID uint, userRole string, reservationID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var reservation Reservation
		if err := tx.First(&reservation, reservationID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrReservationNotFound
			}
			return err
		}

		if userRole != "admin" && reservation.UserID != userID {
			return ErrForbiddenReservationAccess
		}

		if reservation.Status == "cancelled" {
			return nil
		}

		return tx.Model(&reservation).Update("status", "cancelled").Error
	})
}

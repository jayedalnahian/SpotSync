package reservations

import (
	"SpotSync/internal/domain/reservations/dto"
)

type Service interface {
	ReserveSpot(userID uint, req dto.ReserveSpotRequest) (dto.ReservationResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ReserveSpot(userID uint, req dto.ReserveSpotRequest) (dto.ReservationResponse, error) {
	reservation, err := s.repo.ReserveSpot(userID, req.ZoneID, req.LicensePlate)
	if err != nil {
		return dto.ReservationResponse{}, err
	}

	return dto.ReservationResponse{
		ID:           reservation.ID,
		UserID:       reservation.UserID,
		ZoneID:       reservation.ParkingZoneID,
		LicensePlate: reservation.LicensePlate,
		Status:       reservation.Status,
		CreatedAt:    reservation.CreatedAt,
		UpdatedAt:    reservation.UpdatedAt,
	}, nil
}
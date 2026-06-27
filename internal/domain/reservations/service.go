package reservations

import (
	"SpotSync/internal/domain/reservations/dto"
)

type Service interface {
	ReserveSpot(userID uint, req dto.ReserveSpotRequest) (dto.ReservationResponse, error)
	GetMyReservations(userID uint) ([]dto.MyReservationResponse, error)
	CancelReservation(userID uint, userRole string, reservationID uint) error
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

func (s *service) GetMyReservations(userID uint) ([]dto.MyReservationResponse, error) {
	reservations, err := s.repo.GetMyReservations(userID)
	if err != nil {
		return nil, err
	}

	var response []dto.MyReservationResponse
	for _, res := range reservations {
		response = append(response, dto.MyReservationResponse{
			ID:           res.ID,
			LicensePlate: res.LicensePlate,
			Status:       res.Status,
			Zone: dto.ZoneResponse{
				ID:   res.Zone.ID,
				Name: res.Zone.Name,
				Type: res.Zone.Type,
			},
			CreatedAt: res.CreatedAt,
		})
	}

	if response == nil {
		response = []dto.MyReservationResponse{}
	}

	return response, nil
}

func (s *service) CancelReservation(userID uint, userRole string, reservationID uint) error {
	return s.repo.CancelReservation(userID, userRole, reservationID)
}

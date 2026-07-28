package parking_zones

import (
	"SpotSync/internal/domain/parking_zones/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateZone(req dto.CreateZoneRequest) (*dto.CreateZoneResponse, error) {
	zone := ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	err := s.repo.CreateZone(&zone)
	if err != nil {
		return nil, err
	}

	return &dto.CreateZoneResponse{
		ID:            zone.ID,
		Name:          zone.Name,
		Type:          zone.Type,
		TotalCapacity: zone.TotalCapacity,
		PricePerHour:  zone.PricePerHour,
		CreatedAt:     zone.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:     zone.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *service) GetAllZones() ([]dto.GetZoneResponse, error) {
	zones, err := s.repo.GetAllZones()
	if err != nil {
		return nil, err
	}

	var res []dto.GetZoneResponse
	for _, zone := range zones {
		res = append(res, dto.GetZoneResponse{
			ID:             zone.ID,
			Name:           zone.Name,
			Type:           zone.Type,
			TotalCapacity:  zone.TotalCapacity,
			AvailableSpots: zone.AvailableSpots,
			PricePerHour:   zone.PricePerHour,
			CreatedAt:      zone.CreatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return res, nil
}

func (s *service) GetZoneByID(id uint) (*dto.GetZoneResponse, error) {
	zone, err := s.repo.GetZoneByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.GetZoneResponse{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.AvailableSpots,
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *service) UpdateZone(id uint, req dto.UpdateZoneRequest) (*dto.CreateZoneResponse, error) {
	zone, err := s.repo.GetRawZoneByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		zone.Name = req.Name
	}
	if req.Type != "" {
		zone.Type = req.Type
	}
	if req.TotalCapacity > 0 {
		zone.TotalCapacity = req.TotalCapacity
	}
	if req.PricePerHour > 0 {
		zone.PricePerHour = req.PricePerHour
	}

	if err := s.repo.UpdateZone(zone); err != nil {
		return nil, err
	}

	return &dto.CreateZoneResponse{
		ID:            zone.ID,
		Name:          zone.Name,
		Type:          zone.Type,
		TotalCapacity: zone.TotalCapacity,
		PricePerHour:  zone.PricePerHour,
		CreatedAt:     zone.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:     zone.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (s *service) DeleteZone(id uint) error {
	return s.repo.DeleteZone(id)
}

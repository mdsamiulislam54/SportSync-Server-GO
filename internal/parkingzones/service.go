package parkingzones

import (
	"sportsync/internal/parkingzones/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo,
	}
}

func (s *service) CreateZones(req *Zone) (*dto.CreateZoneResponse, error) {
	zones := Zone{
		Model:         req.Model,
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.repo.CreateZone(&zones); err != nil {
		return nil, err
	}

	return &dto.CreateZoneResponse{
		Success: true,
		Message: "Parking zone created successfully",
		Data: dto.ZoneResponse{
			ID:            zones.ID,
			Name:          zones.Name,
			Type:          zones.Name,
			TotalCapacity: zones.TotalCapacity,
			PricePerHour:  zones.PricePerHour,
			CreatedAt:     zones.CreatedAt,
			UpdatedAt:     zones.UpdatedAt,
		},
	}, nil

}

func (s *service) GetAllZones() (*dto.GetAllZoneResponse, error) {

	zones, err := s.repo.GetAllZones()
	if err != nil {
		return nil, err
	}

	response := make([]dto.ZoneResponse, 0, len(zones))

	for _, zone := range zones {
		response = append(response, dto.ZoneResponse{
			ID:            zone.ID,
			Name:          zone.Name,
			Type:          zone.Type,
			TotalCapacity: zone.TotalCapacity,
			PricePerHour:  zone.PricePerHour,
			CreatedAt:     zone.CreatedAt,
			UpdatedAt:     zone.UpdatedAt,
		})
	}

	return &dto.GetAllZoneResponse{
		Success: true,
		Message: "Parking zones retrieved  successfully",
		Data:    response,
	}, nil

}

func (s *service) GetZoneById(id uint) (*dto.CreateZoneResponse, error) {

	zone, err := s.repo.GetZoneById(id)

	if err != nil {
		return nil, err
	}

	return &dto.CreateZoneResponse{
		Success: true,
		Message: "Parking zone retrieved  successfully",
		Data: dto.ZoneResponse{
			ID:            zone.ID,
			Name:          zone.Name,
			Type:          zone.Name,
			TotalCapacity: zone.TotalCapacity,
			PricePerHour:  zone.PricePerHour,
			CreatedAt:     zone.CreatedAt,
			UpdatedAt:     zone.UpdatedAt,
		},
	}, nil

}

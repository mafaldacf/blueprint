package train_ticket2

import (
	"context"
)

type AdminTravelService interface {
	AddTravel(ctx context.Context, travelInfo TravelInfo) (Trip, error)
	DeleteTravel(ctx context.Context, tripID string) error
}

type AdminTravelServiceImpl struct {
	travelService TravelService
}

func NewAdminTravelServiceImpl(ctx context.Context, travelService TravelService) (AdminTravelService, error) {
	return &AdminTravelServiceImpl{travelService: travelService}, nil
}

func (a *AdminTravelServiceImpl) AddTravel(ctx context.Context, travelInfo TravelInfo) (Trip, error) {
	return a.travelService.CreateTrip(ctx, travelInfo)
}

func (a *AdminTravelServiceImpl) DeleteTravel(ctx context.Context, tripID string) error {
	return a.travelService.DeleteTrip(ctx, tripID)
}

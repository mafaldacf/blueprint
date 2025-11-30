package trainticket

import (
	"context"
)

type AdminTravelService interface {
	GetAllTravels(ctx context.Context) ([]AdminTrip, error)
	AddTravel(ctx context.Context, travelInfo TravelInfo) (Trip, error)
	UpdateTravel(ctx context.Context, travelInfo TravelInfo) error
	DeleteTravel(ctx context.Context, tripID string) error
}

type AdminTravelServiceImpl struct {
	travelService TravelService
}

func NewAdminTravelServiceImpl(ctx context.Context, travelService TravelService) (AdminTravelService, error) {
	return &AdminTravelServiceImpl{travelService: travelService}, nil
}

func (a *AdminTravelServiceImpl) GetAllTravels(ctx context.Context) ([]AdminTrip, error) {
	return a.travelService.AdminQueryAll(ctx)
}

func (a *AdminTravelServiceImpl) AddTravel(ctx context.Context, travelInfo TravelInfo) (Trip, error) {
	return a.travelService.CreateTrip(ctx, travelInfo)
}

func (a *AdminTravelServiceImpl) UpdateTravel(ctx context.Context, travelInfo TravelInfo) error {
	return a.travelService.UpdateTrip(ctx, travelInfo)
}

func (a *AdminTravelServiceImpl) DeleteTravel(ctx context.Context, tripID string) error {
	return a.travelService.DeleteTrip(ctx, tripID)
}

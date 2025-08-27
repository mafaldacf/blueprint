package train_ticket2

import (
	"context"
)

type AdminBasicInfoService interface {
	DeleteStation(ctx context.Context, id string) error
	DeleteTrain(ctx context.Context, id string) error
}

type AdminBasicInfoServiceImpl struct {
	stationService StationService
	trainService   TrainService
}

func NewAdminBasicInfoServiceImpl(ctx context.Context, stationService StationService, trainService TrainService) (AdminBasicInfoService, error) {
	return &AdminBasicInfoServiceImpl{stationService: stationService, trainService: trainService}, nil
}

func (a *AdminBasicInfoServiceImpl) DeleteStation(ctx context.Context, id string) error {
	return a.stationService.DeleteStation(ctx, id)
}

func (a *AdminBasicInfoServiceImpl) DeleteTrain(ctx context.Context, id string) error {
	_, err := a.trainService.Delete(ctx, id)
	return err
}

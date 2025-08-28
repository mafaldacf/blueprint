package train_ticket2

import (
	"context"
)

type AdminBasicInfoService interface {
	AddStation(ctx context.Context, station Station) error
	AddTrain(ctx context.Context, trainType TrainType) error
	AddConfig(ctx context.Context, config Config) error
	AddContact(ctx context.Context, contact Contact) error
	AddPrice(ctx context.Context, config PriceConfig) error

	DeleteStation(ctx context.Context, id string) error
	DeleteTrain(ctx context.Context, id string) error
	DeleteConfig(ctx context.Context, name string) error
	DeleteContact(ctx context.Context, id string) error
	DeletePrice(ctx context.Context, id string) error
}

type AdminBasicInfoServiceImpl struct {
	stationService  StationService
	trainService    TrainService
	configService   ConfigService
	contactsService ContactsService
	priceService    PriceService
}

func NewAdminBasicInfoServiceImpl(ctx context.Context,
	stationService StationService,
	trainService TrainService,
	configService ConfigService,
	contactsService ContactsService,
	priceService PriceService,
) (AdminBasicInfoService, error) {
	return &AdminBasicInfoServiceImpl{
		stationService:  stationService,
		trainService:    trainService,
		configService:   configService,
		contactsService: contactsService,
		priceService:    priceService,
	}, nil
}

func (a *AdminBasicInfoServiceImpl) AddStation(ctx context.Context, station Station) error {
	return a.stationService.CreateStation(ctx, station)
}

func (a *AdminBasicInfoServiceImpl) AddTrain(ctx context.Context, trainType TrainType) error {
	return a.trainService.Create(ctx, trainType)
}

func (a *AdminBasicInfoServiceImpl) AddConfig(ctx context.Context, config Config) error {
	return a.configService.Create(ctx, config)
}

func (a *AdminBasicInfoServiceImpl) AddContact(ctx context.Context, contact Contact) error {
	return a.contactsService.CreateContacts(ctx, contact)
}

func (a *AdminBasicInfoServiceImpl) AddPrice(ctx context.Context, config PriceConfig) error {
	return a.priceService.CreateNewPriceConfig(ctx, config)
}

func (a *AdminBasicInfoServiceImpl) DeleteStation(ctx context.Context, id string) error {
	return a.stationService.DeleteStation(ctx, id)
}

func (a *AdminBasicInfoServiceImpl) DeleteTrain(ctx context.Context, id string) error {
	_, err := a.trainService.Delete(ctx, id)
	return err
}

func (a *AdminBasicInfoServiceImpl) DeleteConfig(ctx context.Context, name string) error {
	return a.configService.Delete(ctx, name)
}

func (a *AdminBasicInfoServiceImpl) DeleteContact(ctx context.Context, id string) error {
	return a.contactsService.Delete(ctx, id)
}

func (a *AdminBasicInfoServiceImpl) DeletePrice(ctx context.Context, id string) error {
	return a.priceService.DeletePriceConfig(ctx, id)
}

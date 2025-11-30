package trainticket

import (
	"context"
)

type AdminBasicInfoService interface {
	ModifyContact(ctx context.Context, contact Contact) (bool, error)
	GetAllContacts(ctx context.Context) ([]Contact, error)
	DeleteContact(ctx context.Context, id string) error
	AddContact(ctx context.Context, contact Contact) error

	ModifyStation(ctx context.Context, station Station) (bool, error)
	GetAllStations(ctx context.Context) ([]Station, error)
	DeleteStation(ctx context.Context, id string) error
	AddStation(ctx context.Context, station Station) error

	ModifyTrain(ctx context.Context, ttype TrainType) (bool, error)
	GetAllTrains(ctx context.Context) ([]TrainType, error)
	DeleteTrain(ctx context.Context, id string) error
	AddTrain(ctx context.Context, trainType TrainType) error

	ModifyConfig(ctx context.Context, config Config) (bool, error)
	GetAllConfigs(ctx context.Context) ([]Config, error)
	DeleteConfig(ctx context.Context, name string) error
	AddConfig(ctx context.Context, config Config) error

	ModifyPrice(ctx context.Context, config PriceConfig) (bool, error)
	GetAllPrices(ctx context.Context) ([]PriceConfig, error)
	DeletePrice(ctx context.Context, id string) error
	AddPrice(ctx context.Context, config PriceConfig) error
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

func (a *AdminBasicInfoServiceImpl) ModifyContact(ctx context.Context, contact Contact) (bool, error) {
	return a.contactsService.Modify(ctx, contact)
}

func (a *AdminBasicInfoServiceImpl) GetAllContacts(ctx context.Context) ([]Contact, error) {
	return a.contactsService.GetAllContacts(ctx)
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

func (a *AdminBasicInfoServiceImpl) ModifyConfig(ctx context.Context, config Config) (bool, error) {
	return a.configService.Update(ctx, config)
}

func (a *AdminBasicInfoServiceImpl) AddPrice(ctx context.Context, config PriceConfig) error {
	return a.priceService.CreateNewPriceConfig(ctx, config)
}

func (a *AdminBasicInfoServiceImpl) ModifyStation(ctx context.Context, station Station) (bool, error) {
	return a.stationService.UpdateStation(ctx, station)
}

func (a *AdminBasicInfoServiceImpl) GetAllStations(ctx context.Context) ([]Station, error) {
	return a.stationService.FindAll(ctx)
}

func (a *AdminBasicInfoServiceImpl) DeleteStation(ctx context.Context, id string) error {
	return a.stationService.DeleteStation(ctx, id)
}

func (a *AdminBasicInfoServiceImpl) ModifyTrain(ctx context.Context, ttype TrainType) (bool, error) {
	return a.trainService.Update(ctx, ttype)
}

func (a *AdminBasicInfoServiceImpl) GetAllTrains(ctx context.Context) ([]TrainType, error) {
	return a.trainService.AllTrains(ctx)
}

func (a *AdminBasicInfoServiceImpl) DeleteTrain(ctx context.Context, id string) error {
	_, err := a.trainService.Delete(ctx, id)
	return err
}

func (a *AdminBasicInfoServiceImpl) GetAllConfigs(ctx context.Context) ([]Config, error) {
	return a.configService.FindAll(ctx)
}

func (a *AdminBasicInfoServiceImpl) DeleteConfig(ctx context.Context, name string) error {
	return a.configService.Delete(ctx, name)
}

func (a *AdminBasicInfoServiceImpl) DeleteContact(ctx context.Context, id string) error {
	return a.contactsService.Delete(ctx, id)
}

func (a *AdminBasicInfoServiceImpl) ModifyPrice(ctx context.Context, config PriceConfig) (bool, error) {
	return a.priceService.UpdatePriceConfig(ctx, config)
}

func (a *AdminBasicInfoServiceImpl) GetAllPrices(ctx context.Context) ([]PriceConfig, error) {
	return a.priceService.GetAllPriceConfig(ctx)
}

func (a *AdminBasicInfoServiceImpl) DeletePrice(ctx context.Context, id string) error {
	return a.priceService.DeletePriceConfig(ctx, id)
}

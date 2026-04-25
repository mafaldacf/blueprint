package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var adminBasicInfoServiceRegistry = registry.NewServiceRegistry[trainticket.AdminBasicInfoService]("admin_basic_info_service")

func init() {
	adminBasicInfoServiceRegistry.Register("local", func(ctx context.Context) (trainticket.AdminBasicInfoService, error) {
		stationDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		stationService, err := trainticket.NewStationServiceImpl(ctx, stationDB)
		if err != nil {
			return nil, err
		}

		trainDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		trainService, err := trainticket.NewTrainServiceImpl(ctx, trainDB)
		if err != nil {
			return nil, err
		}

		configDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		configService, err := trainticket.NewConfigServiceImpl(ctx, configDB)
		if err != nil {
			return nil, err
		}

		contactsDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		contactsService, err := trainticket.NewContactsServiceImpl(ctx, contactsDB)
		if err != nil {
			return nil, err
		}

		priceDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		priceService, err := trainticket.NewPriceServiceImpl(ctx, priceDB)
		if err != nil {
			return nil, err
		}

		return trainticket.NewAdminBasicInfoServiceImpl(ctx, stationService, trainService, configService, contactsService, priceService)
	})
}

func TestAdminBasicInfoServiceStation(t *testing.T) {
	ctx := context.Background()
	service, err := adminBasicInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	station := trainticket.Station{
		ID:   "abi_sta001",
		Name: "suzhou",
	}
	err = service.AddStation(ctx, station)
	assert.NoError(t, err)

	all, err := service.GetAllStations(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)

	station.Name = "suzhou_updated"
	updated, err := service.ModifyStation(ctx, station)
	assert.NoError(t, err)
	assert.True(t, updated)

	err = service.DeleteStation(ctx, "abi_sta001")
	assert.NoError(t, err)
}

func TestAdminBasicInfoServiceTrain(t *testing.T) {
	ctx := context.Background()
	service, err := adminBasicInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	train := trainticket.TrainType{
		ID:           "abi_train001",
		Name:         "ABI_GaoTie",
		EconomyClass: 2000,
		ComfortClass: 2500,
		AvgSpeed:     350,
	}
	err = service.AddTrain(ctx, train)
	assert.NoError(t, err)

	all, err := service.GetAllTrains(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)

	train.AvgSpeed = 360
	updated, err := service.ModifyTrain(ctx, train)
	assert.NoError(t, err)
	assert.True(t, updated)

	err = service.DeleteTrain(ctx, "abi_train001")
	assert.NoError(t, err)
}

func TestAdminBasicInfoServiceConfig(t *testing.T) {
	ctx := context.Background()
	service, err := adminBasicInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	config := trainticket.Config{
		Name:        "abi_max_capacity",
		Value:       "200",
		Description: "Max seat capacity",
	}
	err = service.AddConfig(ctx, config)
	assert.NoError(t, err)

	all, err := service.GetAllConfigs(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)

	config.Value = "300"
	updated, err := service.ModifyConfig(ctx, config)
	assert.NoError(t, err)
	assert.True(t, updated)

	err = service.DeleteConfig(ctx, "abi_max_capacity")
	assert.NoError(t, err)
}

func TestAdminBasicInfoServiceContact(t *testing.T) {
	ctx := context.Background()
	service, err := adminBasicInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	contact := trainticket.Contact{
		ID:        "abi_contact001",
		AccountID: "abi_acc001",
		Name:      "Alice Admin",
		PhoneNumber: "13900000001",
	}
	err = service.AddContact(ctx, contact)
	assert.NoError(t, err)

	all, err := service.GetAllContacts(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)

	contact.Name = "Alice Admin Updated"
	updated, err := service.ModifyContact(ctx, contact)
	assert.NoError(t, err)
	assert.True(t, updated)

	err = service.DeleteContact(ctx, "abi_contact001")
	assert.NoError(t, err)
}

func TestAdminBasicInfoServicePrice(t *testing.T) {
	ctx := context.Background()
	service, err := adminBasicInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	priceConfig := trainticket.PriceConfig{
		ID:                  "abi_price001",
		TrainType:           "GaoTie",
		RouteID:             "route_abi001",
		BasicPriceRate:      0.5,
		FirstClassPriceRate: 0.8,
	}
	err = service.AddPrice(ctx, priceConfig)
	assert.NoError(t, err)

	all, err := service.GetAllPrices(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)

	priceConfig.BasicPriceRate = 0.6
	updated, err := service.ModifyPrice(ctx, priceConfig)
	assert.NoError(t, err)
	assert.True(t, updated)

	err = service.DeletePrice(ctx, "abi_price001")
	assert.NoError(t, err)
}

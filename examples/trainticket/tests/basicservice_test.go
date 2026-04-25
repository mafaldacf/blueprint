package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var basicServiceRegistry = registry.NewServiceRegistry[trainticket.BasicService]("basic_service")
var basicServiceStationSvc trainticket.StationService
var basicServiceTrainSvc trainticket.TrainService
var basicServiceRouteSvc trainticket.RouteService
var basicServicePriceSvc trainticket.PriceService

func init() {
	basicServiceRegistry.Register("local", func(ctx context.Context) (trainticket.BasicService, error) {
		stationDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		basicServiceStationSvc, err = trainticket.NewStationServiceImpl(ctx, stationDB)
		if err != nil {
			return nil, err
		}

		trainDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		basicServiceTrainSvc, err = trainticket.NewTrainServiceImpl(ctx, trainDB)
		if err != nil {
			return nil, err
		}

		routeDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		basicServiceRouteSvc, err = trainticket.NewRouteServiceImpl(ctx, routeDB)
		if err != nil {
			return nil, err
		}

		priceDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		basicServicePriceSvc, err = trainticket.NewPriceServiceImpl(ctx, priceDB)
		if err != nil {
			return nil, err
		}

		return trainticket.NewBasicServiceImpl(ctx, basicServiceStationSvc, basicServiceTrainSvc, basicServiceRouteSvc, basicServicePriceSvc)
	})
}

func TestBasicServiceQueryForStationId(t *testing.T) {
	ctx := context.Background()
	service, err := basicServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	station := trainticket.Station{
		ID:   "basic_sta001",
		Name: "hangzhou",
	}
	err = basicServiceStationSvc.CreateStation(ctx, station)
	assert.NoError(t, err)

	found, err := service.QueryForStationId(ctx, "basic_sta001")
	assert.NoError(t, err)
	assert.Equal(t, "hangzhou", found.Name)
}

func TestBasicServiceQueryForTravel(t *testing.T) {
	ctx := context.Background()
	service, err := basicServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// set up stations
	err = basicServiceStationSvc.CreateStation(ctx, trainticket.Station{ID: "basic_sta_sh", Name: "shanghai"})
	assert.NoError(t, err)
	err = basicServiceStationSvc.CreateStation(ctx, trainticket.Station{ID: "basic_sta_bj", Name: "beijing"})
	assert.NoError(t, err)

	// set up train type
	err = basicServiceTrainSvc.Create(ctx, trainticket.TrainType{
		ID:           "basic_train001",
		Name:         "GaoTieBasic",
		EconomyClass: 2000,
		ComfortClass: 2500,
		AvgSpeed:     350,
	})
	assert.NoError(t, err)

	// set up route
	routeInfo := trainticket.RouteInfo{
		ID:           "basic_route001",
		StartStation: "shanghai",
		EndStation:   "beijing",
		StationList:  "shanghai,nanjing,beijing",
		DistanceList: "0,300,1200",
	}
	route, err := basicServiceRouteSvc.CreateAndModify(ctx, routeInfo)
	assert.NoError(t, err)

	// seed price config so QueryForTravel can compute prices
	err = basicServicePriceSvc.CreateNewPriceConfig(ctx, trainticket.PriceConfig{
		ID:                  "basic_price001",
		TrainType:           "GaoTieBasic",
		RouteID:             route.ID,
		BasicPriceRate:      0.5,
		FirstClassPriceRate: 0.8,
	})
	assert.NoError(t, err)

	trip := trainticket.Trip{
		TripID:        "G_BASIC_001",
		TrainTypeName: "GaoTieBasic",
		RouteID:       route.ID,
	}
	travel := trainticket.Travel{
		Trip:          trip,
		StartPlace:    "shanghai",
		EndPlace:      "beijing",
		DepartureTime: "2026-05-01",
	}
	result, err := service.QueryForTravel(ctx, travel)
	assert.NoError(t, err)
	assert.Equal(t, "GaoTieBasic", result.TrainType.Name)
}

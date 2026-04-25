package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var travelServiceRegistry = registry.NewServiceRegistry[trainticket.TravelService]("travel_service")

func init() {
	travelServiceRegistry.Register("local", func(ctx context.Context) (trainticket.TravelService, error) {
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

		routeDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		routeService, err := trainticket.NewRouteServiceImpl(ctx, routeDB)
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

		basicService, err := trainticket.NewBasicServiceImpl(ctx, stationService, trainService, routeService, priceService)
		if err != nil {
			return nil, err
		}

		orderDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		orderService, err := trainticket.NewOrderServiceImpl(ctx, orderDB)
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

		seatService, err := trainticket.NewSeatServiceImpl(ctx, orderService, configService)
		if err != nil {
			return nil, err
		}

		travelDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewTravelServiceImpl(ctx, basicService, seatService, routeService, trainService, travelDB)
	})
}

func TestTravelServiceCreateAndRetrieve(t *testing.T) {
	ctx := context.Background()
	service, err := travelServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.TravelInfo{
		TripID:              "G001",
		TrainTypeName:       "GaoTie",
		RouteID:             "route001",
		StartStationName:    "shanghai",
		TerminalStationName: "beijing",
		StartTime:           "08:00:00",
		EndTime:             "13:00:00",
	}
	trip, err := service.CreateTrip(ctx, info)
	assert.NoError(t, err)
	assert.Equal(t, "G001", trip.TripID)

	found, err := service.Retrieve(ctx, "G001")
	assert.NoError(t, err)
	assert.Equal(t, "G001", found.TripID)
	assert.Equal(t, "GaoTie", found.TrainTypeName)
}

func TestTravelServiceUpdateTrip(t *testing.T) {
	ctx := context.Background()
	service, err := travelServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.TravelInfo{
		TripID:              "G002",
		TrainTypeName:       "GaoTie",
		RouteID:             "route002",
		StartStationName:    "nanjing",
		TerminalStationName: "beijing",
		StartTime:           "09:00:00",
		EndTime:             "14:00:00",
	}
	_, err = service.CreateTrip(ctx, info)
	assert.NoError(t, err)

	info.TrainTypeName = "DongChe"
	err = service.UpdateTrip(ctx, info)
	assert.NoError(t, err)

	found, err := service.Retrieve(ctx, "G002")
	assert.NoError(t, err)
	assert.Equal(t, "DongChe", found.TrainTypeName)
}

func TestTravelServiceDeleteTrip(t *testing.T) {
	ctx := context.Background()
	service, err := travelServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.TravelInfo{
		TripID:        "G003",
		TrainTypeName: "GaoTie",
		RouteID:       "route003",
	}
	_, err = service.CreateTrip(ctx, info)
	assert.NoError(t, err)

	err = service.DeleteTrip(ctx, "G003")
	assert.NoError(t, err)

	_, err = service.Retrieve(ctx, "G003")
	assert.Error(t, err)
}

func TestTravelServiceQueryAll(t *testing.T) {
	ctx := context.Background()
	service, err := travelServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.TravelInfo{
		TripID:        "G004",
		TrainTypeName: "GaoTie",
		RouteID:       "route004",
	}
	_, err = service.CreateTrip(ctx, info)
	assert.NoError(t, err)

	all, err := service.QueryAll(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestTravelServiceGetTripsByRouteId(t *testing.T) {
	ctx := context.Background()
	service, err := travelServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.TravelInfo{
		TripID:        "G005",
		TrainTypeName: "GaoTie",
		RouteID:       "route005",
	}
	_, err = service.CreateTrip(ctx, info)
	assert.NoError(t, err)

	trips, err := service.GetTripsByRouteId(ctx, []string{"route005"})
	assert.NoError(t, err)
	assert.NotEmpty(t, trips)
}

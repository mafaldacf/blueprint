package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var adminTravelServiceRegistry = registry.NewServiceRegistry[trainticket.AdminTravelService]("admin_travel_service")

func init() {
	adminTravelServiceRegistry.Register("local", func(ctx context.Context) (trainticket.AdminTravelService, error) {
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
		travelService, err := trainticket.NewTravelServiceImpl(ctx, basicService, seatService, routeService, trainService, travelDB)
		if err != nil {
			return nil, err
		}

		return trainticket.NewAdminTravelServiceImpl(ctx, travelService)
	})
}

func TestAdminTravelServiceAddTravel(t *testing.T) {
	ctx := context.Background()
	service, err := adminTravelServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.TravelInfo{
		TripID:              "ADM_G001",
		TrainTypeName:       "GaoTie",
		RouteID:             "adm_travel_route001",
		StartStationName:    "shanghai",
		TerminalStationName: "beijing",
		StartTime:           "08:00:00",
		EndTime:             "13:00:00",
	}
	trip, err := service.AddTravel(ctx, info)
	assert.NoError(t, err)
	assert.Equal(t, "ADM_G001", trip.TripID)
}

func TestAdminTravelServiceDeleteTravel(t *testing.T) {
	ctx := context.Background()
	service, err := adminTravelServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.TravelInfo{
		TripID:        "ADM_G002",
		TrainTypeName: "GaoTie",
		RouteID:       "adm_travel_route002",
	}
	_, err = service.AddTravel(ctx, info)
	assert.NoError(t, err)

	err = service.DeleteTravel(ctx, "ADM_G002")
	assert.NoError(t, err)
}

func TestAdminTravelServiceUpdateTravel(t *testing.T) {
	ctx := context.Background()
	service, err := adminTravelServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.TravelInfo{
		TripID:        "ADM_G003",
		TrainTypeName: "GaoTie",
		RouteID:       "adm_travel_route003",
		StartTime:     "09:00:00",
		EndTime:       "14:00:00",
	}
	_, err = service.AddTravel(ctx, info)
	assert.NoError(t, err)

	info.TrainTypeName = "DongChe"
	err = service.UpdateTravel(ctx, info)
	assert.NoError(t, err)
}

package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var rebookServiceRegistry = registry.NewServiceRegistry[trainticket.RebookService]("rebook_service")
var rebookOrderDB *simplenosqldb.SimpleNoSQLDB

func init() {
	rebookServiceRegistry.Register("local", func(ctx context.Context) (trainticket.RebookService, error) {
		var err error

		rebookOrderDB, err = simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		orderService, err := trainticket.NewOrderServiceImpl(ctx, rebookOrderDB)
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
		travelDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		travelService, err := trainticket.NewTravelServiceImpl(ctx, basicService, seatService, routeService, trainService, travelDB)
		if err != nil {
			return nil, err
		}

		insidePayDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		payDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		moneyDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		paymentService, err := trainticket.NewPaymentServiceImpl(ctx, payDB, moneyDB)
		if err != nil {
			return nil, err
		}
		insidePaymentService, err := trainticket.NewInsidePaymentServiceImpl(ctx, insidePayDB, paymentService, orderService)
		if err != nil {
			return nil, err
		}

		return trainticket.NewRebookServiceImpl(ctx, seatService, travelService, orderService, trainService, routeService, insidePaymentService)
	})
}

func TestRebookServicePayDifferenceNotPaid(t *testing.T) {
	ctx := context.Background()
	service, err := rebookServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	orderService, err := trainticket.NewOrderServiceImpl(ctx, rebookOrderDB)
	assert.NoError(t, err)

	// status=0 (NOT_PAID) causes PayDifference to return nil immediately
	order := trainticket.Order{
		ID:     "rebook_ord001",
		Status: trainticket.ORDER_STATUS_NOT_PAID,
	}
	_, err = orderService.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	info := trainticket.RebookInfo{
		OrderId:  "rebook_ord001",
		TripId:   "G_REBOOK_001",
		SeatType: 2,
		Date:     "2026-05-01",
	}
	err = service.PayDifference(ctx, info)
	assert.NoError(t, err)
}

func TestRebookServiceRebookOrderNotFound(t *testing.T) {
	ctx := context.Background()
	service, err := rebookServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.RebookInfo{
		OrderId:  "nonexistent_order",
		TripId:   "G_REBOOK_002",
		SeatType: 2,
		Date:     "2026-05-01",
	}
	err = service.Rebook(ctx, info)
	assert.Error(t, err)
}

package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var seatServiceRegistry = registry.NewServiceRegistry[trainticket.SeatService]("seat_service")
var seatServiceConfigSvc trainticket.ConfigService
var seatServiceOrderSvc trainticket.OrderService

func init() {
	seatServiceRegistry.Register("local", func(ctx context.Context) (trainticket.SeatService, error) {
		orderDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		seatServiceOrderSvc, err = trainticket.NewOrderServiceImpl(ctx, orderDB)
		if err != nil {
			return nil, err
		}

		configDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		seatServiceConfigSvc, err = trainticket.NewConfigServiceImpl(ctx, configDB)
		if err != nil {
			return nil, err
		}

		// seed config required by GetLeftTicketOfInterval
		seedCtx := context.Background()
		_ = seatServiceConfigSvc.Create(seedCtx, trainticket.Config{
			Name:  "DirectTicketAllocationProportion",
			Value: "0.6",
		})

		return trainticket.NewSeatServiceImpl(ctx, seatServiceOrderSvc, seatServiceConfigSvc)
	})
}

func TestSeatServiceDistributeSeat(t *testing.T) {
	ctx := context.Background()
	service, err := seatServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// pre-create a sold order so GetTicketListByDateAndTripID returns non-empty
	_, err = seatServiceOrderSvc.CreateNewOrder(ctx, trainticket.Order{
		ID:          "seat_ord001",
		TravelDate:  "2026-05-01",
		TrainNumber: "G_SEAT_001",
		FromStation: "nanjing",
		ToStation:   "beijing",
		Status:      trainticket.ORDER_STATUS_PAID,
	})
	assert.NoError(t, err)

	req := trainticket.SeatRequest{
		TravelDate:   "2026-05-01",
		TrainNumber:  "G_SEAT_001",
		StartStation: "shanghai",
		DestStation:  "beijing",
		SeatType:     2,
		TotalNum:     100,
		Stations:     []string{"shanghai", "nanjing", "beijing"},
	}
	ticket, err := service.DistributeSeat(ctx, req)
	assert.NoError(t, err)
	assert.Greater(t, ticket.SeatNo, 0)
}

func TestSeatServiceGetLeftTicketOfInterval(t *testing.T) {
	ctx := context.Background()
	service, err := seatServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// use non-G/D prefix to skip the order-service lookup path
	req := trainticket.SeatRequest{
		TravelDate:   "2026-05-01",
		TrainNumber:  "K_SEAT_002",
		StartStation: "shanghai",
		DestStation:  "beijing",
		SeatType:     2,
		TotalNum:     100,
		Stations:     []string{"shanghai", "nanjing", "beijing"},
	}
	left, err := service.GetLeftTicketOfInterval(ctx, req)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, left, 0)
}

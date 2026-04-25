package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var orderServiceRegistry = registry.NewServiceRegistry[trainticket.OrderService]("order_service")

func init() {
	orderServiceRegistry.Register("local", func(ctx context.Context) (trainticket.OrderService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewOrderServiceImpl(ctx, db)
	})
}

func TestOrderServiceCreateAndGet(t *testing.T) {
	ctx := context.Background()
	service, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:          "ord001",
		AccountID:   "acc001",
		TrainNumber: "G001",
		FromStation: "shanghai",
		ToStation:   "beijing",
		Status:      trainticket.ORDER_STATUS_NOT_PAID,
		Price:       "100.00",
	}
	created, err := service.CreateNewOrder(ctx, order)
	assert.NoError(t, err)
	assert.Equal(t, "ord001", created.ID)

	found, err := service.GetOrderById(ctx, "ord001")
	assert.NoError(t, err)
	assert.Equal(t, "acc001", found.AccountID)
	assert.Equal(t, "shanghai", found.FromStation)
}

func TestOrderServiceFindAll(t *testing.T) {
	ctx := context.Background()
	service, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:        "ord002",
		AccountID: "acc002",
		Status:    trainticket.ORDER_STATUS_NOT_PAID,
	}
	_, err = service.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	all, err := service.FindAllOrder(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestOrderServiceGetOrderPrice(t *testing.T) {
	ctx := context.Background()
	service, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:    "ord003",
		Price: "250.00",
	}
	_, err = service.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	price, err := service.GetOrderPrice(ctx, "ord003")
	assert.NoError(t, err)
	assert.Equal(t, "250.00", price)
}

func TestOrderServiceModifyOrder(t *testing.T) {
	ctx := context.Background()
	service, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:     "ord004",
		Status: trainticket.ORDER_STATUS_NOT_PAID,
	}
	_, err = service.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	err = service.ModifyOrder(ctx, "ord004", trainticket.ORDER_STATUS_COLLECTED)
	assert.NoError(t, err)
}

func TestOrderServiceDeleteOrder(t *testing.T) {
	ctx := context.Background()
	service, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:     "ord005",
		Status: trainticket.ORDER_STATUS_NOT_PAID,
	}
	_, err = service.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	err = service.DeleteOrder(ctx, "ord005")
	assert.NoError(t, err)
}

func TestOrderServiceCalculateSoldTicket(t *testing.T) {
	ctx := context.Background()
	service, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:          "ord006",
		AccountID:   "acc006",
		TrainNumber: "G002",
		TravelDate:  "2026-04-25",
		SeatClass:   trainticket.ORDER_SEAT_CLASS_BUSINESS,
		Status:      trainticket.ORDER_STATUS_PAID,
	}
	_, err = service.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	sold, err := service.CalculateSoldTicket(ctx, "2026-04-25", "G002")
	assert.NoError(t, err)
	assert.Equal(t, "G002", sold.TrainNumber)
	assert.Equal(t, 1, sold.BusinessSeat)
}

func TestOrderServiceAddCreateNewOrder(t *testing.T) {
	ctx := context.Background()
	service, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:        "ord007",
		AccountID: "acc007",
		Status:    trainticket.ORDER_STATUS_NOT_PAID,
	}
	err = service.AddCreateNewOrder(ctx, order)
	assert.NoError(t, err)

	found, err := service.GetOrderById(ctx, "ord007")
	assert.NoError(t, err)
	assert.Equal(t, "acc007", found.AccountID)
}

func TestOrderServiceGetTicketListByDateAndTripID(t *testing.T) {
	ctx := context.Background()
	service, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:          "ord008",
		TrainNumber: "G003",
		TravelDate:  "2026-05-01",
		FromStation: "nanjing",
		ToStation:   "beijing",
		Status:      trainticket.ORDER_STATUS_PAID,
	}
	_, err = service.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	req := trainticket.SeatRequest{
		TravelDate:  "2026-05-01",
		TrainNumber: "G003",
	}
	ticketInfo, err := service.GetTicketListByDateAndTripID(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, ticketInfo.SoldTickets)
}

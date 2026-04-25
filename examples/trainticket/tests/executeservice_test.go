package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var executeServiceRegistry = registry.NewServiceRegistry[trainticket.ExecuteService]("execute_service")
var executeOrderServiceDB *simplenosqldb.SimpleNoSQLDB

func init() {
	executeServiceRegistry.Register("local", func(ctx context.Context) (trainticket.ExecuteService, error) {
		var err error
		executeOrderServiceDB, err = simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		orderService, err := trainticket.NewOrderServiceImpl(ctx, executeOrderServiceDB)
		if err != nil {
			return nil, err
		}
		return trainticket.NewExecuteServiceImpl(ctx, orderService)
	})
}

func TestExecuteServiceTicketCollect(t *testing.T) {
	ctx := context.Background()
	service, err := executeServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// pre-populate an order with PAID status using the shared DB
	orderService, err := trainticket.NewOrderServiceImpl(ctx, executeOrderServiceDB)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:     "exec_ord001",
		Status: trainticket.ORDER_STATUS_PAID,
	}
	_, err = orderService.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	err = service.TicketCollect(ctx, "exec_ord001")
	assert.NoError(t, err)
}

func TestExecuteServiceTicketExecute(t *testing.T) {
	ctx := context.Background()
	service, err := executeServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// pre-populate an order with COLLECTED status using the shared DB
	orderService, err := trainticket.NewOrderServiceImpl(ctx, executeOrderServiceDB)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:     "exec_ord002",
		Status: trainticket.ORDER_STATUS_COLLECTED,
	}
	_, err = orderService.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	err = service.TicketExecute(ctx, "exec_ord002")
	assert.NoError(t, err)
}

func TestExecuteServiceTicketCollectNotPaid(t *testing.T) {
	ctx := context.Background()
	service, err := executeServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	orderService, err := trainticket.NewOrderServiceImpl(ctx, executeOrderServiceDB)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:     "exec_ord003",
		Status: trainticket.ORDER_STATUS_NOT_PAID,
	}
	_, err = orderService.CreateNewOrder(ctx, order)
	assert.NoError(t, err)

	err = service.TicketCollect(ctx, "exec_ord003")
	assert.Error(t, err)
}

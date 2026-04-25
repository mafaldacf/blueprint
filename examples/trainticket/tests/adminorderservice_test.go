package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var adminOrderServiceRegistry = registry.NewServiceRegistry[trainticket.AdminOrderService]("admin_order_service")

func init() {
	adminOrderServiceRegistry.Register("local", func(ctx context.Context) (trainticket.AdminOrderService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		orderService, err := trainticket.NewOrderServiceImpl(ctx, db)
		if err != nil {
			return nil, err
		}
		return trainticket.NewAdminOrderServiceImpl(ctx, orderService)
	})
}

func TestAdminOrderServiceAddAndGetAll(t *testing.T) {
	ctx := context.Background()
	service, err := adminOrderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:          "adm_ord001",
		AccountID:   "adm_acc001",
		TrainNumber: "G010",
		FromStation: "shanghai",
		ToStation:   "beijing",
		Status:      trainticket.ORDER_STATUS_NOT_PAID,
		Price:       "150.00",
	}
	created, err := service.AddOrder(ctx, order)
	assert.NoError(t, err)
	assert.Equal(t, "adm_ord001", created.ID)

	all, err := service.GetAllOrders(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestAdminOrderServiceDeleteOrder(t *testing.T) {
	ctx := context.Background()
	service, err := adminOrderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	order := trainticket.Order{
		ID:     "adm_ord002",
		Status: trainticket.ORDER_STATUS_NOT_PAID,
	}
	_, err = service.AddOrder(ctx, order)
	assert.NoError(t, err)

	err = service.DeleteOrder(ctx, "adm_ord002")
	assert.NoError(t, err)
}

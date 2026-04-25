package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplequeue"
	"github.com/stretchr/testify/assert"
)

var deliveryServiceRegistry = registry.NewServiceRegistry[trainticket.DeliveryService]("delivery_service")
var deliveryServiceDB *simplenosqldb.SimpleNoSQLDB

func init() {
	deliveryServiceRegistry.Register("local", func(ctx context.Context) (trainticket.DeliveryService, error) {
		queue, err := simplequeue.NewSimpleQueue(ctx)
		if err != nil {
			return nil, err
		}
		var dbErr error
		deliveryServiceDB, dbErr = simplenosqldb.NewSimpleNoSQLDB(ctx)
		if dbErr != nil {
			return nil, dbErr
		}
		return trainticket.NewDeliveryServiceImpl(ctx, queue, deliveryServiceDB)
	})
}

func TestDeliveryServiceFindDeliveryNotFound(t *testing.T) {
	ctx := context.Background()
	service, err := deliveryServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.FindDelivery(ctx, "nonexistent_order")
	assert.Error(t, err)
}

package shopping_app

import (
	"context"
	"sync"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type ShipmentService interface {
	Run(ctx context.Context) error
}

type ShipmentServiceImpl struct {
	order_service  OrderService
	shipment_db    backend.NoSQLDatabase
	shipment_queue backend.Queue
	num_workers    int
}

func NewShipmentServiceImpl(ctx context.Context, order_service OrderService, shipment_db backend.NoSQLDatabase, shipment_queue backend.Queue) (ShipmentService, error) {
	return &ShipmentServiceImpl{order_service: order_service, shipment_db: shipment_db, shipment_queue: shipment_queue, num_workers: 4}, nil
}

func (s *ShipmentServiceImpl) createPendingShipment(ctx context.Context, message ShipmentMessage) error {
	collection, _ := s.shipment_db.GetCollection(ctx, "shipment_database", "shipment_collection")
	shipment := Shipment{
		OrderID: message.OrderID,
		UserID:  message.UserID,
		Status:  "pending",
	}
	collection.InsertOne(ctx, shipment)
	s.order_service.ReadOrder(ctx, shipment.OrderID)
	return nil
}

func (s *ShipmentServiceImpl) workerThread(ctx context.Context) error {
	var forever chan struct{}
	go func() {
		var event map[string]interface{}
		s.shipment_queue.Pop(ctx, &event)
		workerMessage := ShipmentMessage{
			OrderID: event["OrderID"].(string),
			UserID:  event["UserID"].(string),
		}
		s.createPendingShipment(ctx, workerMessage)
	}()
	<-forever
	return nil
}

func (n *ShipmentServiceImpl) Run(ctx context.Context) error {
	backend.GetLogger().Info(ctx, "initializing %d workers", n.num_workers)
	var wg sync.WaitGroup
	wg.Add(n.num_workers)
	for i := 1; i <= n.num_workers; i++ {
		go func(i int) {
			defer wg.Done()
			err := n.workerThread(ctx)
			if err != nil {
				backend.GetLogger().Error(ctx, "error in worker thread: %s", err.Error())
				panic(err)
			}
		}(i)
	}
	wg.Wait()
	backend.GetLogger().Info(ctx, "joining %d workers", n.num_workers)
	return nil
}

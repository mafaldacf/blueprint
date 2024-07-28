package threechain2

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderService interface {
	CreateOrder(ctx context.Context, cartID string, username string, productID string, quantity int, price int) error
	ReadOrder(ctx context.Context, orderID string) (Order, error)
}

type OrderServiceImpl struct {
	stock_service   StockService
	billing_service BillingService
	order_db        backend.NoSQLDatabase
	shipment_queue  backend.Queue
}

func NewOrderServiceImpl(ctx context.Context, stock_service StockService, billing_service BillingService, order_db backend.NoSQLDatabase, shipment_queue backend.Queue) (OrderService, error) {
	return &OrderServiceImpl{stock_service: stock_service, billing_service: billing_service, order_db: order_db, shipment_queue: shipment_queue}, nil
}

func (c *OrderServiceImpl) ReadOrder(ctx context.Context, orderID string) (Order, error) {
	var order Order
	collection, _ := c.order_db.GetCollection(ctx, "order_database", "order_collection")
	query := bson.D{{Key: "orderID", Value: orderID}}
	result, _ := collection.FindOne(ctx, query)
	result.One(ctx, &order)
	return order, nil
}

func (c *OrderServiceImpl) CreateOrder(ctx context.Context, cartID string, username string, productID string, quantity int, price int) error {
	collection, _ := c.order_db.GetCollection(ctx, "order_database", "order_collection")
	order := Order{
		OrderID:   cartID,
		Username:  username,
		ProductID: productID,
		Quantity:  quantity,
		Timestamp: 1,
	}
	collection.InsertOne(ctx, order)

	c.stock_service.ReserveStock(ctx, productID, quantity)
	c.billing_service.CreateBill(ctx, username, productID, quantity, price)

	message := ShipmentMessage{
		OrderID:  order.OrderID,
		Username: username,
	}
	c.shipment_queue.Push(ctx, message)

	return nil
}

package threechain

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type OrderService interface {
	PlaceOrder(ctx context.Context, username string, product string) error
}

type OrderServiceImpl struct {
	cart_db     backend.NoSQLDatabase
	stock_queue backend.Queue
}

func NewOrderServiceImpl(ctx context.Context, cart_db backend.NoSQLDatabase, stock_queue backend.Queue) (OrderService, error) {
	return &OrderServiceImpl{cart_db: cart_db, stock_queue: stock_queue}, nil
}

func (c *OrderServiceImpl) PlaceOrder(ctx context.Context, username string, product string) error {
	collection, _ := c.cart_db.GetCollection(ctx, "order_database", "order_collection")
	order := Order{
		Username:  username,
		Product:   product,
		Timestamp: 1,
	}
	collection.InsertOne(ctx, order)

	return nil
}

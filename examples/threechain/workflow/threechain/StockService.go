package threechain

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type StockService interface {
	ReserveStock(ctx context.Context, username string, product string) error
}

type StockServiceImpl struct {
	cart_db      backend.NoSQLDatabase
	stock_queue  backend.Queue
	orderService OrderService
}

func NewStockServiceImpl(ctx context.Context, cart_db backend.NoSQLDatabase, stock_queue backend.Queue, orderService OrderService) (StockService, error) {
	return &StockServiceImpl{cart_db: cart_db, stock_queue: stock_queue, orderService: orderService}, nil
}

func (s *StockServiceImpl) ReserveStock(ctx context.Context, username string, product string) error {
	collection, _ := s.cart_db.GetCollection(ctx, "stock_database", "stock_collection")
	stock := Stock{
		Product:  product,
		Quantity: 1,
	}
	collection.InsertOne(ctx, stock)
	s.orderService.PlaceOrder(ctx, username, product)

	/* message := PlaceOrderMessage{
		Username: username,
		Product:  product,
	}
	c.stock_queue.Push(ctx, message) */
	return nil
}

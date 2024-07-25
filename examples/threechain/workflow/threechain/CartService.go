package threechain

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type CartService interface {
	CheckoutCart(ctx context.Context, username string, product string) error
}

type CartServiceImpl struct {
	cart_db      backend.NoSQLDatabase
	stock_queue  backend.Queue
	stockService StockService
}

func NewCartServiceImpl(ctx context.Context, cart_db backend.NoSQLDatabase, stock_queue backend.Queue, stockService StockService) (CartService, error) {
	return &CartServiceImpl{cart_db: cart_db, stock_queue: stock_queue, stockService: stockService}, nil
}

func (c *CartServiceImpl) CheckoutCart(ctx context.Context, username string, product string) error {
	collection, _ := c.cart_db.GetCollection(ctx, "cart_database", "cart_collection")
	cart := Cart{
		Username: username,
		Product:  product,
	}
	collection.InsertOne(ctx, cart)
	c.stockService.ReserveStock(ctx, username, product)

	/* message := ReserveStockMessage{
		Username: username,
		Product:  product,
	}
	c.stock_queue.Push(ctx, message) */
	return nil
}

package shopping_app

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type StockService interface {
	ReserveStock(ctx context.Context, productID string, quantity int) error
}

type StockServiceImpl struct {
	stock_db backend.NoSQLDatabase
}

func NewStockServiceImpl(ctx context.Context, stock_db backend.NoSQLDatabase) (StockService, error) {
	return &StockServiceImpl{stock_db: stock_db}, nil
}

func (s *StockServiceImpl) ReserveStock(ctx context.Context, productID string, quantity int) error {
	collection, _ := s.stock_db.GetCollection(ctx, "stock_database", "stock_collection")
	stock := Stock{
		ProductID: productID,
		Quantity:  quantity,
	}
	collection.InsertOne(ctx, stock)
	return nil
}

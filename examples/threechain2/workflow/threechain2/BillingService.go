package threechain2

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type BillingService interface {
	CreateBill(ctx context.Context, username string, productID string, quantity int, price int) error
}

type BillingServiceImpl struct {
	bill_db backend.NoSQLDatabase
}

func NewBillingServiceImpl(ctx context.Context, bill_db backend.NoSQLDatabase) (BillingService, error) {
	return &BillingServiceImpl{bill_db: bill_db}, nil
}

func (s *BillingServiceImpl) CreateBill(ctx context.Context, username string, productID string, quantity int, pricePerUnit int) error {
	collection, _ := s.bill_db.GetCollection(ctx, "bill_database", "bill_collection")
	bill := Bill{
		Username:     username,
		ProductID:    productID,
		Quantity:     quantity,
		PricePerUnit: pricePerUnit,
		TotalCost:    quantity * pricePerUnit,
	}
	collection.InsertOne(ctx, bill)
	return nil
}

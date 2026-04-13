package simpleshop

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type Inventory struct {
	ID     string
	Amount int
}

type InventoryService interface {
	AddInventory(ctx context.Context, id string, amount int) (Inventory, error)
}

type InventoryServiceImpl struct {
	inventoryDB backend.NoSQLDatabase
}

func NewBarServiceImpl(ctx context.Context, inventoryDB backend.NoSQLDatabase) (InventoryService, error) {
	d := &InventoryServiceImpl{inventoryDB: inventoryDB}
	return d, nil
}

func (s *InventoryServiceImpl) AddInventory(ctx context.Context, id string, amount int) (Inventory, error) {
	bar := Inventory{
		ID:     id,
		Amount: amount,
	}

	collection, err := s.inventoryDB.GetCollection(ctx, "inventory_db", "inventory")
	if err != nil {
		return Inventory{}, err
	}

	err = collection.InsertOne(ctx, bar)
	if err != nil {
		return Inventory{}, err
	}

	return bar, nil
}

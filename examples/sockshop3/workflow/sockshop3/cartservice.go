// Package cart implements the SockShop cart microservice.
package sockshop3

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type CartService interface {
	// carts
	GetCart(ctx context.Context, customerID string) ([]Item, error)
	DeleteCart(ctx context.Context, customerID string) error
	MergeCarts(ctx context.Context, customerID, sessionID string) error

	// items
	AddItem(ctx context.Context, customerID string, item Item) (Item, error)
	UpdateItem(ctx context.Context, customerID string, item Item) error
	RemoveItem(ctx context.Context, customerID, itemID string) error
	GetItem(ctx context.Context, customerID string, itemID string) (Item, error)
}
type CartServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewCartServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (CartService, error) {
	return &CartServiceImpl{db: db}, nil
}

func (s *CartServiceImpl) AddItem(ctx context.Context, customerID string, item Item) (Item, error) {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return Item{}, err
	}

	cart, err := s.getCart(ctx, customerID)
	if err != nil {
		return Item{}, err
	}

	if existingItem := findItem(cart, item.ID); existingItem != nil {
		existingItem.Quantity += item.Quantity
		item = *existingItem
	} else {
		cart.Items = append(cart.Items, item)
	}

	collection.Upsert(ctx, bson.D{{Key: "ID", Value: customerID}}, cart)
	return item, nil
}

func (s *CartServiceImpl) DeleteCart(ctx context.Context, customerID string) error {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return nil
	}
	return collection.DeleteMany(ctx, bson.D{{Key: "ID", Value: customerID}})
}

func (s *CartServiceImpl) GetCart(ctx context.Context, customerID string) ([]Item, error) {
	cart, err := s.getCart(ctx, customerID)
	if err != nil {
		return cart.Items, err
	}

	return cart.Items, nil
}

func (s *CartServiceImpl) GetItem(ctx context.Context, customerID string, itemID string) (Item, error) {
	cart, err := s.getCart(ctx, customerID)
	if err != nil {
		return Item{}, err
	}

	if item := findItem(cart, itemID); item != nil {
		return *item, nil
	}

	return Item{}, nil
}

func (s *CartServiceImpl) MergeCarts(ctx context.Context, customerID string, sessionID string) error {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return nil
	}

	sessionCart, err := s.getCart(ctx, sessionID)
	if err != nil {
		return err
	}

	customerCart, err := s.getCart(ctx, customerID)
	if err != nil {
		return err
	}

	if len(sessionCart.Items) == 0 {
		// No update to perform
		return nil
	}

	// Update quantity of existing items; append new items
	customerCartItems := make(map[string]*Item)
	for i := 0; i < len(customerCart.Items); i++ {
		customerCartItems[customerCart.Items[i].ID] = &customerCart.Items[i]
	}

	for _, item := range sessionCart.Items {
		if existing, exists := customerCartItems[item.ID]; exists {
			existing.Quantity += item.Quantity
			existing.UnitPrice = item.UnitPrice
		} else {
			customerCart.Items = append(customerCart.Items, item)
		}
	}
	
	_, err = collection.Upsert(ctx, bson.D{{Key: "ID", Value: customerID}}, customerCart)
	if err != nil {
		return err
	}

	// Only delete the session after successfully merging over to customer
	return collection.DeleteOne(ctx, bson.D{{Key: "ID", Value: sessionID}})
}

func (s *CartServiceImpl) RemoveItem(ctx context.Context, customerID string, itemID string) error {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return nil
	}

	cart, err := s.getCart(ctx, customerID)
	if err != nil {
		return err
	}

	if removed := removeItem(cart, itemID); !removed {
		return nil
	}

	if len(cart.Items) == 0 {
		return s.DeleteCart(ctx, customerID)
	}
	_, err = collection.ReplaceOne(ctx, bson.D{{Key: "ID", Value: customerID}}, cart)
	return err
}

func (s *CartServiceImpl) UpdateItem(ctx context.Context, customerID string, item Item) error {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return nil
	}

	cart, err := s.getCart(ctx, customerID)
	if err != nil {
		return err
	}

	if existing := findItem(cart, item.ID); existing != nil {
		// Item exists in the cart, update the quantity
		existing.Quantity = item.Quantity
		existing.UnitPrice = item.UnitPrice

		// After updating, item quantity is gone, so remove item from cart
		if existing.Quantity <= 0 {
			removeItem(cart, item.ID)
		}

		// If no items left in cart, delete cart
		if len(cart.Items) == 0 {
			return s.DeleteCart(ctx, customerID)
		}
	} else {
		// Item doesn't exist in cart and no items added, so do nothing
		if item.Quantity <= 0 {
			return nil
		}

		// Item needs to be added to cart
		cart.Items = append(cart.Items, item)
	}
	
	_, err = collection.Upsert(ctx, bson.D{{Key: "ID", Value: customerID}}, cart)
	return err
}

func (s *CartServiceImpl) getCart(ctx context.Context, id string) (*Cart, error) {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return nil, nil
	}

	filter := bson.D{{Key: "ID", Value: id}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	var cart Cart
	ok, err := cursor.One(ctx, &cart)
	if !ok {
		return nil, fmt.Errorf("could not get cart for id (%s)", id)
	}
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func findItem(c *Cart, itemID string) *Item {
	if c == nil {
		return nil
	}
	for i := range c.Items {
		if c.Items[i].ID == itemID {
			return &c.Items[i]
		}
	}
	return nil
}

func removeItem(c *Cart, itemID string) bool {
	removed := false
	for i := 0; i < len(c.Items); i++ {
		if c.Items[i].ID == itemID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			i--
			removed = true
		}
	}
	return removed
}

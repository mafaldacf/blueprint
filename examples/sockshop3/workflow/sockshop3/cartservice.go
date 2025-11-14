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

	// get cart
	filter := bson.D{{Key: "ID", Value: customerID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Item{}, err
	}
	var cart Cart
	ok, err := cursor.One(ctx, &cart)
	if !ok {
		return Item{}, fmt.Errorf("could not get cart for id (%s)", customerID)
	}
	if err != nil {
		return Item{}, err
	}

	// find item
	var found bool
	for i := range cart.Items {
		if cart.Items[i].ID == item.ID {
			cart.Items[i].Quantity += item.Quantity
			item = cart.Items[i]
			break
		}
	}

	if !found {
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
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return nil, nil
	}

	// get cart
	filter := bson.D{{Key: "ID", Value: customerID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	var cart Cart
	ok, err := cursor.One(ctx, &cart)
	if !ok {
		return nil, fmt.Errorf("could not get cart for id (%s)", customerID)
	}
	if err != nil {
		return nil, err
	}

	return cart.Items, nil
}

func (s *CartServiceImpl) GetItem(ctx context.Context, customerID string, itemID string) (Item, error) {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return Item{}, nil
	}

	// get cart
	filter := bson.D{{Key: "ID", Value: customerID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Item{}, err
	}
	var cart Cart
	ok, err := cursor.One(ctx, &cart)
	if !ok {
		return Item{}, fmt.Errorf("could not get cart for id (%s)", customerID)
	}
	if err != nil {
		return Item{}, err
	}

	// find item
	for i := range cart.Items {
		if cart.Items[i].ID == itemID {
			return cart.Items[i], nil
		}
	}

	return Item{}, nil
}

func (s *CartServiceImpl) MergeCarts(ctx context.Context, customerID string, sessionID string) error {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return nil
	}

	// get cart (session)
	filter := bson.D{{Key: "ID", Value: sessionID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return err
	}
	var sessionCart Cart
	ok, err := cursor.One(ctx, &sessionCart)
	if !ok {
		return fmt.Errorf("could not get cart for id (%s)", customerID)
	}
	if err != nil {
		return err
	}

	// get cart (customer)
	filter = bson.D{{Key: "ID", Value: sessionID}}
	cursor, err = collection.FindOne(ctx, filter)
	if err != nil {
		return err
	}
	var customerCart Cart
	ok, err = cursor.One(ctx, &customerCart)
	if !ok {
		return fmt.Errorf("could not get cart for id (%s)", customerID)
	}
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

	// get cart
	filter := bson.D{{Key: "ID", Value: customerID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return err
	}
	var cart Cart
	ok, err := cursor.One(ctx, &cart)
	if !ok {
		return fmt.Errorf("could not get cart for id (%s)", customerID)
	}
	if err != nil {
		return err
	}

	// remove item
	removed := false
	for i := 0; i < len(cart.Items); i++ {
		if cart.Items[i].ID == itemID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			i--
			removed = true
		}
	}
	if !removed {
		return nil
	}

	if len(cart.Items) == 0 {
		// delete cart
		return collection.DeleteMany(ctx, bson.D{{Key: "ID", Value: customerID}})
	}
	_, err = collection.ReplaceOne(ctx, bson.D{{Key: "ID", Value: customerID}}, cart)
	return err
}

func (s *CartServiceImpl) UpdateItem(ctx context.Context, customerID string, item Item) error {
	collection, err := s.db.GetCollection(ctx, "cart_db", "carts")
	if err != nil {
		return nil
	}

	// get cart
	filter := bson.D{{Key: "ID", Value: customerID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return err
	}
	var cart Cart
	ok, err := cursor.One(ctx, &cart)
	if !ok {
		return fmt.Errorf("could not get cart for id (%s)", customerID)
	}
	if err != nil {
		return err
	}
	
	// find item
	var found bool
	for i := range cart.Items {
		if cart.Items[i].ID == item.ID {
			// item exists in the cart
			cart.Items[i].Quantity = item.Quantity
			cart.Items[i].UnitPrice = item.UnitPrice

			// TODO
			// After updating, item quantity is gone, so remove item from cart
			/* if cart.Items[i].Quantity < 0 {
				// remove item
				for i := 0; i < len(cart.Items); i++ {
					if cart.Items[i].ID == item.ID {
						cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
						i--
					}
				}
			} */
			
			// If no items left in cart, delete cart
			if len(cart.Items) == 0 {
				// delete cart
				return collection.DeleteMany(ctx, bson.D{{Key: "ID", Value: customerID}})
			}
			found = true
			break
		}
	}

	if !found {
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

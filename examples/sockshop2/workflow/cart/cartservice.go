  // Package cart implements the SockShop cart microservice.
package cart

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type CartService interface {
	// Get all items in a customer's cart.  A customer might not have a cart,
	// in which case the empty list is returned.  customerID can be a userID
	// for a logged in user, or a sessionID for an anonymous user.
	GetCart(ctx context.Context, customerID string) ([]Item, error)

	// Delete a customer's cart
	DeleteCart(ctx context.Context, customerID string) error

	// Merge two carts.  Used when an anonymous customer logs in
	MergeCarts(ctx context.Context, customerID, sessionID string) error

	// Get a specific item from a customer's cart
	GetItem(ctx context.Context, customerID string, itemID string) (Item, error)

	// Add an item to a customer's cart.
	// If the item already exists in the cart, then the total quantity is
	// updated to reflect the combined total.
	// Returns the current state of the item in the customer's cart.
	AddItem(ctx context.Context, customerID string, item Item) (Item, error)

	// Remove an item from the customer's cart
	RemoveItem(ctx context.Context, customerID, itemID string) error

	// Updates an item in the customer's cart to the value provided.
	UpdateItem(ctx context.Context, customerID string, item Item) error
}

// Implementation of [CartService]
type cartImpl struct {
	db backend.NoSQLDatabase
}

// Creates a [CartService] instance that persists cart data in the provided db
func NewCartService(ctx context.Context, db backend.NoSQLDatabase) (CartService, error) {
	return &cartImpl{db: db}, nil
}

// AddItem implements CartService.
func (s *cartImpl) AddItem(ctx context.Context, customerID string, item Item) (Item, error) {
	collection, _ := s.db.GetCollection(ctx, "cart", "carts")

	//dummy logic for testing purposes
	cart := cart{
		ID: customerID,
		Items: []Item{item},
	}

	collection.InsertOne(ctx, cart)
	//collection.Upsert(ctx, bson.D{{"id", customerID}}, cart)
	return item, nil
}

// DeleteCart implements CartService.
func (s *cartImpl) DeleteCart(ctx context.Context, customerID string) error {
	collection, _ := s.db.GetCollection(ctx, "cart", "carts")
	return collection.DeleteMany(ctx, bson.D{{"id", customerID}})
}

// GetCart implements CartService.
func (s *cartImpl) GetCart(ctx context.Context, customerID string) ([]Item, error) {
	cart := &cart{
		ID: customerID,
	}
	return cart.Items, nil
}

// GetItem implements CartService.
func (s *cartImpl) GetItem(ctx context.Context, customerID string, itemID string) (Item, error) {
	return Item{}, nil
}

// MergeCarts implements CartService.
func (s *cartImpl) MergeCarts(ctx context.Context, customerID string, sessionID string) error {
	cart := &cart{
		ID: customerID,
	}
	collection, _ := s.db.GetCollection(ctx, "cart", "carts")
	collection.Upsert(ctx, bson.D{{"id", customerID}}, cart)

	// Only delete the session after successfully merging over to customer
	return collection.DeleteOne(ctx, bson.D{{"id", sessionID}})
}

// RemoveItem implements CartService.
func (s *cartImpl) RemoveItem(ctx context.Context, customerID string, itemID string) error {
	cart := &cart{
		ID: customerID,
	}
	collection, _ := s.db.GetCollection(ctx, "cart", "carts")
	collection.ReplaceOne(ctx, bson.D{{"id", customerID}}, cart)
	return nil
}

// UpdateItem implements CartService.
func (s *cartImpl) UpdateItem(ctx context.Context, customerID string, item Item) error {
	cart := &cart{
		ID: customerID,
		Items: []Item{item},
	}
	collection, _ := s.db.GetCollection(ctx, "cart", "carts")
	collection.Upsert(ctx, bson.D{{"id", customerID}}, cart)
	return nil
}

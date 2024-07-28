package threechain2

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type CartService interface {
	Checkout(ctx context.Context, cartID string) error
}

type CartServiceImpl struct {
	order_service OrderService
	cart_db       backend.NoSQLDatabase
}

func NewCartServiceImpl(ctx context.Context, order_service OrderService, cart_db backend.NoSQLDatabase) (CartService, error) {
	return &CartServiceImpl{order_service: order_service, cart_db: cart_db}, nil
}

func (c *CartServiceImpl) Checkout(ctx context.Context, cartID string) error {
	var cart Cart
	collection, _ := c.cart_db.GetCollection(ctx, "cart_database", "cart_collection")
	query := bson.D{{Key: "cartID", Value: cartID}}
	result, _ := collection.FindOne(ctx, query)
	result.One(ctx, &cart)

	filter := bson.D{{Key: "cartID", Value: cartID}}
	update := bson.D{{Key: "status", Value: "completed"}}
	collection.Upsert(ctx, filter, update)

	c.order_service.CreateOrder(ctx, cartID, cart.Username, cart.ProductID, cart.Quantity, cart.PricePerUnit)
	return nil
}

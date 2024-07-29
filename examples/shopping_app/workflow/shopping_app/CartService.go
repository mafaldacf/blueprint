package shopping_app

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type CartService interface {
	Checkout(ctx context.Context, cartID string) error
	AddProductToCart(ctx context.Context, cartID string, productID string) error
}

type CartServiceImpl struct {
	order_service   OrderService
	product_service ProductService
	cart_db         backend.NoSQLDatabase
}

func NewCartServiceImpl(ctx context.Context, order_service OrderService, product_service ProductService, cart_db backend.NoSQLDatabase) (CartService, error) {
	return &CartServiceImpl{order_service: order_service, product_service: product_service, cart_db: cart_db}, nil
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

	c.order_service.CreateOrder(ctx, cartID, cart.UserID, cart.ProductID, cart.Quantity, cart.PricePerUnit)

	collection.DeleteOne(ctx, filter)

	return nil
}

func (c *CartServiceImpl) AddProductToCart(ctx context.Context, cartID string, productID string) error {
	var product Product
	product, _ = c.product_service.GetProduct(ctx, productID)

	var cart Cart
	collection, _ := c.cart_db.GetCollection(ctx, "cart_database", "cart_collection")
	query := bson.D{{Key: "cartID", Value: cartID}}
	result, _ := collection.FindOne(ctx, query)
	result.One(ctx, &cart)

	if cart.ProductID == "" {
		cart.ProductID = product.ProductID
	}
	cart.Quantity += 1
	cart.PricePerUnit = product.PricePerUnit
	return collection.InsertOne(ctx, cart)
}

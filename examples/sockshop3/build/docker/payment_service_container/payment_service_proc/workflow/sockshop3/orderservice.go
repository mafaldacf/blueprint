// Package order implements the SockShop orders microservice.
//
// The service calls other services to collect information and then
// submits the order to the shipping service
package sockshop3

import (
	"context"
	"sync"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderService interface {
	// Place an order for the specified items
	NewOrder(ctx context.Context, customerID, addressID, cardID, cartID string) (Order, error)

	// Get all orders for a customer, sorted by date
	GetOrders(ctx context.Context, customerID string) ([]Order, error)

	// Get an order by ID
	GetOrder(ctx context.Context, orderID string) (Order, error)
}

// Creates a new [OrderService] instance.
// Customer, Address, and Card information will be looked up in the provided userService
// Successfully placed orders will be stored in [orderDB]
func NewOrderService(ctx context.Context, userService UserService, cartService CartService, payments PaymentService, shipping ShippingService, orderDB backend.NoSQLDatabase) (OrderService, error) {
	return &orderImpl{
		users:    userService,
		carts:    cartService,
		payments: payments,
		shipping: shipping,
		db:       orderDB,
	}, nil
}

type orderImpl struct {
	users    UserService
	carts    CartService
	payments PaymentService
	shipping ShippingService
	db       backend.NoSQLDatabase
}

// GetOrder implements OrderService.
func (s *orderImpl) GetOrder(ctx context.Context, orderID string) (Order, error) {
	filter := bson.D{{"id", orderID}}
	collection, _ := s.db.GetCollection(ctx, "order_service", "orders")
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Order{}, err
	}
	var order Order
	cursor.One(ctx, &order)
	return order, nil
}

// GetOrders implements OrderService.
func (s *orderImpl) GetOrders(ctx context.Context, customerID string) ([]Order, error) {
	filter := bson.D{{"customerid", customerID}}
	collection, _ := s.db.GetCollection(ctx, "order_service", "orders")
	cursor, err := collection.FindMany(ctx, filter)
	if err != nil {
		return nil, err
	}
	var orders []Order
	if err := cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// NewOrder implements OrderService.
func (s *orderImpl) NewOrder(ctx context.Context, customerID, addressID, cardID, cartID string) (Order, error) {
	// Fetch data concurrently
	var wg sync.WaitGroup
	wg.Add(4)

	var items []Item
	var addresses []Address
	var cards []Card

	go func() {
		defer wg.Done()
		items, _ = s.carts.GetCart(ctx, cartID)
	}()
	go func() {
		defer wg.Done()
		s.users.GetUsers(ctx, customerID)
	}()
	go func() {
		defer wg.Done()
		addresses, _ = s.users.GetAddresses(ctx, addressID)
	}()
	go func() {
		defer wg.Done()
		cards, _ = s.users.GetCards(ctx, cardID)
	}()

	// Await completion and validate responses
	wg.Wait()

	// Calculate total and authorize payment.
	amount := float32(10)
	s.payments.Authorise(ctx, amount)

	// Submit the shipment
	shipment := Shipment{
		ID:     uuid.NewString(),
		Name:   customerID,
		Status: "awaiting shipment",
	}
	shipment, err := s.shipping.PostShipping(ctx, shipment)
	if err != nil {
		return Order{}, err
	}

	// Save the order
	order := Order{
		ID:         shipment.ID,
		CustomerID: customerID,
		Address:    addresses[0],
		Card:       cards[0],
		Items:      items,
		Shipment:   shipment,
		Date:       time.Now().String(),
		Total:      amount,
	}
	collection, _ := s.db.GetCollection(ctx, "order_service", "orders")
	err = collection.InsertOne(ctx, order)
	if err != nil {
		return Order{}, err
	}

	// Delete the cart
	return order, s.carts.DeleteCart(ctx, customerID)
}

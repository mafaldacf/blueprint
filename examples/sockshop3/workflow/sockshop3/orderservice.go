// Package order implements the SockShop orders microservice.
//
// The service calls other services to collect information and then
// submits the order to the shipping service
package sockshop3

import (
	"context"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderService interface {
	NewOrder(ctx context.Context, customerID, addressID, cardID, cartID string) (Order, error)
	GetOrders(ctx context.Context, customerID string) ([]Order, error)
	GetOrder(ctx context.Context, orderID string) (Order, error)
}

func NewOrderServiceImpl(ctx context.Context, userService UserService, cartService CartService, payments PaymentService, shipping ShippingService, orderDB backend.NoSQLDatabase) (OrderService, error) {
	return &OrderServiceImpl{
		users:    userService,
		carts:    cartService,
		payments: payments,
		shipping: shipping,
		db:       orderDB,
	}, nil
}

type OrderServiceImpl struct {
	users    UserService
	carts    CartService
	payments PaymentService
	shipping ShippingService
	db       backend.NoSQLDatabase
}

func (s *OrderServiceImpl) GetOrder(ctx context.Context, orderID string) (Order, error) {
	collection, err := s.db.GetCollection(ctx, "order_db", "orders")
	if err != nil {
		return Order{}, err
	}

	filter := bson.D{{Key: "ID", Value: orderID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Order{}, err
	}
	var order Order
	ok, err := cursor.One(ctx, &order)
	if err != nil {
		return Order{}, err
	}
	if !ok {
		return Order{}, errors.Errorf("order %v does not exist", orderID)
	}
	return order, nil
}

func (s *OrderServiceImpl) GetOrders(ctx context.Context, customerID string) ([]Order, error) {
	collection, err := s.db.GetCollection(ctx, "order_db", "orders")
	if err != nil {
		return nil, err
	}
	
	filter := bson.D{{Key: "CustomerID", Value: customerID}}
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
func (s *OrderServiceImpl) NewOrder(ctx context.Context, customerID, addressID, cardID, cartID string) (Order, error) {
	// All arguments must be provided
	if customerID == "" {
		return Order{}, errors.Errorf("missing customerID")
	} else if addressID == "" {
		return Order{}, errors.Errorf("missing addressID")
	} else if cardID == "" {
		return Order{}, errors.Errorf("missing cardID")
	} else if cartID == "" {
		return Order{}, errors.Errorf("missing cartID")
	}
	
	// TODO
	/* // Fetch data concurrently
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
	wg.Wait() */

	items, err1 := s.carts.GetCart(ctx, cartID)
	if err1 != nil {
		return Order{}, err1
	}
	users, err2 := s.users.GetUsers(ctx, customerID)
	if err2 != nil {
		return Order{}, err2
	}
	addresses, err3 := s.users.GetAddresses(ctx, addressID)
	if err3 != nil {
		return Order{}, err3
	}
	cards, err4 := s.users.GetCards(ctx, cardID)
	if err4 != nil {
		return Order{}, err4
	}


	if len(items) == 0 {
		return Order{}, errors.Errorf("no items in cart")
	} else if len(users) == 0 {
		return Order{}, errors.Errorf("unknown customer %v", customerID)
	} else if len(addresses) == 0 {
		return Order{}, errors.Errorf("invalid address %v", addressID)
	} else if len(cards) == 0 {
		return Order{}, errors.Errorf("invalid card %v", cardID)
	}

	// Calculate total and authorize payment.
	// 1. calculate total
	amount := float32(0)
	shipping := float32(4.99)
	for _, item := range items {
		amount += float32(item.Quantity) * item.UnitPrice
	}
	amount += shipping
	// 2. authorise
	auth, err := s.payments.Authorise(ctx, amount)
	if err != nil {
		return Order{}, err
	} else if !auth.Authorised {
		return Order{}, errors.Errorf("payment not authorized due to %v", auth.Message)
	}

	// Submit the shipment
	shipment := Shipment{
		ID:     uuid.NewString(),
		Name:   customerID,
		Status: "awaiting shipment",
	}
	shipment, err = s.shipping.PostShipping(ctx, shipment)
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
	collection, err := s.db.GetCollection(ctx, "order_db", "orders")
	if err != nil {
		return Order{}, err
	}
	err = collection.InsertOne(ctx, order)
	if err != nil {
		return Order{}, err
	}

	// Delete the cart
	return order, nil
	//return order, s.carts.DeleteCart(ctx, customerID)
}

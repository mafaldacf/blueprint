package order

import (
	"context"
	"fmt"
	"sort"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderService interface {
	CancelOrder(ctx context.Context, command CancelOrderRequest) (CancelOrderResponse, error)
	CreateOrder(ctx context.Context, command CreateOrderCommand) (CreateOrderResult, error)
	UpdateOrder(ctx context.Context, command UpdateOrderCommand) (UpdateOrderResult, error)
	GetOrder(ctx context.Context, query GetOrderRequest) (GetOrderResponse, error)
	GetOrdersByUser(ctx context.Context, query GetOrdersByUserRequest) (GetOrdersByUserResponse, error)
}

type OrderServiceImpl struct {
	database backend.NoSQLDatabase
	queue    backend.Queue
}

func NewOrderServiceImpl(ctx context.Context, database backend.NoSQLDatabase, queue backend.Queue) (OrderService, error) {
	s := &OrderServiceImpl{
		database: database,
		queue:    queue,
	}
	return s, nil
}

func (s *OrderServiceImpl) CreateOrder(ctx context.Context, command CreateOrderCommand) (CreateOrderResult, error) {
	err := s.add(ctx, command.OrderDto)
	if err != nil {
		return CreateOrderResult{}, err
	}
	return CreateOrderResult{Id: command.OrderDto.Id}, nil
}

func (s *OrderServiceImpl) CancelOrder(ctx context.Context, command CancelOrderRequest) (CancelOrderResponse, error) {

}

func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, command UpdateOrderCommand) (UpdateOrderResult, error) {
	orderId := command.OrderDto.Id
	_, err := s.find(ctx, orderId)
	if err != nil {
		return UpdateOrderResult{false}, err
	}

	var newOrder OrderDto
	updateOrderWithNewValues(&newOrder, &command.OrderDto)
	s.update(ctx, newOrder)
	return UpdateOrderResult{true}, nil
}

func (s *OrderServiceImpl) GetOrder(ctx context.Context, query GetOrderRequest) (GetOrderResponse, error) {
	order, err := s.find(ctx, query.CustomerId)
	if err != nil {
		return GetOrderResponse{nil}, err
	}
	return GetOrderResponse{order}, nil
}

func (s *OrderServiceImpl) GetOrdersByUser(ctx context.Context, query GetOrdersByUserRequest) (GetOrdersByUserResponse, error) {
	customerId := query.CustomerId
	orders, err := s.findByUser(ctx, customerId)
	if err != nil {
		return GetOrdersByUserResponse{nil}, err
	}
	return GetOrdersByUserResponse{orders}, nil
}

func (s *OrderServiceImpl) add(ctx context.Context, order OrderDto) error {
	collection, err := s.database.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return err
	}
	return collection.InsertOne(ctx, order)
}

func (s *OrderServiceImpl) remove(ctx context.Context, id uuid.UUID) error {
	collection, err := s.database.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "Id", Value: id}}
	return collection.DeleteOne(ctx, filter)
}

func (s *OrderServiceImpl) find(ctx context.Context, id uuid.UUID) (OrderDto, error) {
	collection, err := s.database.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return OrderDto{}, err
	}
	filter := bson.D{{Key: "Id", Value: id}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return OrderDto{}, err
	}
	var order OrderDto
	ok, err := cursor.One(ctx, &order)
	if err != nil {
		return OrderDto{}, err
	}
	if !ok {
		return OrderDto{}, fmt.Errorf("order not found for id (%s)", id)
	}
	return order, nil
}

func (s *OrderServiceImpl) findByUser(ctx context.Context, customerId uuid.UUID) ([]OrderDto, error) {
	collection, err := s.database.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "CustomerId", Value: customerId}}
	cursor, err := collection.FindMany(ctx, filter)
	if err != nil {
		return nil, err
	}
	var orders []OrderDto
	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	// sort by OrderName
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].OrderName <= orders[j].OrderName
	})
	return orders, nil
}

func (s *OrderServiceImpl) update(ctx context.Context, order OrderDto) (OrderDto, error) {
	collection, err := s.database.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return OrderDto{}, err
	}
	filter := bson.D{{Key: "Id", Value: order.Id}}
	updated, err := collection.ReplaceOne(ctx, filter, order)
	if err != nil {
		return OrderDto{}, err
	}
	if updated == 0 {
		return OrderDto{}, fmt.Errorf("order not found for id (%s)", order.Id)
	}
	return order, nil
}

func updateOrderWithNewValues(order *OrderDto, orderUpdate *OrderDto) {
	var updatedShippingAddress AddressDto = orderUpdate.ShippingAddress
	var updatedBillingAddress AddressDto = orderUpdate.BillingAddress
	var updatedPayment PaymentDto = orderUpdate.Payment

	order.OrderName = orderUpdate.OrderName
	order.ShippingAddress = updatedShippingAddress
	order.BillingAddress = updatedBillingAddress
	order.Payment = updatedPayment
	order.Status = orderUpdate.Status
}

package eshopmicroservices

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
)

type OrderService interface {
	Run(ctx context.Context) error
	CreateNewOrder(ctx context.Context, command CreateOrderCommand) (CreateOrderResult, error)
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

func (s *OrderServiceImpl) Run(ctx context.Context) error {
	for {
		var message BasketChekoutEvent
		ok, err := s.queue.Pop(ctx, &message)
		if err != nil {
			return err
		}
		if !ok {
			continue
		}

		addressDto := AddressDto{message.FirstName, message.LastName, message.EmailAddress, message.AddressLine, message.Country, message.State, message.ZipCode}
		paymentDto := PaymentDto{message.CardName, message.CardNumber, message.Expiration, message.CVV, message.PaymentMethod}
		orderId := uuid.New()
		orderDto := OrderDto{
			Id:              orderId,
			CustomerId:      message.CustomerId,
			OrderName:       message.UserName,
			ShippingAddress: addressDto,
			BillingAddress:  addressDto,
			Payment:         paymentDto,
			Status:          Pending,
			OrderItems: []OrderItemDto{
				{orderId, uuid.MustParse("5334c996-8457-4cf0-815c-ed2b77c4ff61"), 2, 500},
				{orderId, uuid.MustParse("c67d6323-e8b1-4bdf-9a75-b0d0d2e7e914"), 1, 400},
			},
		}
		_, err = s.CreateNewOrder(ctx, CreateOrderCommand{OrderDto: orderDto})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *OrderServiceImpl) CreateNewOrder(ctx context.Context, command CreateOrderCommand) (CreateOrderResult, error) {
	err := s.add(ctx, command.OrderDto)
	if err != nil {
		return CreateOrderResult{}, err
	}
	return CreateOrderResult{Id: command.OrderDto.Id}, nil
}

func (s *OrderServiceImpl) add(ctx context.Context, order OrderDto) error {
	collection, err := s.database.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return err
	}
	return collection.InsertOne(ctx, order)
}

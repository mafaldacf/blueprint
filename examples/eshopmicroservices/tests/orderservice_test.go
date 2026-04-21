package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/order"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplequeue"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var orderServiceRegistry = registry.NewServiceRegistry[order.OrderService]("order_service")

func init() {
	orderServiceRegistry.Register("local", func(ctx context.Context) (order.OrderService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		queue, err := simplequeue.NewSimpleQueue(ctx)
		if err != nil {
			return nil, err
		}
		return order.NewOrderServiceImpl(ctx, db, queue)
	})
}

func makeAddress(firstName, lastName string) order.AddressDto {
	return order.AddressDto{
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: firstName + "@example.com",
		AddressLine:  "123 Test St",
		Country:      "US",
		State:        "NY",
		ZipCode:      "10001",
	}
}

func makePayment() order.PaymentDto {
	return order.PaymentDto{
		CardName:      "Test User",
		CardNumber:    "4242424242424242",
		Expiration:    "12/25",
		CCV:           "123",
		PaymentMethod: 1,
	}
}

func TestOrderServiceCreateNewOrder(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	customerId := uuid.NewString()
	orderDto := order.OrderDto{
		CustomerId:      customerId,
		OrderName:       "order@example.com",
		ShippingAddress: makeAddress("John", "Doe"),
		BillingAddress:  makeAddress("John", "Doe"),
		Payment:         makePayment(),
		Status:          order.Pending,
	}

	result, err := orderService.CreateNewOrder(ctx, order.CreateOrderCommand{OrderDto: orderDto})
	assert.NoError(t, err)
	assert.NotEmpty(t, result.Id)
}

func TestOrderServiceGetOrdersByCustomer(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	customerId := uuid.NewString()

	_, err = orderService.CreateNewOrder(ctx, order.CreateOrderCommand{OrderDto: order.OrderDto{
		CustomerId:      customerId,
		OrderName:       "cust1@example.com",
		ShippingAddress: makeAddress("Alice", "Smith"),
		BillingAddress:  makeAddress("Alice", "Smith"),
		Payment:         makePayment(),
		Status:          order.Pending,
	}})
	assert.NoError(t, err)

	_, err = orderService.CreateNewOrder(ctx, order.CreateOrderCommand{OrderDto: order.OrderDto{
		CustomerId:      customerId,
		OrderName:       "cust2@example.com",
		ShippingAddress: makeAddress("Alice", "Smith"),
		BillingAddress:  makeAddress("Alice", "Smith"),
		Payment:         makePayment(),
		Status:          order.Pending,
	}})
	assert.NoError(t, err)

	result, err := orderService.GetOrdersByCustomer(ctx, order.GetOrdersByCustomerQuery{CustomerId: customerId})
	assert.NoError(t, err)
	assert.True(t, len(result.Orders) >= 2)
}

func TestOrderServiceUpdateOrder(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	customerId := uuid.NewString()
	created, err := orderService.CreateNewOrder(ctx, order.CreateOrderCommand{OrderDto: order.OrderDto{
		CustomerId:      customerId,
		OrderName:       "update@example.com",
		ShippingAddress: makeAddress("Bob", "Jones"),
		BillingAddress:  makeAddress("Bob", "Jones"),
		Payment:         makePayment(),
		Status:          order.Draft,
	}})
	assert.NoError(t, err)
	orderId := created.Id

	result, err := orderService.UpdateOrder(ctx, order.UpdateOrderCommand{OrderDto: order.OrderDto{
		Id:              orderId,
		CustomerId:      customerId,
		OrderName:       "updated@example.com",
		ShippingAddress: makeAddress("Bob", "Jones"),
		BillingAddress:  makeAddress("Bob", "Jones"),
		Payment:         makePayment(),
		Status:          order.Completed,
	}})
	assert.NoError(t, err)
	assert.True(t, result.IsSuccess)
}

func TestOrderServiceDeleteOrder(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	created, err := orderService.CreateNewOrder(ctx, order.CreateOrderCommand{OrderDto: order.OrderDto{
		CustomerId:      uuid.NewString(),
		OrderName:       "delete@example.com",
		ShippingAddress: makeAddress("Carol", "White"),
		BillingAddress:  makeAddress("Carol", "White"),
		Payment:         makePayment(),
		Status:          order.Pending,
	}})
	assert.NoError(t, err)

	result, err := orderService.DeleteOrder(ctx, order.DeleteOrderCommand{Id: created.Id})
	assert.NoError(t, err)
	assert.True(t, result.IsSuccess)
}

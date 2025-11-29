package eshopmicroservices

import "github.com/google/uuid"

type CreateOrderCommand struct {
	OrderDto OrderDto
}

type CreateOrderResult struct {
	Id uuid.UUID
}

type UpdateOrderCommand struct {
	OrderDto OrderDto
}

type UpdateOrderResult struct {
	IsSuccess bool
}

type DeleteOrderCommand struct {
	Id uuid.UUID
}

type DeleteOrderResult struct {
	IsSuccess bool
}

type GetOrdersByCustomerQuery struct {
	CustomerId uuid.UUID
}

type GetOrdersByCustomerResult struct {
	Orders []OrderDto
}

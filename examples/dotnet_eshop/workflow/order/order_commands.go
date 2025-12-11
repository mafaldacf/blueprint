package order

import "github.com/google/uuid"

type CreateOrderCommand struct {
	OrderDto OrderDto
}

type CancelOrderRequest struct {
	OrderDto OrderDto
}

type CreateOrderResult struct {
	Id uuid.UUID
}

type CancelOrderResponse struct {
	Id uuid.UUID
}

type UpdateOrderCommand struct {
	OrderDto OrderDto
}

type UpdateOrderResult struct {
	IsSuccess bool
}

type GetOrdersByUserRequest struct {
	CustomerId uuid.UUID
}

type GetOrdersByUserResponse struct {
	Orders []OrderDto
}

type GetOrderRequest struct {
	CustomerId uuid.UUID
}

type GetOrderResponse struct {
	Orders OrderDto
}

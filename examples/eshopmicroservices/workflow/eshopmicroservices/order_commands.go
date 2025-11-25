package eshopmicroservices

import "github.com/google/uuid"

type CreateOrderCommand struct {
	OrderDto OrderDto
}

type CreateOrderResult struct {
	Id uuid.UUID
}

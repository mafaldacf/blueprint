package order

type CreateOrderCommand struct {
	OrderDto OrderDto
}

type CreateOrderResult struct {
	Id string
}

type UpdateOrderCommand struct {
	OrderDto OrderDto
}

type UpdateOrderResult struct {
	IsSuccess bool
}

type DeleteOrderCommand struct {
	Id string
}

type DeleteOrderResult struct {
	IsSuccess bool
}

type GetOrdersByCustomerQuery struct {
	CustomerId string
}

type GetOrdersByCustomerResult struct {
	Orders []OrderDto
}

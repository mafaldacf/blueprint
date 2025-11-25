package eshopmicroservices

import "github.com/google/uuid"

type OrderStatus int

const (
	Draft OrderStatus = iota + 1
	Pending
	Completed
	Cancelled
)

type OrderDto struct {
	Id              uuid.UUID
	CustomerId      uuid.UUID
	OrderName       string
	ShippingAddress AddressDto
	BillingAddress  AddressDto
	Payment         PaymentDto
	Status          OrderStatus
	OrderItems      []OrderItemDto
}

type OrderItemDto struct {
	OrderId   uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	Price     float64
}

type AddressDto struct {
	FirstName    string
	LastName     string
	EmailAddress string
	AddressLine  string
	Country      string
	State        string
	ZipCode      string
}

type PaymentDto struct {
	CardName      string
	CardNumber    string
	Expiration    string
	CCV           string
	PaymentMethod int
}

package eshopmicroservices

import "github.com/google/uuid"

type BasketCheckoutDto struct {
	UserName   string
	CustomerId uuid.UUID
	TotalPrice float64

	// Shipping and Billing Address
	FirstName    string
	LastName     string
	EmailAddress string
	AddressLine  string
	Country      string
	State        string
	ZipCode      string

	// Payment
	CardName      string
	CardNumber    string
	Expiration    string
	CVV           string
	PaymentMethod int
}

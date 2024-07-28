package order

import (
	"github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow/cart"
	"github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow/shipping"
	"github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow/user"
)

type Order struct {
	ID         string
	CustomerID string
	Customer   user.User
	Address    user.Address
	Card       user.Card
	Items      []cart.Item
	Shipment   shipping.Shipment
	Date       string
	Total      float32
}

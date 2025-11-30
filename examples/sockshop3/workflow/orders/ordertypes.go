package orders

import (
	"github.com/blueprint-uservices/blueprint/examples/sockshop3/workflow/carts"
	"github.com/blueprint-uservices/blueprint/examples/sockshop3/workflow/shipping"
	"github.com/blueprint-uservices/blueprint/examples/sockshop3/workflow/user"
)

type Order struct {
	ID         string
	CustomerID string
	Customer   user.User
	Address    user.Address
	Card       user.Card
	Items      []carts.Item
	Shipment   shipping.Shipment
	Date       string
	Total      float32
}

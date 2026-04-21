package orders

import (
	"github.com/blueprint-uservices/blueprint/examples/sockshop/workflow/carts"
	"github.com/blueprint-uservices/blueprint/examples/sockshop/workflow/shipping"
	"github.com/blueprint-uservices/blueprint/examples/sockshop/workflow/user"
)

type Order struct {
	ID         string          `bson:"ID"`
	CustomerID string          `bson:"CustomerID"`
	Customer   user.User       `bson:"Customer"`
	Address    user.Address    `bson:"Address"`
	Card       user.Card       `bson:"Card"`
	Items      []carts.Item    `bson:"Items"`
	Shipment   shipping.Shipment `bson:"Shipment"`
	Date       string          `bson:"Date"`
	Total      float32         `bson:"Total"`
}

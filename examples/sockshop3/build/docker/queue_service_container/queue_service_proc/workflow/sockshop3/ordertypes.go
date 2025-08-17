package sockshop3

type Order struct {
	ID         string
	CustomerID string
	Customer   User
	Address    Address
	Card       Card
	Items      []Item
	Shipment   Shipment
	Date       string
	Total      float32
}

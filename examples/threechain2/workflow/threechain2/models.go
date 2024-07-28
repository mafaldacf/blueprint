package threechain2

type Stock struct {
	ProductID string
	Quantity  int
}

type Bill struct {
	Username     string
	ProductID    string
	Quantity     int
	PricePerUnit int
	TotalCost    int
}

type Order struct {
	OrderID   string
	Username  string
	ProductID string
	Quantity  int
	Timestamp int64
}

type Cart struct {
	CartID       string
	Username     string
	ProductID    string
	Quantity     int
	PricePerUnit int
	Status       string
}

type Shipment struct {
	OrderID  string
	Username string
	Status   string
}

type ShipmentMessage struct {
	OrderID  string
	Username string
}

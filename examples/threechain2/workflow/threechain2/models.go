package threechain2

type Stock struct {
	Product  string
	Quantity int
}

type Order struct {
	OrderID   string
	Username  string
	Product   string
	Quantity  int
	Timestamp int64
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

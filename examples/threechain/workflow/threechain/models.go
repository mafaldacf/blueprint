package threechain

type Cart struct {
	Username string
	Product  string
}

type Stock struct {
	Product  string
	Quantity int
}

type Order struct {
	Username  string
	Product   string
	Timestamp int64
}

type ReserveStockMessage struct {
	Username string
	Product  string
}

type PlaceOrderMessage struct {
	Username string
	Product  string
}

package shopping_app

type Stock struct {
	ProductID string
	Quantity  int
}

type Product struct {
	ProductID    string
	Description  string
	PricePerUnit int
	Category     string
}

type Analytics struct {
	UserID     string
	Categories []string
}

type Bill struct {
	UserID       string
	ProductID    string
	Quantity     int
	PricePerUnit int
	TotalCost    int
}

type Order struct {
	OrderID   string
	UserID    string
	ProductID string
	Quantity  int
	Timestamp int64
}

type Cart struct {
	CartID       string
	UserID       string
	ProductID    string
	Quantity     int
	PricePerUnit int
	Status       string
}

type Shipment struct {
	OrderID string
	UserID  string
	Status  string
}

type ShipmentMessage struct {
	OrderID string
	UserID  string
}

type AnalyticsMessage struct {
	UserID          string
	ProductCategory string
}

package shopping_simple

type Product struct {
	ProductID    string
	Description  string
	PricePerUnit int
	Category     string
}

type Cart struct {
	CartID        string
	LastProductID string
	TotalQuantity int
	Products      []string
}

type CartProduct struct {
	CartID       string
	ProductID    string
	Quantity     int
	PricePerUnit int
}

type ProductQueueMessage struct {
	ProductID string
	Remove    bool
}

package basket

type CustomerBasket struct {
	BuyerID string
	Items   []BasketItem
}

type BasketItem struct {
	ID          string
	ProductID   int
	ProductName string
	UnitPrice float64
	OldUnitPrice float64
	Quantity int
	PictureUrl string
}

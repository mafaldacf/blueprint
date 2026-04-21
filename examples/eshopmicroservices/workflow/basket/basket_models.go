package basket

type ShoppingCart struct {
	UserName   string             `bson:"UserName"`
	Items      []ShoppingCartItem `bson:"Items"`
	TotalPrice float64            `bson:"TotalPrice"`
}

type ShoppingCartItem struct {
	Quantity    int     `bson:"Quantity"`
	Color       string  `bson:"Color"`
	Price       float64 `bson:"Price"`
	ProductId   string  `bson:"ProductId"`
	ProductName string  `bson:"ProductName"`
}

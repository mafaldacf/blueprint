package carts

type Cart struct {
	ID    string `bson:"ID"`
	Items []Item `bson:"Items"`
}

type Item struct {
	ID        string  `bson:"ID"`
	Quantity  int     `bson:"Quantity"`
	UnitPrice float32 `bson:"UnitPrice"`
}

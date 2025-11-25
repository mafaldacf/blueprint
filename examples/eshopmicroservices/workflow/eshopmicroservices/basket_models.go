package eshopmicroservices

import "github.com/google/uuid"

type ShoppingCart struct {
	UserName   string
	Items      []ShoppingCartItem
	TotalPrice float64
}

type ShoppingCartItem struct {
	Quantity    int
	Color       string
	Price       float64
	ProductId   uuid.UUID
	ProductName string
}

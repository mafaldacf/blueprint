package eshopmicroservices

import "github.com/google/uuid"

type CreateProductCommand struct {
	Name        string
	Category    []string
	Description string
	ImageFile   string
	Price       float64
}

type CreateProductResponse struct {
	Product Product
}

type DeleteProductCommand struct {
	Id uuid.UUID
}

type GetProductByIdQuery struct {
	Id uuid.UUID
}

type GetProductByIdResponse struct {
	Product Product
}

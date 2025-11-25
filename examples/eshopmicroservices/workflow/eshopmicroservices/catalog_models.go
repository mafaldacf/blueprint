package eshopmicroservices

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID
	Name        string
	Category    []string
	Description string
	ImageFile   string
	Price       float64
}

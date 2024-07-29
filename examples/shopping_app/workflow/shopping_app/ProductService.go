package shopping_app

import (
	"context"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductService interface {
	GetProduct(ctx context.Context, productID string) (Product, error)
	CreateProduct(ctx context.Context, productID string, description string, pricePerUnit int, category string) (Product, error)
}

type ProductServiceImpl struct {
	product_db backend.NoSQLDatabase
}

func NewProductServiceImpl(ctx context.Context, product_db backend.NoSQLDatabase) (ProductService, error) {
	return &ProductServiceImpl{product_db: product_db}, nil
}

func (s *ProductServiceImpl) GetProduct(ctx context.Context, productID string) (Product, error) {
	var product Product
	collection, _ := s.product_db.GetCollection(ctx, "product_database", "product_collection")
	query := bson.D{{Key: "productID", Value: productID}}
	result, _ := collection.FindOne(ctx, query)
	_, err := result.One(ctx, &product)
	return product, err
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, productID string, description string, pricePerUnit int, category string) (Product, error) {
	collection, _ := s.product_db.GetCollection(ctx, "product_database", "product_collection")
	product := Product{
		ProductID:    productID,
		Description:  description,
		PricePerUnit: pricePerUnit,
		Category:     category,
	}
	err := collection.InsertOne(ctx, product)
	return product, err
}

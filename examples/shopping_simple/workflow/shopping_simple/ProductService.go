package shopping_simple

import (
	"context"
	"fmt"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductService interface {
	GetProduct(ctx context.Context, productID string) (Product, error)
	GetAllProducts(ctx context.Context) ([]Product, error)
	CreateProduct(ctx context.Context, productID string, description string, pricePerUnit int, category string) (Product, error)
	DeleteProduct(ctx context.Context, productID string) (bool, error)
}

type ProductServiceImpl struct {
	product_db    backend.NoSQLDatabase
	product_queue backend.Queue
	num_workers   int
}

func NewProductServiceImpl(ctx context.Context, product_db backend.NoSQLDatabase, product_queue backend.Queue) (ProductService, error) {
	return &ProductServiceImpl{product_db: product_db, product_queue: product_queue}, nil
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

func (s *ProductServiceImpl) GetProduct(ctx context.Context, productID string) (Product, error) {
	var product Product
	collection, _ := s.product_db.GetCollection(ctx, "product_database", "product_collection")
	query := bson.D{{Key: "productid", Value: productID}}

	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return product, err
	}

	exists, err := result.One(ctx, &product)
	if err != nil {
		return product, fmt.Errorf("no product found for id '%s': %s", productID, err.Error())
	}
	if !exists {
		return product, fmt.Errorf("no product found for id '%s'", productID)
	}
	return product, err
}

func (s *ProductServiceImpl) GetAllProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	collection, _ := s.product_db.GetCollection(ctx, "product_database", "product_collection")
	filter := bson.D{}
	cursor, _ := collection.FindMany(ctx, filter)
	err := cursor.All(ctx, &products)
	return products, err
}

func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, productID string) (bool, error) {
	collection, _ := s.product_db.GetCollection(ctx, "product_database", "product_collection")
	filter := bson.D{{Key: "productid", Value: productID}}
	err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	var val bool
	message := ProductQueueMessage{
		ProductID: productID,
		Remove:    true,
	}
	val, err = s.product_queue.Push(ctx, message)
	return val, err
}

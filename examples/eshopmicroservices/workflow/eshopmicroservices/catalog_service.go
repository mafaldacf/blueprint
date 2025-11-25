package eshopmicroservices

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type CatalogService interface {
	Run(ctx context.Context) error
	CreateProduct(ctx context.Context, command CreateProductCommand) (CreateProductResponse, error)
	DeleteProduct(ctx context.Context, command DeleteProductCommand) error
	GetProductById(ctx context.Context, query GetProductByIdQuery) (GetProductByIdResponse, error)
}

type CatalogServiceImpl struct {
	database backend.NoSQLDatabase
}

func NewCatalogServiceImpl(ctx context.Context, database backend.NoSQLDatabase) (CatalogService, error) {
	s := &CatalogServiceImpl{
		database: database,
	}
	return s, nil
}

func (s *CatalogServiceImpl) Run(ctx context.Context) error {
	return nil
}

func (s *CatalogServiceImpl) CreateProduct(ctx context.Context, command CreateProductCommand) (CreateProductResponse, error) {
	product := Product{
		Name:        command.Name,
		Category:    command.Category,
		Description: command.Description,
		ImageFile:   command.ImageFile,
		Price:       command.Price,
	}
	err := s.store(ctx, product)
	if err != nil {
		return CreateProductResponse{}, err
	}
	return CreateProductResponse{Product: product}, nil
}

func (s *CatalogServiceImpl) DeleteProduct(ctx context.Context, command DeleteProductCommand) error {
	return s.delete(ctx, command.Id)
}

func (s *CatalogServiceImpl) GetProductById(ctx context.Context, query GetProductByIdQuery) (GetProductByIdResponse, error) {
	product, err := s.load(ctx, query.Id)
	if err != nil {
		return GetProductByIdResponse{}, err
	}
	return GetProductByIdResponse{Product: product}, nil
}

func (s *CatalogServiceImpl) store(ctx context.Context, product Product) error {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "product")
	if err != nil {
		return err
	}
	return collection.InsertOne(ctx, product)
}

func (s *CatalogServiceImpl) delete(ctx context.Context, id uuid.UUID) error {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "product")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "Id", Value: id}}
	return collection.DeleteOne(ctx, filter)
}

func (s *CatalogServiceImpl) load(ctx context.Context, id uuid.UUID) (Product, error) {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "product")
	if err != nil {
		return Product{}, err
	}
	filter := bson.D{{Key: "Id", Value: id}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Product{}, err
	}
	var product Product
	ok, err := cursor.One(ctx, &product)
	if err != nil {
		return Product{}, err
	}
	if !ok {
		return Product{}, fmt.Errorf("product not found for id (%s)", id)
	}
	return product, nil
}

package digota

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductService interface {
	New(ctx context.Context, name string, active bool, attributes []string, description string, images []string, metadata map[string]string, shippable bool, productUrl string) (*Product, error)
	Get(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context, page int64, limit int64, sort int32) (*ProductList, error)
	Update(ctx context.Context, id string, name string, active bool, attributes []string, description string, images []string, metadata map[string]string, shippable bool, productUrl string) (*Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewProductServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (ProductService, error) {
	return &ProductServiceImpl{db: db}, nil
}

func (s *ProductServiceImpl) New(ctx context.Context, name string, active bool, attributes []string, description string, images []string, metadata map[string]string, shippable bool, productUrl string) (*Product, error) {
	product := &Product{
		Id:          uuid.NewString(),
		Name:        name,
		Active:      active,
		Attributes:  attributes,
		Description: description,
		Images:      images,
		Metadata:    metadata,
		Shippable:   shippable,
		ProductUrl:  productUrl,
	}

	collection, err := s.db.GetCollection(ctx, "products_db", "products")
	if err != nil {
		return nil, err
	}
	err = collection.InsertOne(ctx, *product)
	return product, err
}

func (s *ProductServiceImpl) Get(ctx context.Context, id string) (*Product, error) {
	collection, err := s.db.GetCollection(ctx, "products_db", "products")
	if err != nil {
		return nil, fmt.Errorf("error getting collection : %v", err)
	}

	query := bson.D{{Key: "Id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying product: %v", err)
	}

	var product Product
	found, err := result.One(ctx, &product)
	if err != nil {
		return nil, fmt.Errorf("error finding product: %v", err)
	}
	if !found {
		return nil, fmt.Errorf("product not found for id (%s)", id)
	}

	return &product, nil
}

func (s *ProductServiceImpl) List(ctx context.Context, page int64, limit int64, sort int32) (*ProductList, error) {
	collection, err := s.db.GetCollection(ctx, "products_db", "products")
	if err != nil {
		return nil, fmt.Errorf("error getting collection: %s\n", err.Error())
	}

	result, err := collection.FindMany(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error finding products: %s\n", err.Error())
	}

	var products []Product
	err = result.All(ctx, &products)
	if err != nil {
		return nil, fmt.Errorf("error loading products: %s\n", err.Error())
	}

	productList := &ProductList{
		Products: products,
		Total:    int32(len(products)),
	}

	return productList, nil
}

func (s *ProductServiceImpl) Update(ctx context.Context, id string, name string, active bool, attributes []string, description string, images []string, metadata map[string]string, shippable bool, productUrl string) (*Product, error) {
	collection, err := s.db.GetCollection(ctx, "products_db", "products")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "Id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error finding product for id (%s): %v", id, err)
	}

	var product *Product = &Product{}
	found, err := result.One(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("error loading product for id (%s): %v", id, err)
	}
	if !found {
		return nil, fmt.Errorf("product not found for id (%s)", id)
	}

	// update fields and keep the rest the same
	product.Shippable = shippable
	product.Active = active

	if name != "" {
		product.Name = name
	}
	if description != "" {
		product.Description = description
	}
	if images != nil {
		product.Images = images
	}
	if attributes != nil {
		product.Attributes = attributes
	}
	if metadata != nil {
		product.Metadata = metadata
	}
	if productUrl != "" {
		product.ProductUrl = productUrl
	}

	filter := bson.D{{Key: "Id", Value: id}}
	n, err := collection.ReplaceOne(ctx, filter, product)
	if n == 0 {
		return nil, fmt.Errorf("product not found for id (%s)", id)
	}

	return product, err
}

func (s *ProductServiceImpl) Delete(ctx context.Context, id string) error {
	collection, err := s.db.GetCollection(ctx, "products_db", "products")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "Id", Value: id}}
	return collection.DeleteOne(ctx, filter)
}

package digota

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/blueprint-uservices/blueprint/examples/digota/workflow/digota/validation"
)

type ProductService interface {
	New(ctx context.Context, name string, active bool, attributes []string, description string, images []string, metadata map[string]string, shippable bool, url string) (*Product, error)
	Get(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context, page int64, limit int64, sort int32) (*ProductList, error)
	Update(ctx context.Context, id string, name string, active bool, attributes []string, description string, images []string, metadata map[string]string, shippable bool, url string) (*Product, error)
	Delete(ctx context.Context, id string) error
}

type ProductServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewProductServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (ProductService, error) {
	s := &ProductServiceImpl{db: db}
	return s, nil
}

func (s *ProductServiceImpl) New(ctx context.Context, name string, active bool, attributes []string, description string, images []string, metadata map[string]string, shippable bool, url string) (*Product, error) {
	product := &Product{
		Name:        name,
		Active:      active,
		Attributes:  attributes,
		Description: description,
		Images:      images,
		Metadata:    metadata,
		Shippable:   shippable,
		Url:         url,
	}

	err := validation.Validate(product)
	if err != nil {
		return nil, err
	}

	collection, err := s.db.GetCollection(ctx, "products", "products")
	if err != nil {
		return nil, err
	}
	err = collection.InsertOne(ctx, *product)
	return product, err
}

func (s *ProductServiceImpl) Get(ctx context.Context, id string) (*Product, error) {
	collection, err := s.db.GetCollection(ctx, "products", "products")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	var product *Product
	found, err := result.One(ctx, product)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("product not found for id (%s)", id)
	}

	return product, nil
}

func (s *ProductServiceImpl) List(ctx context.Context, page int64, limit int64, sort int32) (*ProductList, error) {
	collection, err := s.db.GetCollection(ctx, "products", "products")
	if err != nil {
		return nil, err
	}

	result, err := collection.FindMany(ctx, nil)
	if err != nil {
		return nil, err
	}

	var products []*Product
	err = result.All(ctx, products)
	if err != nil {
		return nil, err
	}

	productList := &ProductList{
		Products: products,
		Total:    int32(len(products)),
	}

	return productList, nil
}

func (s *ProductServiceImpl) Update(ctx context.Context, id string, name string, active bool, attributes []string, description string, images []string, metadata map[string]string, shippable bool, url string) (*Product, error) {
	collection, err := s.db.GetCollection(ctx, "products", "products")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	var product *Product = &Product{}
	found, err := result.One(ctx, product)
	if err != nil {
		return nil, err
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
	if url != "" {
		product.Url = url
	}

	err = collection.InsertOne(ctx, *product)
	return product, err
}

func (s *ProductServiceImpl) Delete(ctx context.Context, id string) error {
	collection, err := s.db.GetCollection(ctx, "products", "products")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "id", Value: id}}
	err = collection.DeleteOne(ctx, filter)
	return err
}

package digota

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type SkuService interface {
	New(ctx context.Context, name string, currency int32, active bool, price uint64, parent string, metadata map[string]string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error)
	Get(ctx context.Context, id string) (*Sku, error)
	List(ctx context.Context, page int64, limit int64, sort int32) (*SkuList, error)
	Update(ctx context.Context, id string, name string, currency int32, active bool, price uint64, parent string, metadata map[string]string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error)
	Delete(ctx context.Context, id string) error
}

type SkuServiceImpl struct {
	db    backend.NoSQLDatabase
	/* queue backend.Queue */
	productService ProductService
}

func NewSkuServiceImpl(ctx context.Context, productService ProductService, db backend.NoSQLDatabase/* , queue backend.Queue */) (SkuService, error) {
	s := &SkuServiceImpl{productService: productService,  db: db/* , queue: queue */}
	return s, nil
}

func (s *SkuServiceImpl) New(ctx context.Context, name string, currency int32, active bool, price uint64, parent string, metadata map[string]string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error) {
	product, err := s.productService.Get(ctx, parent)
	if err != nil {
		return nil, err
	}

	var validAttr = make(map[string]string)
	for k, v := range attributes {
		for _, pv := range product.Attributes {
			if k == pv {
				validAttr[k] = v
			}
		}
	}
	
	sku := &Sku{
		Name:              name,
		Currency:          currency,
		Active:            active,
		Price:             price,
		Parent:            parent,
		Metadata:          metadata,
		PackageDimensions: packageDimensions,
		Inventory:         inventory,
		Attributes:        validAttr,
	}

	collection, err := s.db.GetCollection(ctx, "skus_db", "skus")
	if err != nil {
		return nil, err
	}
	err = collection.InsertOne(ctx, sku)
	return sku, err
}

func (s *SkuServiceImpl) Get(ctx context.Context, id string) (*Sku, error) {
	collection, err := s.db.GetCollection(ctx, "skus_db", "skus")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "Id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	var sku *Sku
	found, err := result.One(ctx, sku)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("sku not found for id (%s)", id)
	}

	return sku, nil
}

func (s *SkuServiceImpl) List(ctx context.Context, page int64, limit int64, sort int32) (*SkuList, error) {
	collection, err := s.db.GetCollection(ctx, "skus_db", "skus")
	if err != nil {
		return nil, err
	}

	result, err := collection.FindMany(ctx, nil)
	if err != nil {
		return nil, err
	}

	var skus []*Sku
	err = result.All(ctx, skus)
	if err != nil {
		return nil, err
	}

	skuList := &SkuList{
		Orders: skus,
		Total:  int32(len(skus)),
	}

	return skuList, nil
}

func (s *SkuServiceImpl) Update(ctx context.Context, id string, name string, currency int32, active bool, price uint64, parent string, metadata map[string]string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error) {
	collection, err := s.db.GetCollection(ctx, "skus_db", "skus")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "Id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	var sku *Sku = &Sku{}
	found, err := result.One(ctx, sku)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("sku not found for id (%s)", id)
	}

	product, err := s.productService.Get(ctx, parent)
	if err != nil {
		return nil, err
	}

	var attrs = make(map[string]string)

	if x := attributes; x != nil {
		attrs = x
	} else {
		attrs = sku.Attributes
	}

	var validAttr = make(map[string]string)
	// save only the valid attr
	for k, v := range attrs {
		for _, pv := range product.Attributes {
			if k == pv {
				validAttr[k] = v
			}
		}
	}
	sku.Attributes = validAttr

	// update fields
	
	sku.Active = active

	if name != "" {
		sku.Name = name
	}
	if price != 0 {
		sku.Price = price
	}
	if !CurrencyIsReserved(currency) {
		sku.Currency = currency
	}
	if metadata != nil {
		sku.Metadata = metadata
	}
	if image != "" {
		sku.Image = image
	}
	if packageDimensions != nil {
		sku.PackageDimensions = packageDimensions
	}
	if inventory != nil {
		sku.Inventory = inventory
	}

	filter := bson.D{{Key: "Id", Value: id}}
	ok, err := collection.Upsert(ctx, filter, sku)
	if !ok {
		return nil, fmt.Errorf("sku not updated for id (%s)", id)
	}
	return sku, err
}


func (s *SkuServiceImpl) Delete(ctx context.Context, id string) error {
	collection, err := s.db.GetCollection(ctx, "skus_db", "skus")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "Id", Value: id}}
	err = collection.DeleteOne(ctx, filter)

	/* if err != nil {
		return err
	}

	message := QueueMessage{
		id: id,
	}
	_, err = s.queue.Push(ctx, message) */

	return err
}

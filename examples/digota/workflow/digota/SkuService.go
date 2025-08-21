package digota

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
	//"github.com/blueprint-uservices/blueprint/examples/digota/workflow/digota/validation"
)

type SkuService interface {
	New(ctx context.Context, name string, currency int32, active bool, price uint64, parent string, metadata map[string]string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error)
	//New2(ctx context.Context, id string, currency int32, price uint64, parent string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error)
	Get(ctx context.Context, id string) (*Sku, error)
	/* List(ctx context.Context, page int64, limit int64, sort int32) (*SkuList, error)
	Update(ctx context.Context, id string, name string, currency int32, active bool, price uint64, parent string, metadata map[string]string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error) */
	Delete(ctx context.Context, id string) error
}

type SkuServiceImpl struct {
	db    backend.NoSQLDatabase
	queue backend.Queue
	/* productService ProductService */
}

func NewSkuServiceImpl(ctx context.Context /* , productService ProductService */, db backend.NoSQLDatabase, queue backend.Queue) (SkuService, error) {
	s := &SkuServiceImpl{ /* productService: productService,  */ db: db, queue: queue}
	return s, nil
}

/* func (s *SkuServiceImpl) New2(ctx context.Context, id string, currency int32, price uint64, parent string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error) {
	product, err := s.productService.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	sku := &Sku{
			Name:              product.Name,
			Currency:          currency,
			Active:            product.Active,
			Price:             price,
			Parent:            parent,
			Metadata:          product.Metadata,
			PackageDimensions: packageDimensions,
			Inventory:         inventory,
			Attributes:        attributes,
		}

	err = validation.Validate(sku)
	if err != nil {
		return nil, err
	}

	collection, err := s.db.GetCollection(ctx, "skus", "skus")
	if err != nil {
		return nil, err
	}
	err = collection.InsertOne(ctx, *sku)
	return sku, err

} */

func (s *SkuServiceImpl) New(ctx context.Context, name string, currency int32, active bool, price uint64, parent string, metadata map[string]string, image string, packageDimensions *PackageDimensions, inventory *Inventory, attributes map[string]string) (*Sku, error) {
	sku := &Sku{
		Name:              name,
		Currency:          currency,
		Active:            active,
		Price:             price,
		Parent:            parent,
		Metadata:          metadata,
		PackageDimensions: packageDimensions,
		Inventory:         inventory,
		Attributes:        attributes,
	}

	/* err := validation.Validate(sku)
	if err != nil {
		return nil, err
	} */

	collection, err := s.db.GetCollection(ctx, "skus_db", "skus")
	if err != nil {
		return nil, err
	}
	err = collection.InsertOne(ctx, *sku)
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

/* func (s *SkuServiceImpl) List(ctx context.Context, page int64, limit int64, sort int32) (*SkuList, error) {
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

	var item *Sku = &Sku{}
	found, err := result.One(ctx, item)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("sku not found for id (%s)", id)
	}

	// update fields
	item.Attributes = attributes
	item.Active = active

	if name != "" {
		item.Name = name
	}
	if price != 0 {
		item.Price = price
	}
	if !CurrencyIsReserved(currency) {
		item.Currency = currency
	}
	if metadata != nil {
		item.Metadata = metadata
	}
	if image != "" {
		item.Image = image
	}
	if packageDimensions != nil {
		item.PackageDimensions = packageDimensions
	}
	if inventory != nil {
		item.Inventory = inventory
	}

	err = collection.InsertOne(ctx, *item)
	return item, err
}
*/

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

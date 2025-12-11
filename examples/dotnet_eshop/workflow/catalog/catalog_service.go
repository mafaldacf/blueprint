package catalog

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/blueprint/pkg/coreplugins/backend"
	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type CatalogService interface {
	CreateItem(ctx context.Context, command CreateItemRequest) (CreateItemResponse, error)
	UpdateItem(ctx context.Context, productToUpdate CatalogItem) (UpdateItemResponse, error)
	DeleteItem(ctx context.Context, command DeleteItemRequest) error
	GetItemsByIDs(ctx context.Context, query GetItemsByIDsRequest) (GetItemsByIDsResponse, error)
	GetItemByID(ctx context.Context, query GetItemByIDRequest) (GetItemByIDResponse, error)
	GetItemByName(ctx context.Context, query GetItemByNameRequest) (GetItemByNameResponse, error)
	GetAllItems(ctx context.Context) (GetAllItemsResponse, error)
}

type CatalogServiceImpl struct {
	database backend.NoSQLDatabase
	queue    backend.Queue
}

func NewCatalogServiceImpl(ctx context.Context, database backend.NoSQLDatabase, queue backend.Queue) (CatalogService, error) {
	s := &CatalogServiceImpl{
		database: database,
		queue:    queue,
	}
	return s, nil
}

func (s *CatalogServiceImpl) CreateItem(ctx context.Context, request CreateItemRequest) (CreateItemResponse, error) {
	item := CatalogItem{
		ID:                request.ID,
		Name:              request.Name,
		Description:       request.Description,
		Price:             request.Price,
		PictureFileName:   request.PriceFileName,
		CatalogTypeID:     request.CatalogTypeID,
		CatalogType:       request.CatalogType,
		CatalogBrandID:    request.CatalogBrandID,
		CatalogBrand:      request.CatalogBrand,
		AvailableStock:    request.AvailableStock,
		RestockThreshold:  request.RestockThreshold,
		MaxStockThreshold: request.MaxStockThreshold,
	}
	err := s.save(ctx, item)
	if err != nil {
		return CreateItemResponse{}, err
	}
	return CreateItemResponse{item}, nil
}

func (s *CatalogServiceImpl) UpdateItem(ctx context.Context, productToUpdate CatalogItem) (UpdateItemResponse, error) {
	catalogItem, err := s.get(ctx, productToUpdate.ID)
	if err != nil {
		return UpdateItemResponse{}, err
	}

	oldPrice := catalogItem.Price

	catalogItem.Name = productToUpdate.Name
	catalogItem.Description = productToUpdate.Description
	catalogItem.Price = productToUpdate.Price
	catalogItem.PictureFileName = productToUpdate.PictureFileName
	catalogItem.AvailableStock = productToUpdate.AvailableStock
	catalogItem.CatalogBrandID = productToUpdate.CatalogBrandID
	catalogItem.CatalogTypeID = productToUpdate.CatalogTypeID

	if oldPrice != catalogItem.Price {
		// Create integration event (equiv. ProductPriceChangedIntegrationEvent)
		priceChangedEvent := &ProductPriceChangedEvent{
			CatalogItemID: catalogItem.ID,
			NewPrice:      catalogItem.Price,
			OldPrice:      oldPrice,
		}

		err := s.update(ctx, catalogItem)
		if err != nil {
			return UpdateItemResponse{}, fmt.Errorf("update catalog item: %w", err)
		}

		s.queue.Push(ctx, priceChangedEvent)
	} else {
		err := s.update(ctx, catalogItem)
		if err != nil {
			return UpdateItemResponse{}, fmt.Errorf("update catalog item: %w", err)
		}
	}

	return UpdateItemResponse{ID: catalogItem.ID}, nil

}

func (s *CatalogServiceImpl) DeleteItem(ctx context.Context, request DeleteItemRequest) error {
	return s.remove(ctx, request.ID)
}

func (s *CatalogServiceImpl) GetItemByID(ctx context.Context, request GetItemByIDRequest) (GetItemByIDResponse, error) {
	item, err := s.get(ctx, request.ID)
	if err != nil {
		return GetItemByIDResponse{}, err
	}
	return GetItemByIDResponse{item}, nil
}

func (s *CatalogServiceImpl) GetItemsByID(ctx context.Context, request GetItemsByIDsRequest) (GetItemsByIDsResponse, error) {
	items, err := s.getMany(ctx, request.IDs)
	if err != nil {
		return GetItemsByIDsResponse{}, err
	}
	return GetItemsByIDsResponse{items}, nil
}

func (s *CatalogServiceImpl) GetItemByName(ctx context.Context, query GetItemByNameRequest) (GetItemByNameResponse, error) {
	item, err := s.getByName(ctx, query.Name)
	if err != nil {
		return GetItemByNameResponse{}, err
	}
	return GetItemByNameResponse{item}, nil
}

func (s *CatalogServiceImpl) GetAllItems(ctx context.Context) (GetAllItemsResponse, error) {
	products, err := s.getAll(ctx)
	if err != nil {
		return GetAllItemsResponse{}, err
	}
	return GetAllItemsResponse{products}, nil
}

func (s *CatalogServiceImpl) save(ctx context.Context, item CatalogItem) error {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "item")
	if err != nil {
		return err
	}
	return collection.InsertOne(ctx, item)
}

func (s *CatalogServiceImpl) update(ctx context.Context, item CatalogItem) error {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "item")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "Id", Value: item.ID}}
	return collection.UpdateOne(ctx, filter, item)
}

func (s *CatalogServiceImpl) remove(ctx context.Context, id string) error {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "item")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "Id", Value: id}}
	return collection.DeleteOne(ctx, filter)
}

func (s *CatalogServiceImpl) get(ctx context.Context, id int) (CatalogItem, error) {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "item")
	if err != nil {
		return CatalogItem{}, err
	}
	filter := bson.D{{Key: "ID", Value: id}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return CatalogItem{}, err
	}
	var item CatalogItem
	ok, err := cursor.One(ctx, &item)
	if err != nil {
		return CatalogItem{}, err
	}
	if !ok {
		return CatalogItem{}, fmt.Errorf("item not found for id (%s)", id)
	}
	return item, nil
}

func (s *CatalogServiceImpl) getByName(ctx context.Context, category string) (CatalogItem, error) {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "item")
	if err != nil {
		return CatalogItem{}, err
	}
	filter := bson.D{{Key: "Name", Value: category}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return CatalogItem{}, err
	}
	var item CatalogItem
	ok, err := cursor.One(ctx, &item)
	if err != nil {
		return CatalogItem{}, err
	}
	if !ok {
		return CatalogItem{}, fmt.Errorf("item not found for category (%s)", category)
	}
	return item, nil
}

func (s *CatalogServiceImpl) getMany(ctx context.Context, ids []int) ([]CatalogItem, error) {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "item")
	if err != nil {
		return nil, err
	}
	filter := bson.D{
		{Key: "ID", Value: bson.D{
			{Key: "$in", Value: ids},
		}},
	}
	cursor, err := collection.FindMany(ctx, filter)
	if err != nil {
		return nil, err
	}
	var products []CatalogItem
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *CatalogServiceImpl) getAll(ctx context.Context) ([]CatalogItem, error) {
	collection, err := s.database.GetCollection(ctx, "catalog_db", "item")
	if err != nil {
		return nil, err
	}
	cursor, err := collection.FindMany(ctx, nil)
	if err != nil {
		return nil, err
	}
	var products []CatalogItem
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

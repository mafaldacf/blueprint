package train_ticket2

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type ConsignService interface {
	// TODO:
	// - UpdateConsign
	// - FindByAccountId
	// - FindByConsignee
	InsertConsign(ctx context.Context, consignRequest ConsignRequest) (ConsignRecord, error)
	FindByOrderId(ctx context.Context, orderID string) (ConsignRecord, error)
}

type ConsignServiceImpl struct {
	consignPriceService ConsignPriceService
	consignDB           backend.NoSQLDatabase
}

func NewConsignServiceImpl(ctx context.Context, consignPriceService ConsignPriceService, consignDB backend.NoSQLDatabase) (ConsignService, error) {
	return &ConsignServiceImpl{consignPriceService: consignPriceService, consignDB: consignDB}, nil
}

func (c *ConsignServiceImpl) InsertConsign(ctx context.Context, consignRequest ConsignRequest) (ConsignRecord, error) {
	price, err := c.consignPriceService.GetPriceByWeightAndRegion(ctx, consignRequest.Weight, consignRequest.IsWithin)
	if err != nil {
		return ConsignRecord{}, err
	}

	consignRecord := ConsignRecord{
		OrderID:    consignRequest.OrderID,
		AccountID:  consignRequest.AccountID,
		HandleDate: consignRequest.HandleDate,
		TargetDate: consignRequest.TargetDate,
		FromPlace:  consignRequest.From,
		ToPlace:    consignRequest.To,
		Consignee:  consignRequest.Consignee,
		Phone:      consignRequest.Phone,
		Weight:     consignRequest.Weight,
		Price:      price,
	}

	collection, err := c.consignDB.GetCollection(ctx, "consign_db", "consign_record")
	if err != nil {
		return ConsignRecord{}, err
	}
	err = collection.InsertOne(ctx, consignRecord)
	if err != nil {
		return ConsignRecord{}, err
	}

	return consignRecord, nil
}

func (c *ConsignServiceImpl) FindByOrderId(ctx context.Context, orderID string) (ConsignRecord, error) {
	collection, err := c.consignDB.GetCollection(ctx, "consign_db", "consign_record")
	if err != nil {
		return ConsignRecord{}, err
	}

	filter := bson.D{{Key: "OrderID", Value: orderID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return ConsignRecord{}, err
	}

	var consignRecord ConsignRecord
	ok, err := cursor.One(ctx, &consignRecord)
	if err != nil {
		return ConsignRecord{}, err
	}
	if !ok {
		return ConsignRecord{}, fmt.Errorf("consign (%s) not found", orderID)
	}
	return consignRecord, nil
}

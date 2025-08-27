package train_ticket2

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type ConsignService interface {
	InsertConsign(ctx context.Context, consignRequest ConsignRequest) (ConsignRecord, error)
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

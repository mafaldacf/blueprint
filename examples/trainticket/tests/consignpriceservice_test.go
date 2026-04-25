package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var consignPriceServiceRegistry = registry.NewServiceRegistry[trainticket.ConsignPriceService]("consignprice_service")

func init() {
	consignPriceServiceRegistry.Register("local", func(ctx context.Context) (trainticket.ConsignPriceService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewConsignPriceServiceImpl(ctx, db)
	})
}

func TestConsignPriceServiceCreateAndGetConfig(t *testing.T) {
	ctx := context.Background()
	service, err := consignPriceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	priceConfig := trainticket.ConsignPrice{
		ID:            "consignprice001",
		Index:         0,
		InitialWeight: 2.0,
		InitialPrice:  2.0,
		WithinPrice:   1.0,
		BeyondPrice:   1.5,
	}
	created, err := service.CreateAndModifyPriceConfig(ctx, priceConfig)
	assert.NoError(t, err)
	assert.Equal(t, 2.0, created.InitialWeight)
	assert.Equal(t, 2.0, created.InitialPrice)
}

func TestConsignPriceServiceGetPriceByWeightAndRegion(t *testing.T) {
	ctx := context.Background()
	service, err := consignPriceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	priceConfig := trainticket.ConsignPrice{
		ID:            "consignprice002",
		Index:         0,
		InitialWeight: 2.0,
		InitialPrice:  2.0,
		WithinPrice:   1.0,
		BeyondPrice:   1.5,
	}
	_, err = service.CreateAndModifyPriceConfig(ctx, priceConfig)
	assert.NoError(t, err)

	// weight within initial limit
	price, err := service.GetPriceByWeightAndRegion(ctx, 1.5, true)
	assert.NoError(t, err)
	assert.Equal(t, 2.0, price)

	// weight beyond initial limit, within region
	price, err = service.GetPriceByWeightAndRegion(ctx, 5.0, true)
	assert.NoError(t, err)
	assert.Greater(t, price, 2.0)
}

func TestConsignPriceServiceGetPriceConfig(t *testing.T) {
	ctx := context.Background()
	service, err := consignPriceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	priceConfig := trainticket.ConsignPrice{
		ID:            "consignprice003",
		Index:         0,
		InitialWeight: 3.0,
		InitialPrice:  3.0,
		WithinPrice:   1.5,
		BeyondPrice:   2.0,
	}
	_, err = service.CreateAndModifyPriceConfig(ctx, priceConfig)
	assert.NoError(t, err)

	config, err := service.GetPriceConfig(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, config.ID)
}

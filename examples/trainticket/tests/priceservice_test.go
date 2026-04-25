package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var priceServiceRegistry = registry.NewServiceRegistry[trainticket.PriceService]("price_service")

func init() {
	priceServiceRegistry.Register("local", func(ctx context.Context) (trainticket.PriceService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewPriceServiceImpl(ctx, db)
	})
}

func TestPriceServiceCreateAndFind(t *testing.T) {
	ctx := context.Background()
	service, err := priceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	config := trainticket.PriceConfig{
		ID:                  "price001",
		TrainType:           "GaoTie",
		RouteID:             "route001",
		BasicPriceRate:      0.38,
		FirstClassPriceRate: 0.50,
	}
	err = service.CreateNewPriceConfig(ctx, config)
	assert.NoError(t, err)

	found, err := service.FindByRouteIDAndTrainType(ctx, "route001", "GaoTie")
	assert.NoError(t, err)
	assert.Equal(t, "price001", found.ID)
	assert.Equal(t, 0.38, found.BasicPriceRate)
	assert.Equal(t, 0.50, found.FirstClassPriceRate)
}

func TestPriceServiceGetAllPriceConfig(t *testing.T) {
	ctx := context.Background()
	service, err := priceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	config := trainticket.PriceConfig{
		ID:                  "price002",
		TrainType:           "DongChe",
		RouteID:             "route002",
		BasicPriceRate:      0.30,
		FirstClassPriceRate: 0.40,
	}
	err = service.CreateNewPriceConfig(ctx, config)
	assert.NoError(t, err)

	all, err := service.GetAllPriceConfig(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestPriceServiceCreateMultipleConfigs(t *testing.T) {
	ctx := context.Background()
	service, err := priceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	config1 := trainticket.PriceConfig{
		ID:             "price003",
		TrainType:      "GaoTie",
		RouteID:        "route003",
		BasicPriceRate: 0.45,
	}
	config2 := trainticket.PriceConfig{
		ID:             "price004",
		TrainType:      "GaoTie",
		RouteID:        "route004",
		BasicPriceRate: 0.55,
	}
	err = service.CreateNewPriceConfig(ctx, config1)
	assert.NoError(t, err)
	err = service.CreateNewPriceConfig(ctx, config2)
	assert.NoError(t, err)

	found1, err := service.FindByRouteIDAndTrainType(ctx, "route003", "GaoTie")
	assert.NoError(t, err)
	assert.Equal(t, "price003", found1.ID)

	found2, err := service.FindByRouteIDAndTrainType(ctx, "route004", "GaoTie")
	assert.NoError(t, err)
	assert.Equal(t, "price004", found2.ID)
}

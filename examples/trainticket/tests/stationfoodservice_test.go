package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var stationFoodServiceRegistry = registry.NewServiceRegistry[trainticket.StationFoodService]("stationfood_service")

func init() {
	stationFoodServiceRegistry.Register("local", func(ctx context.Context) (trainticket.StationFoodService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewStationFoodServiceImpl(ctx, db)
	})
}

func TestStationFoodServiceCreateAndList(t *testing.T) {
	ctx := context.Background()
	service, err := stationFoodServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	store := trainticket.StationFoodStore{
		ID:           "store001",
		StationName:  "Shanghai",
		StoreName:    "Shanghai Dumplings",
		Telephone:    "021-12345678",
		BusinessTime: "08:00-20:00",
		DeliveryFee:  5.0,
	}
	err = service.CreateFoodStore(ctx, store)
	assert.NoError(t, err)

	all, err := service.ListFoodStores(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestStationFoodServiceListByStationName(t *testing.T) {
	ctx := context.Background()
	service, err := stationFoodServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	store1 := trainticket.StationFoodStore{
		ID:          "store002",
		StationName: "Beijing",
		StoreName:   "Beijing Noodles",
		DeliveryFee: 6.0,
	}
	store2 := trainticket.StationFoodStore{
		ID:          "store003",
		StationName: "Beijing",
		StoreName:   "Beijing BBQ",
		DeliveryFee: 8.0,
	}
	err = service.CreateFoodStore(ctx, store1)
	assert.NoError(t, err)
	err = service.CreateFoodStore(ctx, store2)
	assert.NoError(t, err)

	stores, err := service.ListFoodStoresByStationName(ctx, "Beijing")
	assert.NoError(t, err)
	assert.Len(t, stores, 2)
}

func TestStationFoodServiceGetByID(t *testing.T) {
	ctx := context.Background()
	service, err := stationFoodServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	store := trainticket.StationFoodStore{
		ID:          "store004",
		StationName: "Nanjing",
		StoreName:   "Nanjing Salted Duck",
		DeliveryFee: 4.0,
	}
	err = service.CreateFoodStore(ctx, store)
	assert.NoError(t, err)

	found, err := service.GetFoodStoreByID(ctx, "store004")
	assert.NoError(t, err)
	assert.Equal(t, "Nanjing Salted Duck", found.StoreName)
	assert.Equal(t, "Nanjing", found.StationName)
}

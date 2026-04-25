package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var trainFoodServiceRegistry = registry.NewServiceRegistry[trainticket.TrainFoodService]("trainfood_service")

func init() {
	trainFoodServiceRegistry.Register("local", func(ctx context.Context) (trainticket.TrainFoodService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewTrainFoodServiceImpl(ctx, db)
	})
}

func TestTrainFoodServiceCreate(t *testing.T) {
	ctx := context.Background()
	service, err := trainFoodServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	tf := trainticket.TrainFood{
		ID:     "trainfood001",
		TripID: "G1234",
		Foods: []trainticket.Food{
			{Name: "Rice Box", Price: 25.0},
			{Name: "Noodles", Price: 20.0},
		},
	}
	created, err := service.CreateTrainFood(ctx, tf)
	assert.NoError(t, err)
	assert.Equal(t, "trainfood001", created.ID)
	assert.Equal(t, "G1234", created.TripID)
	assert.Len(t, created.Foods, 2)
}

func TestTrainFoodServiceListAll(t *testing.T) {
	ctx := context.Background()
	service, err := trainFoodServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	tf := trainticket.TrainFood{
		ID:     "trainfood002",
		TripID: "D5678",
		Foods:  []trainticket.Food{{Name: "Sandwich", Price: 15.0}},
	}
	_, err = service.CreateTrainFood(ctx, tf)
	assert.NoError(t, err)

	all, err := service.ListTrainFood(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestTrainFoodServiceListByTripID(t *testing.T) {
	ctx := context.Background()
	service, err := trainFoodServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	tf := trainticket.TrainFood{
		ID:     "trainfood003",
		TripID: "G9999",
		Foods: []trainticket.Food{
			{Name: "Bento Box", Price: 30.0},
			{Name: "Fruit Salad", Price: 18.0},
		},
	}
	_, err = service.CreateTrainFood(ctx, tf)
	assert.NoError(t, err)

	foods, err := service.ListTrainFoodByTripID(ctx, "G9999")
	assert.NoError(t, err)
	assert.Len(t, foods, 2)
	assert.Equal(t, "Bento Box", foods[0].Name)
}

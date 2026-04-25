package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var trainServiceRegistry = registry.NewServiceRegistry[trainticket.TrainService]("train_service")

func init() {
	trainServiceRegistry.Register("local", func(ctx context.Context) (trainticket.TrainService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewTrainServiceImpl(ctx, db)
	})
}

func TestTrainServiceCreateAndRetrieve(t *testing.T) {
	ctx := context.Background()
	service, err := trainServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	trainType := trainticket.TrainType{
		ID:           "train001",
		Name:         "GaoTie",
		EconomyClass: 2000,
		ComfortClass: 2500,
		AvgSpeed:     350,
	}
	err = service.Create(ctx, trainType)
	assert.NoError(t, err)

	found, err := service.Retrieve(ctx, "train001")
	assert.NoError(t, err)
	assert.Equal(t, "GaoTie", found.Name)
	assert.Equal(t, int64(350), found.AvgSpeed)
}

func TestTrainServiceRetrieveByName(t *testing.T) {
	ctx := context.Background()
	service, err := trainServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	trainType := trainticket.TrainType{
		ID:           "train002",
		Name:         "DongChe",
		EconomyClass: 1500,
		ComfortClass: 2000,
		AvgSpeed:     250,
	}
	err = service.Create(ctx, trainType)
	assert.NoError(t, err)

	found, err := service.RetrieveByName(ctx, "DongChe")
	assert.NoError(t, err)
	assert.Equal(t, "train002", found.ID)
	assert.Equal(t, 1500, found.EconomyClass)
}

func TestTrainServiceAllTrains(t *testing.T) {
	ctx := context.Background()
	service, err := trainServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	trainType := trainticket.TrainType{
		ID:       "train003",
		Name:     "PuTong",
		AvgSpeed: 120,
	}
	err = service.Create(ctx, trainType)
	assert.NoError(t, err)

	all, err := service.AllTrains(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestTrainServiceUpdate(t *testing.T) {
	ctx := context.Background()
	service, err := trainServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	trainType := trainticket.TrainType{
		ID:       "train004",
		Name:     "TeShu",
		AvgSpeed: 200,
	}
	err = service.Create(ctx, trainType)
	assert.NoError(t, err)

	trainType.AvgSpeed = 220
	updated, err := service.Update(ctx, trainType)
	assert.NoError(t, err)
	assert.True(t, updated)
}

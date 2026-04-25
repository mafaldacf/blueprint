package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var stationServiceRegistry = registry.NewServiceRegistry[trainticket.StationService]("station_service")

func init() {
	stationServiceRegistry.Register("local", func(ctx context.Context) (trainticket.StationService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewStationServiceImpl(ctx, db)
	})
}

func TestStationServiceCreateAndFind(t *testing.T) {
	ctx := context.Background()
	service, err := stationServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	station := trainticket.Station{
		ID:       "station001",
		Name:     "Shanghai",
		StayTime: 5,
	}
	err = service.CreateStation(ctx, station)
	assert.NoError(t, err)

	found, err := service.FindByID(ctx, "station001")
	assert.NoError(t, err)
	assert.Equal(t, "Shanghai", found.Name)
	assert.Equal(t, int64(5), found.StayTime)
}

func TestStationServiceExists(t *testing.T) {
	ctx := context.Background()
	service, err := stationServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	station := trainticket.Station{
		ID:   "station002",
		Name: "Beijing",
	}
	err = service.CreateStation(ctx, station)
	assert.NoError(t, err)

	exists, err := service.Exists(ctx, "Beijing")
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestStationServiceFindID(t *testing.T) {
	ctx := context.Background()
	service, err := stationServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	station := trainticket.Station{
		ID:   "station003",
		Name: "Nanjing",
	}
	err = service.CreateStation(ctx, station)
	assert.NoError(t, err)

	id, err := service.FindID(ctx, "Nanjing")
	assert.NoError(t, err)
	assert.Equal(t, "station003", id)
}

func TestStationServiceFindAll(t *testing.T) {
	ctx := context.Background()
	service, err := stationServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	station1 := trainticket.Station{ID: "station004", Name: "Suzhou"}
	station2 := trainticket.Station{ID: "station005", Name: "Hangzhou"}
	err = service.CreateStation(ctx, station1)
	assert.NoError(t, err)
	err = service.CreateStation(ctx, station2)
	assert.NoError(t, err)

	all, err := service.FindAll(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

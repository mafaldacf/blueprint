package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var configServiceRegistry = registry.NewServiceRegistry[trainticket.ConfigService]("config_service")

func init() {
	configServiceRegistry.Register("local", func(ctx context.Context) (trainticket.ConfigService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewConfigServiceImpl(ctx, db)
	})
}

func TestConfigServiceCreateAndFind(t *testing.T) {
	ctx := context.Background()
	service, err := configServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	conf := trainticket.Config{
		Name:        "max_capacity",
		Value:       "100",
		Description: "Maximum seat capacity per train car",
	}
	err = service.Create(ctx, conf)
	assert.NoError(t, err)

	found, err := service.Find(ctx, "max_capacity")
	assert.NoError(t, err)
	assert.Equal(t, "max_capacity", found.Name)
	assert.Equal(t, "100", found.Value)
	assert.Equal(t, "Maximum seat capacity per train car", found.Description)
}

func TestConfigServiceUpdate(t *testing.T) {
	ctx := context.Background()
	service, err := configServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	conf := trainticket.Config{
		Name:  "booking_window",
		Value: "30",
	}
	err = service.Create(ctx, conf)
	assert.NoError(t, err)

	conf.Value = "60"
	updated, err := service.Update(ctx, conf)
	assert.NoError(t, err)
	assert.True(t, updated)
}

func TestConfigServiceFindAll(t *testing.T) {
	ctx := context.Background()
	service, err := configServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	conf1 := trainticket.Config{Name: "seat_economy", Value: "2000"}
	conf2 := trainticket.Config{Name: "seat_comfort", Value: "2500"}
	err = service.Create(ctx, conf1)
	assert.NoError(t, err)
	err = service.Create(ctx, conf2)
	assert.NoError(t, err)

	all, err := service.FindAll(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

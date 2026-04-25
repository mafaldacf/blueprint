package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_mediamicroservices/workflow/mediamicroservices"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var plotServiceRegistry = registry.NewServiceRegistry[mediamicroservices.PlotService]("plot_service")

func init() {
	plotServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.PlotService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewPlotServiceImpl(ctx, db, cache)
	})
}

func TestPlotServiceWrite(t *testing.T) {
	ctx := context.Background()
	service, err := plotServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	plot, err := service.WritePlot(ctx, 0, 1, "A hero's journey begins in a small village where an unlikely champion rises to face an ancient evil.")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), plot.PlotID)
	assert.Equal(t, "A hero's journey begins in a small village where an unlikely champion rises to face an ancient evil.", plot.Plot)
}

func TestPlotServiceWriteMultiple(t *testing.T) {
	ctx := context.Background()
	service, err := plotServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	plot1, err := service.WritePlot(ctx, 0, 2, "An epic tale of adventure and discovery across uncharted lands.")
	assert.NoError(t, err)
	assert.Equal(t, int64(2), plot1.PlotID)

	plot2, err := service.WritePlot(ctx, 0, 3, "A gripping thriller that keeps you on the edge of your seat until the very end.")
	assert.NoError(t, err)
	assert.Equal(t, int64(3), plot2.PlotID)
}

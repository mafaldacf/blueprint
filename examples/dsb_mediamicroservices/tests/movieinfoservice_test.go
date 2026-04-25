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

var movieInfoServiceRegistry = registry.NewServiceRegistry[mediamicroservices.MovieInfoService]("movieinfo_service")

func init() {
	movieInfoServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.MovieInfoService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		socialGraphDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewMovieInfoServiceImpl(ctx, db, cache, socialGraphDB)
	})
}

func TestMovieInfoServiceWrite(t *testing.T) {
	ctx := context.Background()
	service, err := movieInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	casts := []mediamicroservices.Cast{
		{CastID: "cast1", CastInfoID: "castinfo1", Character: "Hero"},
	}

	movieInfo, err := service.WriteMovieInfo(ctx, 0, "movie001", "The Adventure", casts, 1, []string{}, []string{}, []string{}, "4.5", 100)
	assert.NoError(t, err)
	assert.Equal(t, "movie001", movieInfo.MovieID)
	assert.Equal(t, "The Adventure", movieInfo.Title)
	assert.Len(t, movieInfo.Casts, 1)
	assert.Equal(t, int64(1), movieInfo.PlotID)
}

func TestMovieInfoServiceWriteWithMultipleCasts(t *testing.T) {
	ctx := context.Background()
	service, err := movieInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	casts := []mediamicroservices.Cast{
		{CastID: "cast2", CastInfoID: "castinfo2", Character: "Protagonist"},
		{CastID: "cast3", CastInfoID: "castinfo3", Character: "Antagonist"},
		{CastID: "cast4", CastInfoID: "castinfo4", Character: "Sidekick"},
	}

	movieInfo, err := service.WriteMovieInfo(ctx, 0, "movie002", "The Grand Heist", casts, 2, []string{"thumb1.jpg"}, []string{"photo1.jpg"}, []string{"video1.mp4"}, "3.8", 42)
	assert.NoError(t, err)
	assert.Equal(t, "movie002", movieInfo.MovieID)
	assert.Equal(t, "The Grand Heist", movieInfo.Title)
	assert.Len(t, movieInfo.Casts, 3)
}

func TestMovieInfoServiceWriteNoCasts(t *testing.T) {
	ctx := context.Background()
	service, err := movieInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	movieInfo, err := service.WriteMovieInfo(ctx, 0, "movie003", "The Silent Film", []mediamicroservices.Cast{}, 3, []string{}, []string{}, []string{}, "0.0", 0)
	assert.NoError(t, err)
	assert.Equal(t, "movie003", movieInfo.MovieID)
	assert.Equal(t, "The Silent Film", movieInfo.Title)
	assert.Empty(t, movieInfo.Casts)
}

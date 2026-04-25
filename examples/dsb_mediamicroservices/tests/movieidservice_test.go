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

var movieIdServiceRegistry = registry.NewServiceRegistry[mediamicroservices.MovieIdService]("movieid_service")

func init() {
	movieIdServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.MovieIdService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		composeReviewService, err := composeReviewServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		ratingService, err := ratingServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewMovieIdServiceImpl(ctx, db, cache, composeReviewService, ratingService)
	})
}

func TestMovieIdServiceRegister(t *testing.T) {
	ctx := context.Background()
	service, err := movieIdServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	movieId, err := service.RegisterMovieId(ctx, 0, "tt0111161", "The Shawshank Redemption")
	assert.NoError(t, err)
	assert.Equal(t, "tt0111161", movieId.MovieID)
	assert.Equal(t, "The Shawshank Redemption", movieId.Title)
}

func TestMovieIdServiceRegisterMultiple(t *testing.T) {
	ctx := context.Background()
	service, err := movieIdServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	movie1, err := service.RegisterMovieId(ctx, 0, "tt0068646", "The Godfather")
	assert.NoError(t, err)
	assert.Equal(t, "tt0068646", movie1.MovieID)
	assert.Equal(t, "The Godfather", movie1.Title)

	movie2, err := service.RegisterMovieId(ctx, 0, "tt0071562", "The Godfather Part II")
	assert.NoError(t, err)
	assert.Equal(t, "tt0071562", movie2.MovieID)
	assert.Equal(t, "The Godfather Part II", movie2.Title)
}

func TestMovieIdServiceRead(t *testing.T) {
	ctx := context.Background()
	service, err := movieIdServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.RegisterMovieId(ctx, 0, "tt0468569", "The Dark Knight")
	assert.NoError(t, err)

	movieId, err := service.ReadMovieId(ctx, 0, "The Dark Knight")
	assert.NoError(t, err)
	assert.Equal(t, "tt0468569", movieId.MovieID)
	assert.Equal(t, "The Dark Knight", movieId.Title)
}

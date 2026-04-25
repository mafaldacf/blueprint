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

var movieReviewServiceRegistry = registry.NewServiceRegistry[mediamicroservices.MovieReviewService]("movie_review_service")

func init() {
	movieReviewServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.MovieReviewService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		reviewStorageService, err := reviewStorageServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewMovieReviewServiceImpl(ctx, db, cache, reviewStorageService)
	})
}

func TestMovieReviewServiceUpload(t *testing.T) {
	ctx := context.Background()
	service, err := movieReviewServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.UploadMovieReview(ctx, 0, "movie001", 10, 1700000000)
	assert.NoError(t, err)
}

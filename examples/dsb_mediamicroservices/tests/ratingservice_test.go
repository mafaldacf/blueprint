package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_mediamicroservices/workflow/mediamicroservices"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/stretchr/testify/assert"
)

var ratingServiceRegistry = registry.NewServiceRegistry[mediamicroservices.RatingService]("rating_service")

func init() {
	ratingServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.RatingService, error) {
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		composeReviewService, err := composeReviewServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewRatingServiceImpl(ctx, cache, composeReviewService)
	})
}

func TestRatingServiceUpload(t *testing.T) {
	ctx := context.Background()
	service, err := ratingServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Use a unique reqID to avoid triggering compose-and-upload (counter < NUM_COMPONENTS)
	err = service.UploadNewRating(ctx, 2000, "movie001", 4)
	assert.NoError(t, err)
}

func TestRatingServiceUploadDifferentMovies(t *testing.T) {
	ctx := context.Background()
	service, err := ratingServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.UploadNewRating(ctx, 2001, "movie002", 5)
	assert.NoError(t, err)

	err = service.UploadNewRating(ctx, 2002, "movie003", 3)
	assert.NoError(t, err)
}

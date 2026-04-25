package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_mediamicroservices/workflow/mediamicroservices"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/stretchr/testify/assert"
)

var composeReviewServiceRegistry = registry.NewServiceRegistry[mediamicroservices.ComposeReviewService]("compose_review_service")

func init() {
	composeReviewServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.ComposeReviewService, error) {
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		reviewStorageService, err := reviewStorageServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		userReviewService, err := userReviewServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		movieReviewService, err := movieReviewServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewComposeReviewServiceImpl(ctx, cache, reviewStorageService, userReviewService, movieReviewService)
	})
}

func TestComposeReviewServiceUploadText(t *testing.T) {
	ctx := context.Background()
	service, err := composeReviewServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Use a unique reqID to avoid triggering compose-and-upload (counter < NUM_COMPONENTS)
	err = service.UploadText(ctx, 1000, "This is a compelling and thought-provoking film.")
	assert.NoError(t, err)
}

func TestComposeReviewServiceUploadRating(t *testing.T) {
	ctx := context.Background()
	service, err := composeReviewServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.UploadRating(ctx, 1001, 4)
	assert.NoError(t, err)
}

func TestComposeReviewServiceUploadMovieId(t *testing.T) {
	ctx := context.Background()
	service, err := composeReviewServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.UploadMovieId(ctx, 1002, "tt0111161")
	assert.NoError(t, err)
}

func TestComposeReviewServiceUploadUserId(t *testing.T) {
	ctx := context.Background()
	service, err := composeReviewServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.UploadUserId(ctx, 1003, 300)
	assert.NoError(t, err)
}

func TestComposeReviewServiceUploadUniqueId(t *testing.T) {
	ctx := context.Background()
	service, err := composeReviewServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.UploadUniqueId(ctx, 1004, 99999)
	assert.NoError(t, err)
}

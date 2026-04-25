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

var userReviewServiceRegistry = registry.NewServiceRegistry[mediamicroservices.UserReviewService]("user_review_service")

func init() {
	userReviewServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.UserReviewService, error) {
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
		return mediamicroservices.NewUserReviewServiceImpl(ctx, db, cache, reviewStorageService)
	})
}

func TestUserReviewServiceUpload(t *testing.T) {
	ctx := context.Background()
	service, err := userReviewServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.UploadUserReview(ctx, 0, 200, 10, 1700000000)
	assert.NoError(t, err)
}

func TestUserReviewServiceUploadForNewUser(t *testing.T) {
	ctx := context.Background()
	service, err := userReviewServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.UploadUserReview(ctx, 1, 201, 11, 1700000001)
	assert.NoError(t, err)
}

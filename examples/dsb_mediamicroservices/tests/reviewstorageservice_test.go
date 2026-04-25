package tests

import (
	"context"
	"testing"
	"time"

	"github.com/blueprint-uservices/blueprint/examples/dsb_mediamicroservices/workflow/mediamicroservices"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var reviewStorageServiceRegistry = registry.NewServiceRegistry[mediamicroservices.ReviewStorageService]("review_storage_service")

func init() {
	reviewStorageServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.ReviewStorageService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewReviewStorageServiceImpl(ctx, db, cache)
	})
}

func TestReviewStorageServiceStore(t *testing.T) {
	ctx := context.Background()
	service, err := reviewStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	review := mediamicroservices.Review{
		ReviewID:  1,
		ReqID:     0,
		Timestamp: time.Now().UnixNano(),
		UserID:    100,
		MovieID:   "movie001",
		Text:      "A captivating film with stunning visuals.",
		Rating:    5,
	}

	stored, err := service.StoreReview(ctx, 0, review)
	assert.NoError(t, err)
	assert.Equal(t, review.ReviewID, stored.ReviewID)
	assert.Equal(t, review.Text, stored.Text)
	assert.Equal(t, review.Rating, stored.Rating)
	assert.Equal(t, review.MovieID, stored.MovieID)
	assert.Equal(t, review.UserID, stored.UserID)
}

func TestReviewStorageServiceStoreMultiple(t *testing.T) {
	ctx := context.Background()
	service, err := reviewStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	review1 := mediamicroservices.Review{
		ReviewID: 2,
		UserID:   101,
		MovieID:  "movie002",
		Text:     "A masterpiece of modern cinema.",
		Rating:   5,
	}
	review2 := mediamicroservices.Review{
		ReviewID: 3,
		UserID:   102,
		MovieID:  "movie002",
		Text:     "Entertaining but predictable.",
		Rating:   3,
	}

	stored1, err := service.StoreReview(ctx, 1, review1)
	assert.NoError(t, err)
	assert.Equal(t, review1.ReviewID, stored1.ReviewID)

	stored2, err := service.StoreReview(ctx, 2, review2)
	assert.NoError(t, err)
	assert.Equal(t, review2.ReviewID, stored2.ReviewID)
}

func TestReviewStorageServiceReadEmpty(t *testing.T) {
	ctx := context.Background()
	service, err := reviewStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	reviews, err := service.ReadReviews(ctx, 0, []int64{})
	assert.NoError(t, err)
	assert.Nil(t, reviews)
}

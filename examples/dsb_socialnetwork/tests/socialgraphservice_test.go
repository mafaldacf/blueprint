package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_socialnetwork/workflow/socialnetwork"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var socialGraphServiceRegistry = registry.NewServiceRegistry[socialnetwork.SocialGraphService]("socialgraph_service")

func init() {
	socialGraphServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.SocialGraphService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		userIDService, err := userIDServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewSocialGraphServiceImpl(ctx, cache, db, userIDService)
	})
}

func TestSocialGraphServiceInsertUser(t *testing.T) {
	ctx := context.Background()
	service, err := socialGraphServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.InsertUser(ctx, 0, 10)
	assert.NoError(t, err)
}

func TestSocialGraphServiceFollowAndGetFollowers(t *testing.T) {
	ctx := context.Background()
	service, err := socialGraphServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// user 20 follows user 21 — Follow populates the cache used by GetFollowers
	err = service.InsertUser(ctx, 0, 20)
	assert.NoError(t, err)
	err = service.InsertUser(ctx, 0, 21)
	assert.NoError(t, err)

	err = service.Follow(ctx, 0, 20, 21)
	assert.NoError(t, err)

	// GetFollowers for user 21 resolves from cache (populated by Follow)
	followers, err := service.GetFollowers(ctx, 0, 21)
	assert.NoError(t, err)
	assert.Contains(t, followers, int64(20))
}

func TestSocialGraphServiceFollowAndGetFollowees(t *testing.T) {
	ctx := context.Background()
	service, err := socialGraphServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.InsertUser(ctx, 0, 30)
	assert.NoError(t, err)
	err = service.InsertUser(ctx, 0, 31)
	assert.NoError(t, err)

	err = service.Follow(ctx, 0, 30, 31)
	assert.NoError(t, err)

	// GetFollowees for user 30 resolves from cache (populated by Follow)
	followees, err := service.GetFollowees(ctx, 0, 30)
	assert.NoError(t, err)
	assert.Contains(t, followees, int64(31))
}

func TestSocialGraphServiceFollowAndUnfollow(t *testing.T) {
	ctx := context.Background()
	service, err := socialGraphServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.InsertUser(ctx, 0, 40)
	assert.NoError(t, err)
	err = service.InsertUser(ctx, 0, 41)
	assert.NoError(t, err)

	err = service.Follow(ctx, 0, 40, 41)
	assert.NoError(t, err)

	err = service.Unfollow(ctx, 0, 40, 41)
	assert.NoError(t, err)

	// After unfollow, cache is updated — followees of 40 should be empty
	followees, err := service.GetFollowees(ctx, 0, 40)
	assert.NoError(t, err)
	assert.NotContains(t, followees, int64(41))
}

func TestSocialGraphServiceMultipleFollowers(t *testing.T) {
	ctx := context.Background()
	service, err := socialGraphServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.InsertUser(ctx, 0, 50)
	assert.NoError(t, err)
	err = service.InsertUser(ctx, 0, 51)
	assert.NoError(t, err)
	err = service.InsertUser(ctx, 0, 52)
	assert.NoError(t, err)

	err = service.Follow(ctx, 0, 51, 50)
	assert.NoError(t, err)
	err = service.Follow(ctx, 0, 52, 50)
	assert.NoError(t, err)

	followers, err := service.GetFollowers(ctx, 0, 50)
	assert.NoError(t, err)
	assert.Contains(t, followers, int64(51))
	assert.Contains(t, followers, int64(52))
}

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/blueprint-uservices/blueprint/examples/dsb_socialnetwork/workflow/socialnetwork"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var userTimelineServiceRegistry = registry.NewServiceRegistry[socialnetwork.UserTimelineService]("usertimeline_service")

func init() {
	userTimelineServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.UserTimelineService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		postStorageService, err := postStorageServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewUserTimelineServiceImpl(ctx, cache, db, postStorageService)
	})
}

func TestUserTimelineServiceReadEmptyRange(t *testing.T) {
	ctx := context.Background()
	service, err := userTimelineServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	ids, err := service.ReadUserTimeline(ctx, 0, 9001, 0, 0)
	assert.NoError(t, err)
	assert.Empty(t, ids)
}

func TestUserTimelineServiceReadInvalidRange(t *testing.T) {
	ctx := context.Background()
	service, err := userTimelineServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	ids, err := service.ReadUserTimeline(ctx, 0, 9001, 5, 3)
	assert.NoError(t, err)
	assert.Empty(t, ids)
}

func TestUserTimelineServiceWriteUserTimeline(t *testing.T) {
	ctx := context.Background()
	service, err := userTimelineServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.WriteUserTimeline(ctx, 0, 8001, 9002, time.Now().UnixNano())
	assert.NoError(t, err)
}

func TestUserTimelineServiceWriteAndRead(t *testing.T) {
	ctx := context.Background()

	// Store the post so ReadPosts can find it
	postService, err := postStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	post := makePost(8100, 9003, "alice", "Timeline test post.")
	err = postService.StorePost(ctx, 0, post)
	assert.NoError(t, err)

	service, err := userTimelineServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.WriteUserTimeline(ctx, 0, 8100, 9003, time.Now().UnixNano())
	assert.NoError(t, err)

	// ReadUserTimeline resolves from cache (populated by WriteUserTimeline), avoids $slice DB query
	ids, err := service.ReadUserTimeline(ctx, 0, 9003, 0, 1)
	assert.NoError(t, err)
	assert.Contains(t, ids, int64(8100))
}

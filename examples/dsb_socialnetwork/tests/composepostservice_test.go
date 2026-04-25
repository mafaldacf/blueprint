package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_socialnetwork/workflow/socialnetwork"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/stretchr/testify/assert"
)

var composePostServiceRegistry = registry.NewServiceRegistry[socialnetwork.ComposePostService]("composepost_service")

func init() {
	composePostServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.ComposePostService, error) {
		postStorageService, err := postStorageServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		userTimelineService, err := userTimelineServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		userService, err := userServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		uniqueIdService, err := uniqueIdServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		mediaService, err := mediaServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		textService, err := textServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		homeTimelineService, err := homeTimelineServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewComposePostServiceImpl(ctx, postStorageService, userTimelineService, userService, uniqueIdService, mediaService, textService, homeTimelineService)
	})
}

func TestComposePostServiceComposePost(t *testing.T) {
	ctx := context.Background()

	// Register the author and set up social graph so WriteHomeTimeline can resolve followers
	userService, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	err = userService.RegisterUserWithId(ctx, 0, "Hank", "Adams", "hankadams", "hankpass", 1000)
	assert.NoError(t, err)

	socialGraph, err := socialGraphServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	err = socialGraph.InsertUser(ctx, 0, 1001)
	assert.NoError(t, err)
	// user 1001 follows user 1000 — populates "1000:followers" cache
	err = socialGraph.Follow(ctx, 0, 1001, 1000)
	assert.NoError(t, err)

	service, err := composePostServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	postID, mentionIDs, err := service.ComposePost(ctx, 0, "hankadams", 1000, "Hello, world!", []int64{}, []string{}, socialnetwork.POST)
	assert.NoError(t, err)
	assert.Greater(t, postID, int64(0))
	assert.Empty(t, mentionIDs)
}

func TestComposePostServiceComposePostWithMedia(t *testing.T) {
	ctx := context.Background()

	userService, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	err = userService.RegisterUserWithId(ctx, 0, "Iris", "Baker", "irisbaker", "irispass", 1100)
	assert.NoError(t, err)

	socialGraph, err := socialGraphServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	err = socialGraph.InsertUser(ctx, 0, 1101)
	assert.NoError(t, err)
	err = socialGraph.Follow(ctx, 0, 1101, 1100)
	assert.NoError(t, err)

	service, err := composePostServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	postID, _, err := service.ComposePost(ctx, 0, "irisbaker", 1100, "Check out this photo!", []int64{9001}, []string{"photo"}, socialnetwork.POST)
	assert.NoError(t, err)
	assert.Greater(t, postID, int64(0))
}

func TestComposePostServiceComposePostWithUrl(t *testing.T) {
	ctx := context.Background()

	userService, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	err = userService.RegisterUserWithId(ctx, 0, "Jack", "Cole", "jackcole", "jackpass", 1200)
	assert.NoError(t, err)

	socialGraph, err := socialGraphServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	err = socialGraph.InsertUser(ctx, 0, 1201)
	assert.NoError(t, err)
	err = socialGraph.Follow(ctx, 0, 1201, 1200)
	assert.NoError(t, err)

	service, err := composePostServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	postID, _, err := service.ComposePost(ctx, 0, "jackcole", 1200, "Interesting: http://example.com/article", []int64{}, []string{}, socialnetwork.POST)
	assert.NoError(t, err)
	assert.Greater(t, postID, int64(0))
}

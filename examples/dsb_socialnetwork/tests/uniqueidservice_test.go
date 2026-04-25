package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_socialnetwork/workflow/socialnetwork"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/stretchr/testify/assert"
)

var uniqueIdServiceRegistry = registry.NewServiceRegistry[socialnetwork.UniqueIdService]("uniqueid_service")

func init() {
	uniqueIdServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.UniqueIdService, error) {
		return socialnetwork.NewUniqueIdServiceImpl(ctx)
	})
}

func TestUniqueIdServiceComposeUniqueId(t *testing.T) {
	ctx := context.Background()
	service, err := uniqueIdServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	id, err := service.ComposeUniqueId(ctx, 0, socialnetwork.POST)
	assert.NoError(t, err)
	assert.Greater(t, id, int64(0))
}

func TestUniqueIdServiceComposeUniqueIdDifferentPostTypes(t *testing.T) {
	ctx := context.Background()
	service, err := uniqueIdServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	id1, err := service.ComposeUniqueId(ctx, 0, socialnetwork.POST)
	assert.NoError(t, err)
	id2, err := service.ComposeUniqueId(ctx, 0, socialnetwork.REPOST)
	assert.NoError(t, err)
	assert.NotEqual(t, id1, id2)
}

func TestUniqueIdServiceComposeMultipleUniqueIds(t *testing.T) {
	ctx := context.Background()
	service, err := uniqueIdServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	seen := make(map[int64]bool)
	for i := 0; i < 5; i++ {
		id, err := service.ComposeUniqueId(ctx, int64(i), socialnetwork.POST)
		assert.NoError(t, err)
		assert.False(t, seen[id], "duplicate unique id generated")
		seen[id] = true
	}
}

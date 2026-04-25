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

var userMentionServiceRegistry = registry.NewServiceRegistry[socialnetwork.UserMentionService]("usermention_service")

func init() {
	userMentionServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.UserMentionService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewUserMentionServiceImpl(ctx, cache, db)
	})
}

func TestUserMentionServiceComposeEmpty(t *testing.T) {
	ctx := context.Background()
	service, err := userMentionServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	mentions, err := service.ComposeUserMentions(ctx, 0, []string{})
	assert.NoError(t, err)
	assert.Nil(t, mentions)
}

func TestUserMentionServiceComposeUserMentions(t *testing.T) {
	ctx := context.Background()
	service, err := userMentionServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	mentions, err := service.ComposeUserMentions(ctx, 0, []string{"alice", "bob"})
	assert.NoError(t, err)
	assert.Nil(t, mentions)
}

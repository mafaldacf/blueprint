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

var userIDServiceRegistry = registry.NewServiceRegistry[socialnetwork.UserIDService]("userid_service")

func init() {
	userIDServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.UserIDService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewUserIDServiceImpl(ctx, cache, db)
	})
}

func TestUserIDServiceGetUserIdNotFound(t *testing.T) {
	ctx := context.Background()
	service, err := userIDServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.GetUserId(ctx, 0, "nonexistentuser")
	assert.Error(t, err)
}

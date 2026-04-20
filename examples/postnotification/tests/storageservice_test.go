package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var storageServiceRegistry = registry.NewServiceRegistry[postnotification.StorageService]("storage_service")

func init() {
	storageServiceRegistry.Register("local", func(ctx context.Context) (postnotification.StorageService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return postnotification.NewStorageServiceImpl(ctx, db)
	})
}

func TestStorageServiceStorePost(t *testing.T) {
	ctx := context.Background()
	storageService, err := storageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	postID, err := storageService.StorePost(ctx, 1, "Hello, world!")
	assert.NoError(t, err)
	assert.NotZero(t, postID)
}

func TestStorageServiceReadPost(t *testing.T) {
	ctx := context.Background()
	storageService, err := storageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Store a post first
	postID, err := storageService.StorePost(ctx, 1, "Read test post")
	assert.NoError(t, err)

	// Read it back
	post, err := storageService.ReadPost(ctx, 1, postID)
	assert.NoError(t, err)
	assert.Equal(t, postID, post.PostID)
	assert.Equal(t, "Read test post", post.Text)
}

func TestStorageServiceDeletePost(t *testing.T) {
	ctx := context.Background()
	storageService, err := storageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Store a post
	postID, err := storageService.StorePost(ctx, 1, "Post to delete")
	assert.NoError(t, err)

	// Delete it
	err = storageService.DeletePost(ctx, postID)
	assert.NoError(t, err)
}

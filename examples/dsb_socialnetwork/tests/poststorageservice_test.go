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

var postStorageServiceRegistry = registry.NewServiceRegistry[socialnetwork.PostStorageService]("post_storage_service")

func init() {
	postStorageServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.PostStorageService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewPostStorageServiceImpl(ctx, cache, db)
	})
}

func makePost(postID int64, userID int64, username string, text string) socialnetwork.Post {
	return socialnetwork.Post{
		PostID:  postID,
		Creator: socialnetwork.Creator{UserID: userID, Username: username},
		Text:    text,
	}
}

func TestPostStorageServiceStorePost(t *testing.T) {
	ctx := context.Background()
	service, err := postStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	post := makePost(1001, 1, "alice", "My first post!")
	err = service.StorePost(ctx, 0, post)
	assert.NoError(t, err)
}

func TestPostStorageServiceReadPost(t *testing.T) {
	ctx := context.Background()
	service, err := postStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	post := makePost(1002, 1, "alice", "A post to read back.")
	err = service.StorePost(ctx, 0, post)
	assert.NoError(t, err)

	result, err := service.ReadPost(ctx, 0, 1002)
	assert.NoError(t, err)
	assert.Equal(t, int64(1002), result.PostID)
	assert.Equal(t, "A post to read back.", result.Text)
	assert.Equal(t, "alice", result.Creator.Username)
}

func TestPostStorageServiceReadPostNotFound(t *testing.T) {
	ctx := context.Background()
	service, err := postStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.ReadPost(ctx, 0, 9999)
	assert.Error(t, err)
}

func TestPostStorageServiceReadPosts(t *testing.T) {
	ctx := context.Background()
	service, err := postStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	post1 := makePost(2001, 2, "bob", "Post one.")
	post2 := makePost(2002, 2, "bob", "Post two.")
	assert.NoError(t, service.StorePost(ctx, 0, post1))
	assert.NoError(t, service.StorePost(ctx, 0, post2))

	posts, err := service.ReadPosts(ctx, 0, []int64{2001, 2002})
	assert.NoError(t, err)
	assert.Len(t, posts, 2)
}

func TestPostStorageServiceReadPostsEmpty(t *testing.T) {
	ctx := context.Background()
	service, err := postStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	posts, err := service.ReadPosts(ctx, 0, []int64{})
	assert.NoError(t, err)
	assert.Empty(t, posts)
}

func TestPostStorageServiceReadPostsDuplicate(t *testing.T) {
	ctx := context.Background()
	service, err := postStorageServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.ReadPosts(ctx, 0, []int64{3001, 3001})
	assert.Error(t, err)
}

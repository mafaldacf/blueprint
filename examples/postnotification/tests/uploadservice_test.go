package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplequeue"
	"github.com/stretchr/testify/assert"
)

var uploadServiceRegistry = registry.NewServiceRegistry[postnotification.UploadService]("upload_service")

func init() {
	uploadServiceRegistry.Register("local", func(ctx context.Context) (postnotification.UploadService, error) {
		storageService, err := storageServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		queue, err := simplequeue.NewSimpleQueue(ctx)
		if err != nil {
			return nil, err
		}
		return postnotification.NewUploadServiceImpl(ctx, storageService, queue)
	})
}

func TestUploadServiceUploadPost(t *testing.T) {
	ctx := context.Background()
	uploadService, err := uploadServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	postID, err := uploadService.UploadPost(ctx, "testuser", "Hello, world!")
	assert.NoError(t, err)
	assert.NotZero(t, postID)
}

func TestUploadServiceDeletePost(t *testing.T) {
	ctx := context.Background()
	uploadService, err := uploadServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Upload a post first
	postID, err := uploadService.UploadPost(ctx, "testuser", "Post to delete")
	assert.NoError(t, err)

	// Delete it
	err = uploadService.DeletePost(ctx, postID)
	assert.NoError(t, err)
}

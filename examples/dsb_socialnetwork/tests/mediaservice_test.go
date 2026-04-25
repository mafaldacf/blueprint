package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_socialnetwork/workflow/socialnetwork"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/stretchr/testify/assert"
)

var mediaServiceRegistry = registry.NewServiceRegistry[socialnetwork.MediaService]("media_service")

func init() {
	mediaServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.MediaService, error) {
		return socialnetwork.NewMediaServiceImpl(ctx)
	})
}

func TestMediaServiceComposeEmpty(t *testing.T) {
	ctx := context.Background()
	service, err := mediaServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	medias, err := service.ComposeMedia(ctx, 0, []string{}, []int64{})
	assert.NoError(t, err)
	assert.Empty(t, medias)
}

func TestMediaServiceComposeSingle(t *testing.T) {
	ctx := context.Background()
	service, err := mediaServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	medias, err := service.ComposeMedia(ctx, 0, []string{"photo"}, []int64{1001})
	assert.NoError(t, err)
	assert.Len(t, medias, 1)
	assert.Equal(t, int64(1001), medias[0].MediaID)
	assert.Equal(t, "photo", medias[0].MediaType)
}

func TestMediaServiceComposeMultiple(t *testing.T) {
	ctx := context.Background()
	service, err := mediaServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	medias, err := service.ComposeMedia(ctx, 0, []string{"photo", "video"}, []int64{2001, 2002})
	assert.NoError(t, err)
	assert.Len(t, medias, 2)
	assert.Equal(t, int64(2001), medias[0].MediaID)
	assert.Equal(t, "photo", medias[0].MediaType)
	assert.Equal(t, int64(2002), medias[1].MediaID)
	assert.Equal(t, "video", medias[1].MediaType)
}

func TestMediaServiceComposeMismatchedLengths(t *testing.T) {
	ctx := context.Background()
	service, err := mediaServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.ComposeMedia(ctx, 0, []string{"photo", "video"}, []int64{1001})
	assert.Error(t, err)
}

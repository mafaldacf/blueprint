package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_socialnetwork/workflow/socialnetwork"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var urlShortenServiceRegistry = registry.NewServiceRegistry[socialnetwork.UrlShortenService]("urlshorten_service")

func init() {
	urlShortenServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.UrlShortenService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewUrlShortenServiceImpl(ctx, db)
	})
}

func TestUrlShortenServiceComposeEmpty(t *testing.T) {
	ctx := context.Background()
	service, err := urlShortenServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	urls, err := service.ComposeUrls(ctx, 0, []string{})
	assert.NoError(t, err)
	assert.Empty(t, urls)
}

func TestUrlShortenServiceComposeSingleUrl(t *testing.T) {
	ctx := context.Background()
	service, err := urlShortenServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	urls, err := service.ComposeUrls(ctx, 0, []string{"http://example.com/some/long/path"})
	assert.NoError(t, err)
	assert.Len(t, urls, 1)
	assert.Equal(t, "http://example.com/some/long/path", urls[0].ExpandedUrl)
	assert.NotEmpty(t, urls[0].ShortenedUrl)
	assert.NotEqual(t, urls[0].ExpandedUrl, urls[0].ShortenedUrl)
}

func TestUrlShortenServiceComposeMultipleUrls(t *testing.T) {
	ctx := context.Background()
	service, err := urlShortenServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	rawUrls := []string{"http://example.com", "https://another.org/page"}
	urls, err := service.ComposeUrls(ctx, 0, rawUrls)
	assert.NoError(t, err)
	assert.Len(t, urls, 2)
	assert.Equal(t, rawUrls[0], urls[0].ExpandedUrl)
	assert.Equal(t, rawUrls[1], urls[1].ExpandedUrl)
	assert.NotEqual(t, urls[0].ShortenedUrl, urls[1].ShortenedUrl)
}

func TestUrlShortenServiceGetExtendedUrls(t *testing.T) {
	ctx := context.Background()
	service, err := urlShortenServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	extended, err := service.GetExtendedUrls(ctx, 0, []string{"http://short-url/abcdefghij"})
	assert.NoError(t, err)
	assert.Empty(t, extended)
}

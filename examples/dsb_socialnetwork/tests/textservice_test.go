package tests

import (
	"context"
	"strings"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_socialnetwork/workflow/socialnetwork"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/stretchr/testify/assert"
)

var textServiceRegistry = registry.NewServiceRegistry[socialnetwork.TextService]("text_service")

func init() {
	textServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.TextService, error) {
		urlShortenService, err := urlShortenServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		userMentionService, err := userMentionServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewTextServiceImpl(ctx, urlShortenService, userMentionService)
	})
}

func TestTextServiceComposePlainText(t *testing.T) {
	ctx := context.Background()
	service, err := textServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	text, mentions, urls, err := service.ComposeText(ctx, 0, "Hello world, no urls or mentions here.")
	assert.NoError(t, err)
	assert.Equal(t, "Hello world, no urls or mentions here.", text)
	assert.Nil(t, mentions)
	assert.Empty(t, urls)
}

func TestTextServiceComposeTextWithUrl(t *testing.T) {
	ctx := context.Background()
	service, err := textServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// The URL regex does not include '/' in the path character class, so it captures up to the domain only.
	text, mentions, urls, err := service.ComposeText(ctx, 0, "Check this out: http://example.com")
	assert.NoError(t, err)
	assert.Nil(t, mentions)
	assert.Len(t, urls, 1)
	assert.Equal(t, "http://example.com", urls[0].ExpandedUrl)
	assert.NotContains(t, text, "http://example.com")
	assert.True(t, strings.Contains(text, urls[0].ShortenedUrl))
}

func TestTextServiceComposeTextWithMultipleUrls(t *testing.T) {
	ctx := context.Background()
	service, err := textServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	text, _, urls, err := service.ComposeText(ctx, 0, "Visit http://foo.com and https://bar.org/page")
	assert.NoError(t, err)
	assert.Len(t, urls, 2)
	assert.True(t, strings.Contains(text, urls[0].ShortenedUrl))
	assert.True(t, strings.Contains(text, urls[1].ShortenedUrl))
}

func TestTextServiceComposeTextWithMention(t *testing.T) {
	ctx := context.Background()
	service, err := textServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	text, mentions, urls, err := service.ComposeText(ctx, 0, "Hello @alice, how are you?")
	assert.NoError(t, err)
	assert.Equal(t, "Hello @alice, how are you?", text)
	assert.Nil(t, mentions)
	assert.Empty(t, urls)
}

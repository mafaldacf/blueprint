package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/stretchr/testify/assert"
)

var newsServiceRegistry = registry.NewServiceRegistry[trainticket.NewsService]("news_service")

func init() {
	newsServiceRegistry.Register("local", func(ctx context.Context) (trainticket.NewsService, error) {
		return trainticket.NewNewsServiceImpl(ctx)
	})
}

func TestNewsServiceHello(t *testing.T) {
	ctx := context.Background()
	service, err := newsServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	result, err := service.Hello(ctx, "test")
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

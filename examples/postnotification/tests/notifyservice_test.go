package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplequeue"
	"github.com/stretchr/testify/assert"
)

var notifyServiceRegistry = registry.NewServiceRegistry[postnotification.NotifyService]("notify_service")

func init() {
	notifyServiceRegistry.Register("local", func(ctx context.Context) (postnotification.NotifyService, error) {
		storageService, err := storageServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		queue, err := simplequeue.NewSimpleQueue(ctx)
		if err != nil {
			return nil, err
		}
		return postnotification.NewNotifyServiceImpl(ctx, storageService, queue)
	})
}

func TestNotifyServiceRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	notifyService, err := notifyServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	done := make(chan error)
	go func() {
		done <- notifyService.Run(ctx)
	}()

	cancel()
	err = <-done
	assert.NoError(t, err)
}

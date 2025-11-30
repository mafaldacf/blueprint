package postnotification_simple

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type NotifyService interface {
	Run(ctx context.Context) error
}

type NotifyServiceImpl struct {
	storageService     StorageService
	notificationsQueue backend.Queue
	exit_on_error      bool
}

func NewNotifyServiceImpl(ctx context.Context, storageService StorageService, notificationsQueue backend.Queue) (NotifyService, error) {
	n := &NotifyServiceImpl{storageService: storageService, notificationsQueue: notificationsQueue, exit_on_error: false}
	return n, nil
}

func (n *NotifyServiceImpl) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			var workerMessage Message
			ok, err := n.notificationsQueue.Pop(ctx, &workerMessage)
			if err != nil && n.exit_on_error {
				return err
			}
			if !ok {
				continue
			}
			n.storageService.ReadPost(ctx, workerMessage.ReqID, workerMessage.PostID)
		}
	}
	return nil
}

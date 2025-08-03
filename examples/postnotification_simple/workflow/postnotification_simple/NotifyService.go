package postnotification_simple

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"

	//"github.com/blueprint-uservices/blueprint/examples/postnotification_simple/workflow/postnotification_simple/common"
)

// does not expose any methods to other services
// it defines Run that runs workers that pull messages from the notificationsQueue
type NotifyService interface {
	Run(ctx context.Context) error
	/* Notify(ctx context.Context, message Message) error */
}

type NotifyServiceImpl struct {
	storageService     StorageService
	notificationsQueue backend.Queue
	num_workers        int
}

func NewNotifyServiceImpl(ctx context.Context, storageService StorageService, notificationsQueue backend.Queue) (NotifyService, error) {
	n := &NotifyServiceImpl{storageService: storageService, notificationsQueue: notificationsQueue, num_workers: 4}
	return n, nil
}

/* func (n *NotifyServiceImpl) workerThread(ctx context.Context) error {
	//var forever chan struct{}
	//go func() {
	var workerMessage Message
	n.notificationsQueue.Pop(ctx, &workerMessage)
	n.storageService.ReadPost(ctx, workerMessage.ReqID, workerMessage.PostID)
	//}()
	//<-forever
	return nil
} */

func (n *NotifyServiceImpl) Run(ctx context.Context) error {
	backend.GetLogger().Info(ctx, "initializing %d workers", n.num_workers)

	var workerMessage Message
	n.notificationsQueue.Pop(ctx, &workerMessage)
	n.storageService.ReadPost(ctx, workerMessage.ReqID, workerMessage.PostID)

	/* var wg sync.WaitGroup
	wg.Add(n.num_workers)

	fn := func(i int) {
		defer wg.Done()
		err := n.workerThread(ctx)
		if err != nil {
			backend.GetLogger().Error(ctx, "error in worker thread (%d): %s", i, err.Error())
			panic(err)
		}
	}

	for i := 1; i <= n.num_workers; i++ { // for
		go fn(i)
	}
	wg.Wait()
	backend.GetLogger().Info(ctx, "joining %d workers", n.num_workers) */
	return nil
}

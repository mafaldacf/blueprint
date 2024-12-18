package postnotification

import (
	"context"
	"sync"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification/common"
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

/* func (n *NotifyServiceImpl) Notify(ctx context.Context, message Message) error {

	reqID, err := common.StringToInt64(message.ReqID)
	if err != nil {
		return nil
	}
	postID, err := common.StringToInt64(message.PostID)
	if err != nil {
		return nil
	}
	_, err = n.storageService.ReadPost(ctx, reqID, postID)
	return err
} */

func (n *NotifyServiceImpl) handleMessage(ctx context.Context, message Message) error {
	reqID, err := common.StringToInt64(message.ReqID)
	if err != nil {
		return err
	}
	postID_NOTIFY_SVC, err := common.StringToInt64(message.PostID)
	if err != nil {
		return err
	}

	_, _, err = n.storageService.ReadPostNoSQL(ctx, reqID, postID_NOTIFY_SVC)
	//_, err = n.storageService.ReadPost(ctx, reqID, postID)
	if err != nil {
		return err
	}
	return nil
}

/* func (n *NotifyServiceImpl) workerThread(ctx context.Context) error {
	var forever chan struct{}
	go func() {
		var notifyEvent map[string]interface{}
		n.notificationsQueue.Pop(ctx, &notifyEvent)
		workerMessage := Message{
			ReqID:     notifyEvent["ReqID"].(string),
			PostID:    notifyEvent["PostID"].(string),
			Timestamp: notifyEvent["Timestamp"].(string),
		}
		//reqID, _ := common.StringToInt64(notification.ReqID)
		//postID, _ := common.StringToInt64(notification.PostID)
		//n.storageService.ReadPost(ctx, reqID, postID)
		n.handleMessage(ctx, workerMessage)
	}()
	<-forever
	return nil
} */

func (n *NotifyServiceImpl) workerThread(ctx context.Context) error {
	var forever chan struct{}
	go func() {
		var notifyEvent Message
		n.notificationsQueue.Pop(ctx, &notifyEvent)
		n.handleMessage(ctx, notifyEvent)
	}()
	<-forever
	return nil
}

/* func (n *NotifyServiceImpl) workerThread(ctx context.Context, workerID int) error {
	var forever chan struct{}
	go func() {
		var message map[string]interface{}
		backend.GetLogger().Info(ctx, "[worker %d] waiting for message...", workerID)

		// blueprint uses backend.CopyResult in backend.Pop that requires as source argument
		// an interface that is converted to map[string]interface after retrieving the element
		// from the notificationsQueue, and dst (message) argument needs to match the source, otherwise we get:
		// an ERROR: "unable to copy incompatible types map[string]interface {} and postnotification.Message"
		// we also must use values as strings in the message otherwise convertion outputs incorrect values (due to float?)
		result, err := n.notificationsQueue.Pop(ctx, &message)
		backend.GetLogger().Info(ctx, "[worker %d] received message %w", workerID, message)
		if err != nil {
			backend.GetLogger().Error(ctx, "error retrieving message from notificationsQueue: %s", err.Error())
			time.Sleep(1 * time.Second)
			return
		}
		if !result {
			backend.GetLogger().Error(ctx, "could not retrieve message from notificationsQueue")
			return
		}
		notification := Message {
			ReqID: message["ReqID"].(string),
			PostID: message["PostID"].(string),
			Timestamp: message["Timestamp"].(string),
		}
		reqID, err := common.StringToInt64(notification.ReqID)
		if err != nil {
			return
		}
		postID, err := common.StringToInt64(notification.PostID)
		if err != nil {
			return
		}
		_, err = n.storageService.ReadPost(ctx, reqID, postID)
		if err != nil {
			return
		}
		err = n.handleMessage(ctx, notification)
		if err != nil {
			return
		}
	}()
	<-forever
	return nil
} */

func (n *NotifyServiceImpl) Run(ctx context.Context) error {
	backend.GetLogger().Info(ctx, "initializing %d workers", n.num_workers)
	var wg sync.WaitGroup
	wg.Add(n.num_workers)
	for i := 1; i <= n.num_workers; i++ {
		go func(i int) {
			defer wg.Done()
			err := n.workerThread(ctx)
			if err != nil {
				backend.GetLogger().Error(ctx, "error in worker thread (%d): %s", i, err.Error())
				panic(err)
			}
		}(i)
	}
	wg.Wait()
	backend.GetLogger().Info(ctx, "joining %d workers", n.num_workers)
	return nil
}

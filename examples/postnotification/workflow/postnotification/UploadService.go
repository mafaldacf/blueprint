package postnotification

import (
	"context"
	"math/rand"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification/common"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type UploadService interface {
	UploadPost(ctx context.Context, username string, text string) (int64, error)
}

type UploadServiceImpl struct {
	storageService  StorageService
	/* notifyService  	NotifyService */
	queue    		backend.Queue
}

func NewUploadServiceImpl(ctx context.Context, storageService StorageService, queue backend.Queue) (UploadService, error) {
	return &UploadServiceImpl{storageService: storageService, /*  notifyService: notifyService */ queue: queue}, nil
}

func (u *UploadServiceImpl) UploadPost(ctx context.Context, username string, text string) (int64, error) {
	reqID := rand.Int63()
	post := Post{
		ReqID:     reqID,
		PostID:    rand.Int63(),
		Text:      text,
	}
	u.storageService.StorePost(ctx, post.ReqID, post)

	message := Message{
		ReqID:     common.Int64ToString(post.ReqID),
		PostID:    common.Int64ToString(post.PostID),
	}
	/* err := u.notifyService.Notify(ctx, message) */
	_, err := u.queue.Push(ctx, message)
	if err != nil {
		return 0, err
	}

	return post.PostID, nil
}

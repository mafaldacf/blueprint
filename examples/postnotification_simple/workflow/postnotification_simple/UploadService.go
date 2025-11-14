package postnotification_simple

import (
	"context"
	"math/rand"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type UploadService interface {
	UploadPost(ctx context.Context, username string, text string) (int64, error)
	DeletePost(ctx context.Context, postID int64) error
}

type UploadServiceImpl struct {
	storageService     StorageService
	notificationsQueue backend.Queue
}

func NewUploadServiceImpl(ctx context.Context, storageService StorageService, notificationsQueue backend.Queue) (UploadService, error) {
	return &UploadServiceImpl{storageService: storageService, notificationsQueue: notificationsQueue}, nil
}

func (u *UploadServiceImpl) UploadPost(ctx context.Context, username string, text string) (int64, error) {
	reqID := rand.Int63()
	postID, _ := u.storageService.StorePost(ctx, reqID, text)

	message := Message{
		ReqID:  reqID,
		PostID: postID,
	}
	_, err := u.notificationsQueue.Push(ctx, message)
	if err != nil {
		return 0, err
	}
	return postID, nil
}

func (u *UploadServiceImpl) DeletePost(ctx context.Context, postID int64) error {
	return u.storageService.DeletePost(ctx, postID)
}

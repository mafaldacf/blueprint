package postnotification_simple

import (
	"context"
	"math/rand"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type UploadService interface {
	UploadPost(ctx context.Context, username string, text string) (int64, error)
}

type UploadServiceImpl struct {
	storageService     StorageService
	notificationsQueue backend.Queue
}

func NewUploadServiceImpl(ctx context.Context, storageService StorageService, notificationsQueue backend.Queue) (UploadService, error) {
	return &UploadServiceImpl{storageService: storageService, notificationsQueue: notificationsQueue}, nil
}

/* func (u *UploadServiceImpl) UploadPost(ctx context.Context, username string, text string) (int64, error) {
	reqID := rand.Int63()
	postID_UploadSVC := rand.Int63()

	timestamp := rand.Int63()
	mentions := []string{"alice", "bob"}
	post := Post{
		ReqID:     reqID,
		PostID:    postID_UploadSVC,
		Text:      text,
		Mentions:  mentions,
		Timestamp: timestamp,
		Creator: Creator{
			Username: "some username",
		},
	}
	u.storageService.StorePost(ctx, post.ReqID, post)

	message := Message{
		ReqID:          post.ReqID,
		PostID_MESSAGE: post.PostID,
	}
	_, err := u.notificationsQueue.Push(ctx, message)
	if err != nil {
		return 0, err
	}
	return post.PostID, nil
} */

func (u *UploadServiceImpl) UploadPost(ctx context.Context, username string, text string) (int64, error) {
	reqID := rand.Int63()

	postID_UploadSVC, _ := u.storageService.StorePost(ctx, reqID, text)

	message := Message{
		ReqID:          reqID,
		PostID_MESSAGE: postID_UploadSVC,
	}
	_, err := u.notificationsQueue.Push(ctx, message)
	if err != nil {
		return 0, err
	}
	return postID_UploadSVC, nil
}

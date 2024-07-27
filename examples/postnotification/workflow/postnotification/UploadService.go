package postnotification

import (
	"context"
	"math/rand"

	bp_backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification/common"
)

type UploadService interface {
	UploadPost(ctx context.Context, username string, text string) (int64, error)
}

type UploadServiceImpl struct {
	storage_service     StorageService
	notifications_queue bp_backend.Queue
	/* notify_service  	NotifyService  */
}

func NewUploadServiceImpl(ctx context.Context, storage_service StorageService, notifications_queue bp_backend.Queue) (UploadService, error) {
	return &UploadServiceImpl{storage_service: storage_service /*  notify_service: notify_service */, notifications_queue: notifications_queue}, nil
}

func (u *UploadServiceImpl) UploadPost(ctx context.Context, username string, text string) (int64, error) {
	reqID := rand.Int63()
	postID := rand.Int63()
	timestamp := rand.Int63()
	mentions := []string{"alice", "bob"}
	post := Post{
		ReqID:     reqID,
		PostID:    postID,
		Text:      text,
		Mentions:  mentions,
		Timestamp: timestamp,
		Creator: Creator{
			Username: "some username",
		},
	}
	u.storage_service.StorePostNoSQL(ctx, post.ReqID, post)
	//u.storage_service.StorePost(ctx, post.ReqID, post)

	message := Message{
		ReqID:  common.Int64ToString(post.ReqID),
		PostID: common.Int64ToString(post.PostID),
	}
	/* err := u.notify_service.Notify(ctx, message) */
	_, err := u.notifications_queue.Push(ctx, message)
	if err != nil {
		return 0, err
	}

	return post.PostID, nil
}

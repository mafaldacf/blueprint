package postnotification

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification/common"
)

type UploadService interface {
	UploadPost(ctx context.Context, username string, text string) (int64, error)
	ReadPostMedia(ctx context.Context, reqID int64, postID int64) (Media, error)
}

type UploadServiceImpl struct {
	storage_service     StorageService
	media_service       MediaService
	notifications_queue backend.Queue
	timeline_cache      backend.Cache
}

func NewUploadServiceImpl(ctx context.Context, storage_service StorageService, media_service MediaService, notifications_queue backend.Queue, timeline_cache backend.Cache) (UploadService, error) {
	return &UploadServiceImpl{storage_service: storage_service, media_service: media_service, notifications_queue: notifications_queue, timeline_cache: timeline_cache}, nil
}

func (u *UploadServiceImpl) ReadPostMedia(ctx context.Context, reqID int64, postID int64) (Media, error) {
	var media Media
	media, _ = u.storage_service.ReadMedia(ctx, reqID, postID)
	return media, nil
}

func (u *UploadServiceImpl) UploadPost(ctx context.Context, username string, text string) (int64, error) {
	reqID := rand.Int63()
	postID := rand.Int63()
	mediaID := rand.Int63()

	media := Media{
		PostID:  postID,
		MediaID: mediaID,
		Content: common.HELLO_WORLD_CONST,
	}
	u.media_service.StoreMedia(ctx, media)

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
	u.storage_service.StorePostCache(ctx, post.ReqID, post)
	u.storage_service.StorePostNoSQL(ctx, post.ReqID, post)

	message := Message{
		ReqID:  common.Int64ToString(post.ReqID),
		PostID: common.Int64ToString(post.PostID),
	}
	_, err := u.notifications_queue.Push(ctx, message)
	if err != nil {
		return 0, err
	}

	reqIDStr := strconv.FormatInt(reqID, 10)
	timeline := Timeline{
		ReqID:  reqID,
		PostID: postID,
	}
	return post.PostID, u.timeline_cache.Put(ctx, reqIDStr, timeline)
}

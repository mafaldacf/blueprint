package postnotification

import (
	"context"
	"math/rand"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification/common"
)

type UploadService interface {
	UploadPost(ctx context.Context, username string, text string) (int64, error)
	ReadMedia(ctx context.Context, postID int64) (Media, error)
}

type UploadServiceImpl struct {
	storageService     StorageService
	mediaService       MediaService
	notificationsQueue backend.Queue
	timelineCache      backend.Cache
}

func NewUploadServiceImpl(ctx context.Context, storageService StorageService, mediaService MediaService, notificationsQueue backend.Queue, timelineCache backend.Cache) (UploadService, error) {
	return &UploadServiceImpl{storageService: storageService, mediaService: mediaService, notificationsQueue: notificationsQueue, timelineCache: timelineCache}, nil
}

func (u *UploadServiceImpl) ReadMedia(ctx context.Context, postID int64) (Media, error) {
	reqID := rand.Int63()

	var media Media
	media, _ = u.storageService.ReadPostMedia(ctx, reqID, postID)
	return media, nil
}

func (u *UploadServiceImpl) UploadPost(ctx context.Context, username string, text string) (int64, error) {
	reqID := rand.Int63()

	media := Media{
		Content: common.HELLO_WORLD_CONST,
	}
	mediaID, _ := u.mediaService.StoreMedia(ctx, media)

	timestamp := rand.Int63()
	mentions := []string{"alice", "bob"}
	post := Post{
		ReqID:     reqID,
		MediaID:   mediaID,
		Text:      text,
		Mentions:  mentions,
		Timestamp: timestamp,
		Creator: Creator{
			Username: "some username",
		},
	}

	//postID, _ := u.storageService.StorePostCache(ctx, reqID, post)
	postID, _ := u.storageService.StorePostNoSQL(ctx, reqID, post)

	reqIDStr := common.Int64ToString(reqID)
	message := Message{
		ReqID:     reqIDStr,
		PostID:    common.Int64ToString(postID),
		Timestamp: common.Int64ToString(timestamp),
	}

	_, err := u.notificationsQueue.Push(ctx, message)
	if err != nil {
		return 0, err
	}

	timeline := Timeline{
		ReqID:  reqID,
		PostID: postID,
	}
	return postID, u.timelineCache.Put(ctx, reqIDStr, timeline)
}

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
	storageService     StorageService
	mediaService       MediaService
	notificationsQueue backend.Queue
	timelineCache      backend.Cache
}

func NewUploadServiceImpl(ctx context.Context, storageService StorageService, mediaService MediaService, notificationsQueue backend.Queue, timelineCache backend.Cache) (UploadService, error) {
	return &UploadServiceImpl{storageService: storageService, mediaService: mediaService, notificationsQueue: notificationsQueue, timelineCache: timelineCache}, nil
}

func (u *UploadServiceImpl) ReadPostMedia(ctx context.Context, reqID int64, postID int64) (Media, error) {
	var media Media
	media, _ = u.storageService.ReadMedia(ctx, reqID, postID)
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
	u.mediaService.StoreMedia(ctx, media)

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
	u.storageService.StorePostCache(ctx, post.ReqID, post)
	u.storageService.StorePostNoSQL(ctx, post.ReqID, post)

	message := Message{
		ReqID:  common.Int64ToString(post.ReqID),
		PostID: common.Int64ToString(post.PostID),
	}
	_, err := u.notificationsQueue.Push(ctx, message)
	if err != nil {
		return 0, err
	}

	reqIDStr := strconv.FormatInt(reqID, 10)
	timeline := Timeline{
		ReqID:  reqID,
		PostID: postID,
	}
	return post.PostID, u.timelineCache.Put(ctx, reqIDStr, timeline)
}

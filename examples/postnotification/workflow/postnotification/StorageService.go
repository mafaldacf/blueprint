package postnotification

import (
	"context"
	"strconv"
	
	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type StorageService interface {
	StorePost(ctx context.Context, reqID int64, post Post) error
	ReadPost(ctx context.Context, reqID int64, postID int64) (Post, error)
}

type StorageServiceImpl struct {
	cache 		 backend.Cache
	/* dummyService DummyService */
}

func NewStorageServiceImpl(ctx context.Context, cache backend.Cache/* , dummyService DummyService */) (StorageService, error) {
	u := &StorageServiceImpl{cache: cache, /* dummyService: dummyService */}
	return u, nil
}

func (u *StorageServiceImpl) StorePost(ctx context.Context, reqID int64, post Post) error {
	postIDStr := strconv.FormatInt(post.PostID, 10)
	err := u.cache.Put(ctx, postIDStr, post)
	if err != nil {
		return err
	}
	return nil
}

func (u *StorageServiceImpl) ReadPost(ctx context.Context, reqID int64, postID int64) (Post, error) {
	var post Post
	postIDStr := strconv.FormatInt(postID, 10)
	_, err := u.cache.Get(ctx, postIDStr, &post)
	if err != nil {
		return post, err
	}
	/* u.dummyService.Dummy(ctx, postID) */
	return post, nil
}

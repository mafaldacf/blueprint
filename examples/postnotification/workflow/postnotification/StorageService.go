package postnotification

import (
	"context"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification/common"
)

type StorageService interface {
	StorePostCache(ctx context.Context, reqID int64, post Post) error
	StorePostNoSQL(ctx context.Context, reqID int64, post Post) error
	ReadPostCache(ctx context.Context, reqID int64, postID int64) (Post, error)
	ReadPostNoSQL(ctx context.Context, reqID int64, postID int64) (Post, Analytics, error)
	ReadMedia(ctx context.Context, reqID int64, postID int64) (Media, error)
}

type StorageServiceImpl struct {
	analyticsService AnalyticsService
	mediaService     MediaService
	posts_cache       backend.Cache
	postsDb          backend.NoSQLDatabase
	analyticsQueue   backend.Queue
}

func NewStorageServiceImpl(ctx context.Context, analyticsService AnalyticsService, mediaService MediaService, posts_cache backend.Cache, postsDb backend.NoSQLDatabase, analyticsQueue backend.Queue) (StorageService, error) {
	s := &StorageServiceImpl{analyticsService: analyticsService, mediaService: mediaService, posts_cache: posts_cache, postsDb: postsDb, analyticsQueue: analyticsQueue}
	return s, nil
}

func (s *StorageServiceImpl) ReadMedia(ctx context.Context, reqID int64, postID int64) (Media, error) {
	var post Post

	postIDStr := strconv.FormatInt(postID, 10)
	s.posts_cache.Get(ctx, postIDStr, &post)

	var media Media
	mediaID := post.PostID
	media, _ = s.mediaService.ReadMedia(ctx, mediaID)
	return media, nil
}

func (s *StorageServiceImpl) StorePostCache(ctx context.Context, reqID int64, post Post) error {
	postIDStr := strconv.FormatInt(post.PostID, 10)
	return s.posts_cache.Put(ctx, postIDStr, post)
}

func (s *StorageServiceImpl) StorePostNoSQL(ctx context.Context, reqID int64, post Post) error {
	collection, err := s.postsDb.GetCollection(ctx, "post", "post")
	if err != nil {
		return err
	}
	err = collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}
	message := TriggerAnalyticsMessage{
		PostID: common.Int64ToString(post.PostID),
	}
	_, err = s.analyticsQueue.Push(ctx, message)
	return err
}

func (s *StorageServiceImpl) ReadPostCache(ctx context.Context, reqID int64, postID int64) (Post, error) {
	var post Post
	postIDStr := strconv.FormatInt(postID, 10)
	_, err := s.posts_cache.Get(ctx, postIDStr, &post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (s *StorageServiceImpl) ReadPostNoSQL(ctx context.Context, reqID int64, postID int64) (Post, Analytics, error) {
	var post Post
	var analytics Analytics
	collection, err := s.postsDb.GetCollection(ctx, "post", "post")
	if err != nil {
		return post, analytics, err
	}
	query := bson.D{{Key: "postid", Value: postID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return post, analytics, err
	}
	res, err := result.One(ctx, &post)
	if !res || err != nil {
		return post, analytics, err
	}
	analytics, err = s.analyticsService.ReadAnalytics(ctx, postID)
	return post, analytics, err
}

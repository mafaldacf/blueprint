package postnotification

import (
	"context"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/blueprint-uservices/blueprint/examples/postnotification/workflow/postnotification/common"
)

type StorageService interface {
	StorePost(ctx context.Context, reqID int64, post Post) error
	StorePostNoSQL(ctx context.Context, reqID int64, post Post) error
	ReadPost(ctx context.Context, reqID int64, postID int64) (Post, error)
	ReadPostNoSQL(ctx context.Context, reqID int64, postID int64) (Post, Analytics, error)
}

type StorageServiceImpl struct {
	analytics_service AnalyticsService
	posts_cache       backend.Cache
	posts_db          backend.NoSQLDatabase
	analytics_queue   backend.Queue
}

func NewStorageServiceImpl(ctx context.Context, analytics_service AnalyticsService, posts_cache backend.Cache, posts_db backend.NoSQLDatabase, analytics_queue backend.Queue) (StorageService, error) {
	s := &StorageServiceImpl{analytics_service: analytics_service, posts_cache: posts_cache, posts_db: posts_db, analytics_queue: analytics_queue}
	return s, nil
}

func (s *StorageServiceImpl) StorePost(ctx context.Context, reqID int64, post Post) error {
	postIDStr := strconv.FormatInt(post.PostID, 10)
	return s.posts_cache.Put(ctx, postIDStr, post)
}

func (s *StorageServiceImpl) StorePostNoSQL(ctx context.Context, reqID int64, post Post) error {
	collection, err := s.posts_db.GetCollection(ctx, "post", "post")
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
	/* err := u.notifyService.Notify(ctx, message) */
	_, err = s.analytics_queue.Push(ctx, message)
	return err
}

func (s *StorageServiceImpl) ReadPost(ctx context.Context, reqID int64, postID int64) (Post, error) {
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
	collection, err := s.posts_db.GetCollection(ctx, "post", "post")
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
	analytics, err = s.analytics_service.ReadAnalytics(ctx, postID)
	return post, analytics, err
}

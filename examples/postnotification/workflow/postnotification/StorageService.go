package postnotification

import (
	"context"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type StorageService interface {
	StorePost(ctx context.Context, reqID int64, post Post) error
	StorePostNoSQL(ctx context.Context, reqID int64, post Post) error
	ReadPost(ctx context.Context, reqID int64, postID int64) (Post, error)
	ReadPostNoSQL(ctx context.Context, reqID int64, postID int64) (Post, error)
}

type StorageServiceImpl struct {
	cache 	backend.Cache
	db      backend.NoSQLDatabase
}

func NewStorageServiceImpl(ctx context.Context, cache backend.Cache, db backend.NoSQLDatabase) (StorageService, error) {
	s := &StorageServiceImpl{cache: cache, db: db}
	return s, nil
}

func (s *StorageServiceImpl) StorePost(ctx context.Context, reqID int64, post Post) error {
	postIDStr := strconv.FormatInt(post.PostID, 10)
	return s.cache.Put(ctx, postIDStr, post)
}

func (p *StorageServiceImpl) StorePostNoSQL(ctx context.Context, reqID int64, post Post) error {
	collection, err := p.db.GetCollection(ctx, "post", "post")
	if err != nil {
		return err
	}
	return collection.InsertOne(ctx, post)
}

func (s *StorageServiceImpl) ReadPost(ctx context.Context, reqID int64, postID int64) (Post, error) {
	var post Post
	postIDStr := strconv.FormatInt(postID, 10)
	_, err := s.cache.Get(ctx, postIDStr, &post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (s *StorageServiceImpl) ReadPostNoSQL(ctx context.Context, reqID int64, postID int64) (Post, error) {
	var post Post
	collection, err := s.db.GetCollection(ctx, "post", "post")
	if err != nil {
		return post, err
	}
	query := bson.D{{Key: "postid", Value: postID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return post, err
	}
	res, err := result.One(ctx, &post)
	if !res || err != nil {
		return post, err
	}
	return post, nil
}

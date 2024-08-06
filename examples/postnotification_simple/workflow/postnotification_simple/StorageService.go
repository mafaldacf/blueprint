package postnotification_simple

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type StorageService interface {
	StorePost(ctx context.Context, reqID int64, post Post) error
	ReadPost(ctx context.Context, reqID int64, postID int64) (Post, error)
}

type StorageServiceImpl struct {
	postsDb          backend.NoSQLDatabase
}

func NewStorageServiceImpl(ctx context.Context, postsDb backend.NoSQLDatabase) (StorageService, error) {
	s := &StorageServiceImpl{postsDb: postsDb}
	return s, nil
}

func (s *StorageServiceImpl) StorePost(ctx context.Context, reqID int64, post Post) error {
	collection, err := s.postsDb.GetCollection(ctx, "post", "post")
	if err != nil {
		return err
	}
	err = collection.InsertOne(ctx, post)
	return err
}

func (s *StorageServiceImpl) ReadPost(ctx context.Context, reqID int64, postID int64) (Post, error) {
	var post Post
	collection, err := s.postsDb.GetCollection(ctx, "post", "post")
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
	return post, err
}

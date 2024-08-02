package postnotification

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type MediaService interface {
	StoreMedia(ctx context.Context, media Media) error
	ReadMedia(ctx context.Context, mediaID int64) (Media, error)
}

type MediaServiceImpl struct {
	media_db backend.NoSQLDatabase
}

func NewMediaServiceImpl(ctx context.Context, media_db backend.NoSQLDatabase) (MediaService, error) {
	s := &MediaServiceImpl{media_db: media_db}
	return s, nil
}

func (s *MediaServiceImpl) StoreMedia(ctx context.Context, media Media) error {
	collection, err := s.media_db.GetCollection(ctx, "media", "media")
	if err != nil {
		return err
	}
	return collection.InsertOne(ctx, media)
}


func (s *MediaServiceImpl) ReadMedia(ctx context.Context, mediaID int64) (Media, error) {
	var media Media
	collection, err := s.media_db.GetCollection(ctx, "media", "media")
	if err != nil {
		return media, err
	}
	query := bson.D{{Key: "mediaid", Value: mediaID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return media, err
	}
	_, err = result.One(ctx, &media)
	return media, err
}

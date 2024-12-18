package postnotification

import (
	"context"
	"math/rand"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type MediaService interface {
	StoreMedia(ctx context.Context, media Media) (int64, error)
	ReadMedia(ctx context.Context, mediaID int64) (Media, error)
}

type MediaServiceImpl struct {
	mediaDb backend.NoSQLDatabase
}

func NewMediaServiceImpl(ctx context.Context, mediaDb backend.NoSQLDatabase) (MediaService, error) {
	s := &MediaServiceImpl{mediaDb: mediaDb}
	return s, nil
}

func (s *MediaServiceImpl) StoreMedia(ctx context.Context, media Media) (int64, error) {
	mediaID := rand.Int63()
	media.MediaID = mediaID
	collection, err := s.mediaDb.GetCollection(ctx, "media", "media")
	if err != nil {
		return mediaID, err
	}
	return mediaID, collection.InsertOne(ctx, media)
}

func (s *MediaServiceImpl) ReadMedia(ctx context.Context, mediaID int64) (Media, error) {
	var media Media
	collection, err := s.mediaDb.GetCollection(ctx, "media", "media")
	if err != nil {
		return media, err
	}
	mediaQuery := bson.D{{Key: "mediaid", Value: mediaID}}
	result, err := collection.FindOne(ctx, mediaQuery)
	if err != nil {
		return media, err
	}
	_, err = result.One(ctx, &media)
	return media, err
}

package mediamicroservices_nosql

import (
	"context"
	"errors"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Review struct {
	ReviewID  int64 `bson:"_id"`
	ReqID     int64
	Timestamp int64
	UserID    int64
	MovieID   string
	Text      string
	Rating    int
}

type ReviewStorageService interface {
	StoreReview(ctx context.Context, reqID int64, review Review) (Review, error)
	ReadReviews(ctx context.Context, reqID int64, reviewIDs []int64) ([]Review, error)
}

type ReviewStorageServiceImpl struct {
	database backend.NoSQLDatabase
	cache    backend.Cache
}

func NewReviewStorageServiceImpl(ctx context.Context, database backend.NoSQLDatabase, cache backend.Cache) (ReviewStorageService, error) {
	s := &ReviewStorageServiceImpl{database: database, cache: cache}
	return s, nil
}

func (s *ReviewStorageServiceImpl) StoreReview(ctx context.Context, reqID int64, review Review) (Review, error) {
	collection, err := s.database.GetCollection(ctx, "review_storage_db", "review")
	if err != nil {
		return Review{}, err
	}
	err = collection.InsertOne(ctx, review)
	if err != nil {
		return Review{}, err
	}

	return review, err
}

func (s *ReviewStorageServiceImpl) ReadReviews(ctx context.Context, reqID int64, reviewIDs []int64) ([]Review, error) {
	if len(reviewIDs) == 0 {
		return nil, nil
	}

	reviewIDsNotCached := make(map[int64]bool)
	for _, pid := range reviewIDs {
		reviewIDsNotCached[pid] = true
	}
	if len(reviewIDsNotCached) != len(reviewIDs) {
		return nil, errors.New("review_ids are duplicated")
	}

	var keys []string
	for _, rid := range reviewIDs {
		keys = append(keys, strconv.FormatInt(rid, 10))
	}

	cachedReviews := make([]Review, len(keys))
	var cachedValues []interface{}
	for idx := range cachedReviews {
		cachedValues = append(cachedValues, &cachedReviews[idx])
	}

	err := s.cache.Mget(ctx, keys, cachedValues)
	if err != nil {
		return nil, err
	}

	var returnReviews []Review
	for _, review := range cachedReviews {
		if review.ReviewID != 0 {
			delete(reviewIDsNotCached, review.ReviewID)
			returnReviews = append(returnReviews, review)
		}
	}

	if len(reviewIDsNotCached) != 0 {
		var reviews []Review
		var unique_pids []int64
		for k := range reviewIDsNotCached {
			unique_pids = append(unique_pids, k)
		}
		collection, err := s.database.GetCollection(ctx, "review_storage_db", "review")
		if err != nil {
			return []Review{}, err
		}

		query := bson.D{
			{Key: "PostID", Value: bson.D{
				{Key: "$in", Value: unique_pids},
			}},
		}

		cursor, err := collection.FindMany(ctx, query)
		if err != nil {
			return []Review{}, err
		}
		err = cursor.All(ctx, &reviews)
		if err != nil {
			return []Review{}, err
		}
		returnReviews = append(returnReviews, reviews...)
		for _, review := range reviews {
			s.cache.Put(ctx, strconv.FormatInt(review.ReviewID, 10), review)
		}
	}
	return returnReviews, nil
}

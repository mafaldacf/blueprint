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

	unique_review_ids := make(map[int64]bool)
	for _, pid := range reviewIDs {
		unique_review_ids[pid] = true
	}
	if len(unique_review_ids) != len(reviewIDs) {
		return nil, errors.New("review_ids are duplicated")
	}

	var keys []string
	for _, rid := range reviewIDs {
		keys = append(keys, strconv.FormatInt(rid, 10))
	}
	values := make([]Review, len(keys))
	var retvals []interface{}
	for idx := range values {
		retvals = append(retvals, &values[idx])
	}

	s.cache.Mget(ctx, keys, retvals)
	var retreviews []Review
	for _, review := range values {
		if review.ReviewID != 0 {
			delete(unique_review_ids, review.ReviewID)
			retreviews = append(retreviews, review)
		}
	}

	if len(unique_review_ids) != 0 {
		var review []Review
		var unique_pids []int64
		for k := range unique_review_ids {
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

		vals, err := collection.FindMany(ctx, query)
		if err != nil {
			return []Review{}, err
		}
		err = vals.All(ctx, &review)
		if err != nil {
			return []Review{}, err
		}
		retreviews = append(retreviews, review...)
		for _, review := range review {
			s.cache.Put(ctx, strconv.FormatInt(review.ReviewID, 10), review)
		}
	}
	return retreviews, nil
}

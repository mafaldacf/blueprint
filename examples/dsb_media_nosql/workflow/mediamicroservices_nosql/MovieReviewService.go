package mediamicroservices_nosql

import (
	"context"
	"errors"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type MovieReviewData struct {
	ReviewID  int64
	Timestamp int64
}

type MovieReview struct {
	MovieID string `bson:"_id"`
	Reviews []MovieReviewData
}

type MovieReviewService interface {
	UploadMovieReview(ctx context.Context, reqID int64, movieID string, reviewID int64, timestamp int64) error
	ReadMovieReviews(ctx context.Context, reqID int64, movieID string, start int, stop int) ([]Review, error)
}

type MovieReviewServiceImpl struct {
	database             backend.NoSQLDatabase
	cache                backend.Cache
	reviewStorageService ReviewStorageService
}

func NewMovieReviewServiceImpl(ctx context.Context, database backend.NoSQLDatabase, cache backend.Cache, reviewStorageService ReviewStorageService) (MovieReviewService, error) {
	s := &MovieReviewServiceImpl{database: database, cache: cache, reviewStorageService: reviewStorageService}
	return s, nil
}

func (s *MovieReviewServiceImpl) UploadMovieReview(ctx context.Context, reqID int64, movieID string, reviewID int64, timestamp int64) error {
	collection, err := s.database.GetCollection(ctx, "movie_review_db", "movie_review")
	if err != nil {
		return err
	}

	var movieReview MovieReview
	query := bson.D{{Key: "MovieID", Value: movieID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return err
	}
	found, err := result.One(ctx, &movieReview)
	if err != nil {
		return err
	}

	if !found {
		movieReview = MovieReview{
			MovieID: movieID,
			Reviews: []MovieReviewData{{
				ReviewID:  reviewID,
				Timestamp: timestamp,
			}},
		}
		err := collection.InsertOne(ctx, movieReview)
		if err != nil {
			return err
		}
	} else {
		update := bson.D{
			{Key: "$push", Value: bson.D{
				{Key: "Reviews", Value: bson.D{
					{Key: "$each", Value: bson.A{
						bson.D{
							{Key: "ReviewID", Value: reviewID},
							{Key: "Timestamp", Value: timestamp},
						},
					}},
					{Key: "$position", Value: 0},
				}},
			}},
		}
		_, err = collection.UpdateMany(ctx, query, update)
		if err != nil {
			return errors.New("Failed to insert user timeline user to Database")
		}
	}

	var reviews []MovieReviewData
	// Ignore error check for Get!
	_, err = s.cache.Get(ctx, movieID, &reviews)
	if err != nil {
		return err
	}
	reviews = append(reviews, MovieReviewData{ReviewID: reviewID, Timestamp: timestamp})
	return s.cache.Put(ctx, movieID, reviews)
}

func (s *MovieReviewServiceImpl) ReadMovieReviews(ctx context.Context, reqID int64, movieID string, start int, stop int) ([]Review, error) {
	if stop <= start || start < 0 {
		return nil, nil
	}

	var reviews []MovieReviewData
	exists, err := s.cache.Get(ctx, movieID, &reviews)
	if err != nil {
		return nil, err
	}

	var reviewIds []int64
	seen_reviews_ids := make(map[int64]bool)
	for _, review := range reviews {
		reviewIds = append(reviewIds, review.ReviewID)
		seen_reviews_ids[review.ReviewID] = true
	}
	db_start := start + len(reviewIds)
	var new_reviews_ids []int64
	if db_start < stop {
		collection, err := s.database.GetCollection(ctx, "movie_review_db", "movie_review")
		if err != nil {
			return nil, err
		}

		query := bson.D{{Key: "MovieID", Value: movieID}}

		projection := bson.D{
			{Key: "posts", Value: bson.D{
				{Key: "$slice", Value: bson.A{0, stop}},
			}},
		}

		cursor, err := collection.FindOne(ctx, query, projection)
		if err != nil {
			return nil, err
		}
		var movieReview MovieReview
		exists, err = cursor.One(ctx, &movieReview)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, errors.New("failed to find movie reviews in database")
		}
		for _, review := range movieReview.Reviews {
			// Avoid duplicated reviews
			if _, ok := seen_reviews_ids[review.ReviewID]; ok {
				continue
			}
			new_reviews_ids = append(new_reviews_ids, review.ReviewID)
		}
	}

	reviewIds = append(new_reviews_ids, reviewIds...)

	// update reviews []Review
	ret_reviews, err := s.reviewStorageService.ReadReviews(ctx, reqID, reviewIds)
	if err != nil {
		return nil, err
	}

	if len(new_reviews_ids) > 0 {
		err = s.cache.Put(ctx, movieID, reviews)
		if err != nil {
			return nil, err
		}
	}

	return ret_reviews, nil
}

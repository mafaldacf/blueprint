package mediamicroservices_nosql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type RatingService interface {
	UploadRating(ctx context.Context, reqID int64, movieID string, rating int) error
}

type RatingServiceImpl struct {
	cache                backend.Cache
	composeReviewService ComposeReviewService
}

func NewRatingServiceImpl(ctx context.Context, cache backend.Cache, composeReviewService ComposeReviewService) (RatingService, error) {
	s := &RatingServiceImpl{cache: cache, composeReviewService: composeReviewService}
	return s, nil
}

func (s *RatingServiceImpl) UploadRating(ctx context.Context, reqID int64, movieID string, rating int) error {
	err := s.composeReviewService.UploadRating(ctx, reqID, rating)
	if err != nil {
		return err
	}

	// TODO: should be IncrBy(rating)
	// e.g.,
	// redis_client->incrby(movie_id + ":uncommit_sum", rating);
	// redis_client->incr(movie_id + ":uncommit_num");
	// redis_client->sync_commit();
	_, err = s.cache.Incr(ctx, movieID)
	return err
}

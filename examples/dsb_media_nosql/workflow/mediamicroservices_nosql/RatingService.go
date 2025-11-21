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

	var cachedUncommitSum int
	_, err = s.cache.Get(ctx, movieID+":uncommit_sum", &cachedUncommitSum)
	if err != nil {
		return err
	}
	err = s.cache.Put(ctx, movieID+":uncommit_sum", cachedUncommitSum+rating)
	if err != nil {
		return err
	}

	_, err = s.cache.Incr(ctx, movieID+":uncommit_num")
	if err != nil {
		return err
	}

	return err
}

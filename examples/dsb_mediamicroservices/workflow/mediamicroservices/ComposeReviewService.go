package mediamicroservices

import (
	"context"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

const NUM_COMPONENTS = 5

type ComposeReview struct {
	MovieID string `bson:"_id"`
	Title   string
	CastID  string
	PlotID  string
}

type ComposeReviewService interface {
	UploadText(ctx context.Context, reqID int64, text string) error
	UploadRating(ctx context.Context, reqID int64, rating int) error
	UploadUniqueId(ctx context.Context, reqID int64, reviewID int64) error
	UploadMovieId(ctx context.Context, reqID int64, movieID string) error
	UploadUserId(ctx context.Context, reqID int64, userID int64) error
}

type ComposeReviewServiceImpl struct {
	cache                backend.Cache
	reviewStorageService ReviewStorageService
	userReviewService    UserReviewService
	movieReviewService   MovieReviewService
}

func NewComposeReviewServiceImpl(ctx context.Context, cache backend.Cache, reviewStorageService ReviewStorageService, userReviewService UserReviewService, movieReviewService MovieReviewService) (ComposeReviewService, error) {
	s := &ComposeReviewServiceImpl{cache: cache, reviewStorageService: reviewStorageService, userReviewService: userReviewService, movieReviewService: movieReviewService}
	return s, nil
}

func (s *ComposeReviewServiceImpl) _ComposeAndUpload(ctx context.Context, reqID int64) error {
	keyUniqueID := strconv.FormatInt(reqID, 10) + ":review_id"
	keyMovieID := strconv.FormatInt(reqID, 10) + ":movie_id"
	keyUserID := strconv.FormatInt(reqID, 10) + ":user_id"
	keyText := strconv.FormatInt(reqID, 10) + ":text"
	keyRating := strconv.FormatInt(reqID, 10) + ":rating"

	var newReview Review
	var keys = []string{keyUniqueID, keyMovieID, keyUserID, keyText, keyRating}
	var vals []interface{}
	err := s.cache.Mget(ctx, keys, vals)
	if err != nil {
		return err
	}

	for i, val := range vals {
		key := keys[i]
		switch key {
		case keyUniqueID:
			newReview.ReviewID = val.(int64)
		case keyMovieID:
			newReview.MovieID = val.(string)
		case keyUserID:
			newReview.UserID = val.(int64)
		case keyText:
			newReview.Text = val.(string)
		case keyRating:
			newReview.Rating = val.(int)
		}
	}

	_, err = s.reviewStorageService.StoreReview(ctx, reqID, newReview)
	if err != nil {
		return err
	}

	err = s.userReviewService.UploadUserReview(ctx, reqID, newReview.UserID, newReview.ReviewID, newReview.Timestamp)
	if err != nil {
		return err
	}

	err = s.movieReviewService.UploadMovieReview(ctx, reqID, newReview.MovieID, newReview.ReviewID, newReview.Timestamp)
	if err != nil {
		return err
	}

	return nil
}

func (s *ComposeReviewServiceImpl) UploadText(ctx context.Context, reqID int64, text string) error {
	keyCounter := strconv.FormatInt(reqID, 10) + ":counter"

	// Store text to memcached
	var counterValue int64
	keyUserID := strconv.FormatInt(reqID, 10) + ":text"
	err := s.cache.Put(ctx, keyUserID, text)
	if err != nil {
		return err
	}

	counterValue, err = s.cache.Incr(ctx, keyCounter)
	if err != nil {
		return err
	}

	if counterValue == NUM_COMPONENTS {
		s._ComposeAndUpload(ctx, reqID)
	}
	return nil
}

func (s *ComposeReviewServiceImpl) UploadRating(ctx context.Context, reqID int64, rating int) error {
	keyCounter := strconv.FormatInt(reqID, 10) + ":counter"

	// Store rating to memcached
	var counterValue int64
	keyUserID := strconv.FormatInt(reqID, 10) + ":rating"
	ratingStr := strconv.Itoa(rating)
	err := s.cache.Put(ctx, keyUserID, ratingStr)
	if err != nil {
		return err
	}

	counterValue, err = s.cache.Incr(ctx, keyCounter)
	if err != nil {
		return err
	}

	if counterValue == NUM_COMPONENTS {
		s._ComposeAndUpload(ctx, reqID)
	}
	return nil
}

func (s *ComposeReviewServiceImpl) UploadUniqueId(ctx context.Context, reqID int64, reviewID int64) error {
	keyCounter := strconv.FormatInt(reqID, 10) + ":counter"

	// Store review_id to memcached
	var counterValue int64
	keyUserID := strconv.FormatInt(reqID, 10) + ":review_id"
	err := s.cache.Put(ctx, keyUserID, reviewID)
	if err != nil {
		return err
	}

	counterValue, err = s.cache.Incr(ctx, keyCounter)
	if err != nil {
		return err
	}

	if counterValue == NUM_COMPONENTS {
		s._ComposeAndUpload(ctx, reqID)
	}
	return nil
}

func (s *ComposeReviewServiceImpl) UploadMovieId(ctx context.Context, reqID int64, movieID string) error {
	keyCounter := strconv.FormatInt(reqID, 10) + ":counter"

	// Store movie_id to memcached
	var counterValue int64
	keyUserID := strconv.FormatInt(reqID, 10) + ":movie_id"
	err := s.cache.Put(ctx, keyUserID, movieID)
	if err != nil {
		return err
	}

	counterValue, err = s.cache.Incr(ctx, keyCounter)
	if err != nil {
		return err
	}

	if counterValue == NUM_COMPONENTS {
		s._ComposeAndUpload(ctx, reqID)
	}
	return nil
}

func (s *ComposeReviewServiceImpl) UploadUserId(ctx context.Context, reqID int64, userID int64) error {
	keyCounter := strconv.FormatInt(reqID, 10) + ":counter"

	// Store user_id to memcached
	var counterValue int64
	keyUserID := strconv.FormatInt(reqID, 10) + ":user_id"
	userIDStr := strconv.FormatInt(userID, 10)
	err := s.cache.Put(ctx, keyUserID, userIDStr)
	if err != nil {
		return err
	}

	counterValue, err = s.cache.Incr(ctx, keyCounter)
	if err != nil {
		return err
	}

	if counterValue == NUM_COMPONENTS {
		s._ComposeAndUpload(ctx, reqID)
	}
	return nil
}

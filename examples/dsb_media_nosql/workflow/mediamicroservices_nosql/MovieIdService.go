package mediamicroservices_nosql

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type MovieId struct {
	MovieID string `bson:"_id"`
	Title   string
}

type MovieIdService interface {
	RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error)
	UploadMovieId(ctx context.Context, reqID int64, title string, rating int) error
	ReadMovieId(ctx context.Context, reqID int64, title string) (MovieId, error)
}

type MovieIdServiceImpl struct {
	database             backend.NoSQLDatabase
	cache                backend.Cache
	composeReviewService ComposeReviewService
	ratingService        RatingService
}

func NewMovieIdServiceImpl(ctx context.Context, database backend.NoSQLDatabase, cache backend.Cache, composeReviewService ComposeReviewService, ratingService RatingService) (MovieIdService, error) {
	s := &MovieIdServiceImpl{database: database, cache: cache, composeReviewService: composeReviewService, ratingService: ratingService}
	return s, nil
}

func (s *MovieIdServiceImpl) RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error) {
	movie := MovieId{
		MovieID: movieID,
		Title:   title,
	}

	collection, err := s.database.GetCollection(ctx, "movie_id_db", "movie")
	if err != nil {
		return MovieId{}, err
	}
	err = collection.InsertOne(ctx, movie)
	if err != nil {
		return MovieId{}, err
	}

	return movie, err
}

func (s *MovieIdServiceImpl) UploadMovieId(ctx context.Context, reqID int64, title string, rating int) error {
	return nil
	var movieID string

	ok, err := s.cache.Get(ctx, title, &movieID)
	if err != nil {
		return err
	}
	if !ok {
		// if not cached in memcached
		var movie MovieId
		collection, err := s.database.GetCollection(ctx, "movie_id_db", "movie")
		if err != nil {
			return err
		}
		query := bson.D{{Key: "Title", Value: title}}
		result, err := collection.FindOne(ctx, query)
		if err != nil {
			return err
		}
		found, err := result.One(ctx, &movie)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("movie %s not found in MongoDB", title)
		}
		movieID = movie.MovieID
	}

	err = s.cache.Put(ctx, title, movieID)
	if err != nil {
		return err
	}

	err = s.composeReviewService.UploadMovieId(ctx, reqID, movieID)
	if err != nil {
		return err
	}

	err = s.ratingService.UploadRating(ctx, reqID, movieID, rating)
	if err != nil {
		return err
	}

	return nil
}

func (s *MovieIdServiceImpl) ReadMovieId(ctx context.Context, reqID int64, title string) (MovieId, error) {
	var movieID string
	var movie MovieId

	ok, err := s.cache.Get(ctx, title, &movieID)
	if err != nil {
		return MovieId{}, err
	}
	if ok {
		movie = MovieId{MovieID: movieID, Title: title}
	} else {
		// if not cached in memcached
		collection, err := s.database.GetCollection(ctx, "movie_id_db", "movie")
		if err != nil {
			return MovieId{}, err
		}
		query := bson.D{{Key: "Title", Value: title}}
		result, err := collection.FindOne(ctx, query)
		if err != nil {
			return MovieId{}, err
		}
		found, err := result.One(ctx, &movie)
		if err != nil {
			return MovieId{}, err
		}
		if !found {
			return MovieId{}, fmt.Errorf("movie %s not found in MongoDB", title)
		}
		movieID = movie.MovieID
	}

	err = s.cache.Put(ctx, title, movieID)
	if err != nil {
		return MovieId{}, err
	}

	return movie, nil
}

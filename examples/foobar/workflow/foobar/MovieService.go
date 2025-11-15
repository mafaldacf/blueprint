package foobar

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Movie struct {
	MovieID string
	PlotID string
}

type MovieService interface {
	ReadMovie(ctx context.Context, movieID string) (Movie, error)
}

type MovieServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewMovieServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (MovieService, error) {
	d := &MovieServiceImpl{db: db}
	return d, nil
}

func (s *MovieServiceImpl) ReadMovie(ctx context.Context, movieID string) (Movie, error) {
	var movie Movie

	collection, err := s.db.GetCollection(ctx, "movie_db", "movie")
	if err != nil {
		return Movie{}, err
	}

	query := bson.D{{Key: "MovieID", Value: movieID}}
	cursor, err := collection.FindOne(ctx, query)
	if err != nil {
		return Movie{}, err
	}

	res, err := cursor.One(ctx, &movie)
	if !res || err != nil {
		return Movie{}, err
	}

	return movie, nil
}

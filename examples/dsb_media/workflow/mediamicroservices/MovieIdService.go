package mediamicroservices

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type MovieId struct {
	MovieID string `bson:"_id"`
	Title   string `bson:"title"`
}

type MovieIdService interface {
	RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error)
}

type MovieIdServiceImpl struct {
	movieIdDB backend.NoSQLDatabase
}

func NewMovieIdServiceImpl(ctx context.Context, movieIdDB backend.NoSQLDatabase) (MovieIdService, error) {
	return &MovieIdServiceImpl{movieIdDB: movieIdDB}, nil
}

func (m *MovieIdServiceImpl) RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error) {
	collection, err := m.movieIdDB.GetCollection(ctx, "movie-id", "movie-id")
	if err != nil {
		return MovieId{}, err
	}
	movieId := MovieId{
		MovieID: movieID,
		Title:   title,
	}
	return movieId, collection.InsertOne(ctx, movieId)
}

package mediamicroservices

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type MovieInfo struct {
	MovieID string `bson:"_id"`
	Title   string `bson:"title"`
	Casts   string `bson:"casts"`
}

type MovieInfoService interface {
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error)
}

type MovieInfoServiceImpl struct {
	movieInfoDB backend.NoSQLDatabase
}

func NewMovieInfoServiceImpl(ctx context.Context, movieIdDB backend.NoSQLDatabase) (MovieInfoService, error) {
	return &MovieInfoServiceImpl{movieInfoDB: movieIdDB}, nil
}

func (m *MovieInfoServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error) {
	collection, err := m.movieInfoDB.GetCollection(ctx, "movie-info", "movie-info")
	if err != nil {
		return MovieInfo{}, err
	}
	movieInfo := MovieInfo{
		MovieID: movieID,
		Title:   title,
		Casts:   casts,
	}
	return movieInfo, collection.InsertOne(ctx, movieInfo)
}

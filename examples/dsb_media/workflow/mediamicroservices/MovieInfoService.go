package mediamicroservices

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type MovieInfo struct {
	MovieID string `bson:"_id"`
	Title   string `bson:"title"`
	Casts   string `bson:"casts"`
}

type MovieInfoService interface {
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error)
	ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error)
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

func (m *MovieInfoServiceImpl) ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error) {
	var movieInfo MovieInfo

	collection, err := m.movieInfoDB.GetCollection(ctx, "movie-info", "movie-info")
	if err != nil {
		return movieInfo, err
	}

	query := bson.D{{Key: "movieid", Value: movieID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return movieInfo, err
	}

	res, err := result.One(ctx, &movieInfo)
	if !res || err != nil {
		return movieInfo, err
	}

	return movieInfo, err
}

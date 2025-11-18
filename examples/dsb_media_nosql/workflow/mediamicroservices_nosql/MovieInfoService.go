package mediamicroservices_nosql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Cast struct {
	CastID     string
	CastInfoID string
	Character  string
}

type MovieInfo struct {
	MovieID string `bson:"_id"`
	Title   string
	Casts   []Cast
	PlotID  string
}

type MovieInfoService interface {
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts []Cast, plotID string) (MovieInfo, error)
	ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error)
}

type MovieInfoServiceImpl struct {
	database backend.NoSQLDatabase
}

func NewMovieInfoServiceImpl(ctx context.Context, database backend.NoSQLDatabase) (MovieInfoService, error) {
	s := &MovieInfoServiceImpl{database: database}
	return s, nil
}

func (s *MovieInfoServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts []Cast, plotID string) (MovieInfo, error) {
	movieInfo := MovieInfo{
		MovieID: movieID,
		Title:   title,
		Casts:   casts,
		PlotID:  plotID,
	}
	collection, err := s.database.GetCollection(ctx, "movie_info_db", "movie_info")
	if err != nil {
		return MovieInfo{}, err
	}
	err = collection.InsertOne(ctx, movieInfo)
	if err != nil {
		return MovieInfo{}, err
	}

	return movieInfo, err
}

func (s *MovieInfoServiceImpl) ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error) {
	var movieInfo MovieInfo
	collection, err := s.database.GetCollection(ctx, "movie_info_db", "movie_info")
	if err != nil {
		return movieInfo, err
	}
	query := bson.D{{Key: "MovieID", Value: movieID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return movieInfo, err
	}
	res, err := result.One(ctx, &movieInfo)
	if !res || err != nil {
		return MovieInfo{}, err
	}
	return movieInfo, err
}

package mediamicroservices_nosql

import (
	"context"
	"fmt"

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
	PlotID  int64
}

type Rating struct {
	AvgRating float64
	NumRating int
}

type MovieInfoService interface {
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts []Cast, plotID int64, thumbnailIDs []string, photoIDs []string, videoIDs []string, avgRating string, numRating int) (MovieInfo, error)
	ReadMovieInfo(ctx context.Context, reqID int64, movieID string) (MovieInfo, error)
	UpdateRating(ctx context.Context, reqID int64, movieID string, sumUncommittedRating int, numUncommittedRating int) error
}

type MovieInfoServiceImpl struct {
	database      backend.NoSQLDatabase
	cache         backend.Cache
	socialGraphDB backend.NoSQLDatabase
}

func NewMovieInfoServiceImpl(ctx context.Context, database backend.NoSQLDatabase, cache backend.Cache, socialGraphDB backend.NoSQLDatabase) (MovieInfoService, error) {
	s := &MovieInfoServiceImpl{database: database, cache: cache, socialGraphDB: socialGraphDB}
	return s, nil
}

func (s *MovieInfoServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts []Cast, plotID int64, thumbnailIDs []string, photoIDs []string, videoIDs []string, avgRating string, numRating int) (MovieInfo, error) {
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

	var cachedMovieInfo interface{}
	found, err := s.cache.Get(ctx, movieID, &cachedMovieInfo)
	if err == nil {
		return movieInfo, nil
	}
	if found {
		movieInfo = cachedMovieInfo.(MovieInfo)
		return movieInfo, nil
	}

	// if not cached in memcached
	collection, err := s.database.GetCollection(ctx, "movie_info_db", "movie_info")
	if err != nil {
		return movieInfo, err
	}
	query := bson.D{{Key: "MovieID", Value: movieID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return movieInfo, err
	}
	found, err = result.One(ctx, &movieInfo)
	if err != nil {
		return MovieInfo{}, err
	}
	if !found {
		return MovieInfo{}, fmt.Errorf("movie %s not found in MongoDB", movieID)
	}

	// upload movie-info to memcached
	err = s.cache.Put(ctx, movieID, movieInfo)
	if err != nil {
		return movieInfo, err
	}

	return movieInfo, err
}

func (s *MovieInfoServiceImpl) UpdateRating(ctx context.Context, reqID int64, movieID string, sumUncommittedRating int, numUncommittedRating int) error {
	collection, err := s.socialGraphDB.GetCollection(ctx, "social_graph_db", "social_graph")
	if err != nil {
		return err
	}
	var rating Rating
	query := bson.D{{Key: "MovieID", Value: movieID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return err
	}
	found, err := result.One(ctx, &rating)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("movie %s not found", movieID)
	}
	avgRating := rating.AvgRating
	numRating := rating.NumRating

	avgRating = (avgRating*float64(numRating) + float64(sumUncommittedRating)) / float64(numRating+numUncommittedRating)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "AvgRating", Value: avgRating}, {Key: "NumRating", Value: numRating}}}}
	ok, err := collection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}
	if ok == 0 {
		return fmt.Errorf("movie %s not found", movieID)
	}
	return nil
}

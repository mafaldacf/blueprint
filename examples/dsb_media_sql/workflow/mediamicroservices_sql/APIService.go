package mediamicroservices_sql

import (
	"context"
)

type APIService interface {
	//RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error)
	//WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error)
	RegisterMovie(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieId, MovieInfo, error)
	ReadMovie(ctx context.Context, reqID int64, movieID string) (MovieId, MovieInfo, error)
}

type APIServiceImpl struct {
	movieIdService   MovieIdService
	movieInfoService MovieInfoService
}

func NewAPIServiceImpl(ctx context.Context, movieIdService MovieIdService, movieInfoService MovieInfoService) (APIService, error) {
	return &APIServiceImpl{movieIdService: movieIdService, movieInfoService: movieInfoService}, nil
}

/* func (api *APIServiceImpl) RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error) {
	return api.movieIdService.RegisterMovieId(ctx, reqID, movieID, title)
} */

/* func (api *APIServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error) {
	return api.movieInfoService.WriteMovieInfo(ctx, reqID, movieID, title, casts)
} */

func (api *APIServiceImpl) RegisterMovie(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieId, MovieInfo, error) {
	movieId, err := api.movieIdService.RegisterMovieId(ctx, reqID, movieID, title)
	if err != nil {
		return MovieId{}, MovieInfo{}, err
	}

	movieInfo, err := api.movieInfoService.WriteMovieInfo(ctx, reqID, movieID, title, casts)
	if err != nil {
		return MovieId{}, MovieInfo{}, err
	}

	return movieId, movieInfo, nil
}

func (api *APIServiceImpl) ReadMovie(ctx context.Context, reqID int64, movieId string) (MovieId, MovieInfo, error) {
	movie1, err1 := api.movieIdService.ReadMovieId(ctx, reqID, movieId)
	movie2, err2 := api.movieInfoService.ReadMovieInfo(ctx, reqID, movieId)
	if err1 != nil {
		return MovieId{}, MovieInfo{}, err1
	}
	if err2 != nil {
		return MovieId{}, MovieInfo{}, err1
	}
	return movie1, movie2, nil
}

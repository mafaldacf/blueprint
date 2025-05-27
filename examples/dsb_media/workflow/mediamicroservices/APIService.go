package mediamicroservices

import (
	"context"
)

type APIService interface {
	RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error)
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error)
}

type APIServiceImpl struct {
	movieIdService   MovieIdService
	movieInfoService MovieInfoService
}

func NewAPIServiceImpl(ctx context.Context, movieIdService MovieIdService, movieInfoService MovieInfoService) (APIService, error) {
	return &APIServiceImpl{movieIdService: movieIdService, movieInfoService: movieInfoService}, nil
}

func (api *APIServiceImpl) RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error) {
	return api.movieIdService.RegisterMovieId(ctx, reqID, movieID, title)
}

func (api *APIServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error) {
	return api.movieInfoService.WriteMovieInfo(ctx, reqID, movieID, title, casts)
}

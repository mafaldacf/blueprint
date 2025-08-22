package mediamicroservices_sql

import (
	"context"
)

type APIService interface {
	//RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error)
	//WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error)
	RegisterMovie(ctx context.Context, reqID int64, movieID string, title string, castID string, castName string, castGender string, castIntro string, plotID string, plotText string) (MovieId, MovieInfo, CastInfo, Plot, error)
	//ReadMovie(ctx context.Context, reqID int64, movieID string) (MovieId, MovieInfo, error)
	ReadPage(ctx context.Context, reqID int64, title string) (MovieId, MovieInfo, CastInfo, Plot, error)
}

type APIServiceImpl struct {
	movieIdService   MovieIdService
	movieInfoService MovieInfoService
	castInfoService  CastInfoService
	plotService      PlotService
}

func NewAPIServiceImpl(ctx context.Context, movieIdService MovieIdService, movieInfoService MovieInfoService, castInfoService CastInfoService, plotService PlotService) (APIService, error) {
	return &APIServiceImpl{movieIdService: movieIdService, movieInfoService: movieInfoService, castInfoService: castInfoService, plotService: plotService}, nil
}

/* func (api *APIServiceImpl) RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error) {
	return api.movieIdService.RegisterMovieId(ctx, reqID, movieID, title)
} */

/* func (api *APIServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts string) (MovieInfo, error) {
	return api.movieInfoService.WriteMovieInfo(ctx, reqID, movieID, title, casts)
} */

func (api *APIServiceImpl) RegisterMovie(ctx context.Context, reqID int64, movieID string, title string, castID string, castName string, castGender string, castIntro string, plotID string, plotText string) (MovieId, MovieInfo, CastInfo, Plot, error) {
	movie, _, err := api.movieIdService.RegisterMovieId(ctx, reqID, movieID, title)
	if err != nil {
		return MovieId{}, MovieInfo{}, CastInfo{}, Plot{}, err
	}

	movieInfo, err := api.movieInfoService.WriteMovieInfo(ctx, reqID, movie.movieid, title, castID, plotID)
	if err != nil {
		return MovieId{}, MovieInfo{}, CastInfo{}, Plot{}, err
	}

	/* return movie, movieInfo, CastInfo{}, Plot{}, nil */

	castInfo, err := api.castInfoService.WriteCastInfo(ctx, reqID, castID, castName, castGender, castIntro)
	if err != nil {
		return MovieId{}, MovieInfo{}, CastInfo{}, Plot{}, err
	}

	plot, err := api.plotService.WritePlot(ctx, reqID, plotID, plotText)
	if err != nil {
		return MovieId{}, MovieInfo{}, CastInfo{}, Plot{}, err
	}

	return movie, movieInfo, castInfo, plot, nil
}

/* func (api *APIServiceImpl) ReadMovie(ctx context.Context, reqID int64, movieId string) (MovieId, MovieInfo, error) {
	movie1, err1 := api.movieIdService.ReadMovieId(ctx, reqID, movieId)
	movie2, err2 := api.movieInfoService.ReadMovieInfo(ctx, reqID, movieId)
	if err1 != nil {
		return MovieId{}, MovieInfo{}, err1
	}
	if err2 != nil {
		return MovieId{}, MovieInfo{}, err1
	}
	return movie1, movie2, nil
} */

func (api *APIServiceImpl) ReadPage(ctx context.Context, reqID int64, movieID string) (MovieId, MovieInfo, CastInfo, Plot, error) {
	movie, err1 := api.movieIdService.ReadMovieId(ctx, reqID, movieID)
	if err1 != nil {
		return MovieId{}, MovieInfo{}, CastInfo{}, Plot{}, err1
	}
	movieInfo, err2 := api.movieInfoService.ReadMovieInfo(ctx, reqID, movie.movieid)
	if err2 != nil {
		return MovieId{}, MovieInfo{}, CastInfo{}, Plot{}, err2
	}

	/* return movie, movieInfo, CastInfo{}, Plot{}, nil */

	castInfo, err3 := api.castInfoService.ReadCastInfo(ctx, reqID, movieInfo.plotid)
	if err3 != nil {
		return MovieId{}, MovieInfo{}, CastInfo{}, Plot{}, err3
	}
	plot, err4 := api.plotService.ReadPlot(ctx, reqID, movieInfo.plotid)
	if err4 != nil {
		return MovieId{}, MovieInfo{}, CastInfo{}, Plot{}, err4
	}
	return movie, movieInfo, castInfo, plot, nil
}

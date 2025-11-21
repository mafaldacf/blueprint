package mediamicroservices_nosql

import (
	"context"
)

type APIService interface {
	RegisterMovie(ctx context.Context, reqID int64, movieID string, title string, castRequest []CastRequest, plotID int64, plotText string, thumbnailIDs []string, photoIDs []string, videoIDs []string, avgRating string, numRating int) (MovieId, MovieInfo, Plot, error)
	ReadPage(ctx context.Context, reqID int64, title string, reviewStart int, reviewStop int) (MovieId, MovieInfo, []Review, []CastInfo, Plot, error)

	RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error)
	WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts []Cast, plotID int64, thumbnailIDs []string, photoIDs []string, videoIDs []string, avgRating string, numRating int) (MovieInfo, error)
	WriteCastInfo(ctx context.Context, reqID int64, castInfoID string, name string, gender string, intro string) (CastInfo, error)
	WritePlot(ctx context.Context, reqID int64, plotID int64, plotText string) (Plot, error)

	RegisterUser(ctx context.Context, reqID string, firstName string, lastName string, username string, password string) (User, error)
	Login(ctx context.Context, reqID int64, username string, password string) error

	UploadUserWithUsername(ctx context.Context, reqID int64, username string) error
	UploadText(ctx context.Context, reqID int64, text string) error
	UploadMovieId(ctx context.Context, reqID int64, title string, rating int) error
	UploadUniqueId(ctx context.Context, reqID int64) error
}

type APIServiceImpl struct {
	userService     UserService
	textService     TextService
	movieIdService  MovieIdService
	uniqueIdService UniqueIdService

	movieInfoService MovieInfoService
	castInfoService  CastInfoService
	plotService      PlotService
	pageService      PageService
}

func NewAPIServiceImpl(ctx context.Context, userService UserService, textService TextService, movieIdService MovieIdService, uniqueIdService UniqueIdService, movieInfoService MovieInfoService, castInfoService CastInfoService, plotService PlotService, pageService PageService) (APIService, error) {
	return &APIServiceImpl{userService: userService, textService: textService, movieIdService: movieIdService, uniqueIdService: uniqueIdService, movieInfoService: movieInfoService, castInfoService: castInfoService, plotService: plotService, pageService: pageService}, nil
}

type CastRequest struct {
	CastID     string
	CastInfoID string
	Character  string

	Name   string
	Gender string
	Intro  string
}

func (api *APIServiceImpl) RegisterMovie(ctx context.Context, reqID int64, movieID string, title string, castRequest []CastRequest, plotID int64, plotText string, thumbnailIDs []string, photoIDs []string, videoIDs []string, avgRating string, numRating int) (MovieId, MovieInfo, Plot, error) {
	movie, err := api.movieIdService.RegisterMovieId(ctx, reqID, movieID, title)
	if err != nil {
		return MovieId{}, MovieInfo{}, Plot{}, err
	}

	var casts []Cast
	for _, cast := range castRequest {
		casts = append(casts, Cast{
			CastID:     cast.CastID,
			CastInfoID: cast.CastInfoID,
			Character:  cast.Character,
		})
	}

	movieInfo, err := api.movieInfoService.WriteMovieInfo(ctx, reqID, movie.MovieID, title, casts, plotID, thumbnailIDs, photoIDs, videoIDs, avgRating, numRating)
	if err != nil {
		return MovieId{}, MovieInfo{}, Plot{}, err
	}

	for _, cast := range castRequest {
		_, err := api.castInfoService.WriteCastInfo(ctx, reqID, cast.CastInfoID, cast.Name, cast.Gender, cast.Intro)
		if err != nil {
			return MovieId{}, MovieInfo{}, Plot{}, err
		}
	}

	plot, err := api.plotService.WritePlot(ctx, reqID, plotID, plotText)
	if err != nil {
		return MovieId{}, MovieInfo{}, Plot{}, err
	}

	return movie, movieInfo, plot, nil
}

func (api *APIServiceImpl) ReadPage(ctx context.Context, reqID int64, title string, reviewStart int, reviewStop int) (MovieId, MovieInfo, []Review, []CastInfo, Plot, error) {
	movie, err := api.movieIdService.ReadMovieId(ctx, reqID, title)
	if err != nil {
		return MovieId{}, MovieInfo{}, nil, nil, Plot{}, err
	}

	movieInfo, reviews, castInfos, plot, err := api.pageService.ReadPage(ctx, reqID, movie.MovieID, reviewStart, reviewStop)
	if err != nil {
		return MovieId{}, MovieInfo{}, nil, nil, Plot{}, err
	}

	return movie, movieInfo, reviews, castInfos, plot, nil
}

func (api *APIServiceImpl) RegisterMovieId(ctx context.Context, reqID int64, movieID string, title string) (MovieId, error) {
	return api.movieIdService.RegisterMovieId(ctx, reqID, movieID, title)
}

func (api *APIServiceImpl) WriteMovieInfo(ctx context.Context, reqID int64, movieID string, title string, casts []Cast, plotID int64, thumbnailIDs []string, photoIDs []string, videoIDs []string, avgRating string, numRating int) (MovieInfo, error) {
	return api.movieInfoService.WriteMovieInfo(ctx, reqID, movieID, title, casts, plotID, thumbnailIDs, photoIDs, videoIDs, avgRating, numRating)
}

func (api *APIServiceImpl) WriteCastInfo(ctx context.Context, reqID int64, castInfoID string, name string, gender string, intro string) (CastInfo, error) {
	return api.castInfoService.WriteCastInfo(ctx, reqID, castInfoID, name, gender, intro)
}

func (api *APIServiceImpl) WritePlot(ctx context.Context, reqID int64, plotID int64, plotText string) (Plot, error) {
	return api.plotService.WritePlot(ctx, reqID, plotID, plotText)
}

func (api *APIServiceImpl) RegisterUser(ctx context.Context, reqID string, firstName string, lastName string, username string, password string) (User, error) {
	return api.userService.RegisterUser(ctx, reqID, firstName, lastName, username, password)
}

func (api *APIServiceImpl) Login(ctx context.Context, reqID int64, username string, password string) error {
	return api.userService.Login(ctx, reqID, username, password)
}

func (api *APIServiceImpl) UploadUserWithUsername(ctx context.Context, reqID int64, username string) error {
	return api.UploadUserWithUsername(ctx, reqID, username)
}

func (api *APIServiceImpl) UploadText(ctx context.Context, reqID int64, text string) error {
	return api.UploadText(ctx, reqID, text)
}

func (api *APIServiceImpl) UploadMovieId(ctx context.Context, reqID int64, title string, rating int) error {
	return api.UploadMovieId(ctx, reqID, title, rating)
}

func (api *APIServiceImpl) UploadUniqueId(ctx context.Context, reqID int64) error {
	return api.UploadUniqueId(ctx, reqID)
}

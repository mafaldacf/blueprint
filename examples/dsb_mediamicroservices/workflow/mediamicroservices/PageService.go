package mediamicroservices

import (
	"context"
)

type PageService interface {
	ReadPage(ctx context.Context, reqID int64, movieID string, reviewStart int, reviewStop int) (MovieInfo, []Review, []CastInfo, Plot, error)
}

type PageServiceImpl struct {
	movieInfoService   MovieInfoService
	movieReviewService MovieReviewService
	castInfoService    CastInfoService
	plotService        PlotService
}

func NewPageServiceImpl(ctx context.Context, movieInfoService MovieInfoService, movieReviewService MovieReviewService, castInfoService CastInfoService, plotService PlotService) (PageService, error) {
	return &PageServiceImpl{movieInfoService: movieInfoService, movieReviewService: movieReviewService, castInfoService: castInfoService, plotService: plotService}, nil
}

func (s *PageServiceImpl) ReadPage(ctx context.Context, reqID int64, movieID string, reviewStart int, reviewStop int) (MovieInfo, []Review, []CastInfo, Plot, error) {
	movieInfo, err := s.movieInfoService.ReadMovieInfo(ctx, reqID, movieID)
	if err != nil {
		return MovieInfo{}, nil, nil, Plot{}, err
	}

	reviews, err := s.movieReviewService.ReadMovieReviews(ctx, reqID, movieID, reviewStart, reviewStop)
	if err != nil {
		return MovieInfo{}, nil, nil, Plot{}, err
	}

	var castInfoIDs []string
	for _, cast := range movieInfo.Casts {
		castInfoIDs = append(castInfoIDs, cast.CastInfoID)
	}

	castInfos, err := s.castInfoService.ReadCastInfos(ctx, reqID, castInfoIDs)
	if err != nil {
		return MovieInfo{}, nil, nil, Plot{}, err
	}

	plot, err := s.plotService.ReadPlot(ctx, reqID, movieInfo.PlotID)
	if err != nil {
		return MovieInfo{}, nil, nil, Plot{}, err
	}

	return movieInfo, reviews, castInfos, plot, nil
}

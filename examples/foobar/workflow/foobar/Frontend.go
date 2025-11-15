package foobar

import (
	"context"
)

type Frontend interface {
	ReadMoviePlot(ctx context.Context, movieID string) (Movie, Plot, error)
	ReadRoutePrice(ctx context.Context, routeID string) (Route, Price, error)
}

type FrontendImpl struct {
	barService   MovieService
	movieService MovieService
	plotService  PlotService
	priceService PriceService
	routeService RouteService
}

func NewFrontendImpl(ctx context.Context, barService MovieService, movieService MovieService, plotService PlotService, priceService PriceService, routeService RouteService) (Frontend, error) {
	f := &FrontendImpl{barService: barService, movieService: movieService, plotService: plotService, priceService: priceService}
	return f, nil
}

func (f *FrontendImpl) ReadMoviePlot(ctx context.Context, movieID string) (Movie, Plot, error) {
	movie, err := f.movieService.ReadMovie(ctx, movieID)
	if err != nil {
		return Movie{}, Plot{}, err
	}
	plot, err := f.plotService.ReadPlot(ctx, movie.PlotID)
	if err != nil {
		return Movie{}, Plot{}, err
	}
	return movie, plot, nil
}

func (f *FrontendImpl) ReadRoutePrice(ctx context.Context, routeID string) (Route, Price, error) {
	route, err := f.routeService.ReadRoute(ctx, routeID)
	if err != nil {
		return Route{}, Price{}, err
	}
	price, err := f.priceService.ReadPriceByRouteID(ctx, routeID)
	if err != nil {
		return Route{}, Price{}, err
	}
	return route, price, nil
}

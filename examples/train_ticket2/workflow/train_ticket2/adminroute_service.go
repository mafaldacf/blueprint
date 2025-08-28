package train_ticket2

import (
	"context"
)

type AdminRouteService interface {
	AddRoute(ctx context.Context, info RouteInfo) (Route, error)
	DeleteRoute(ctx context.Context, id string) error
}

type AdminRouteServiceImpl struct {
	routeService RouteService
}

func NewAdminRouteServiceImpl(ctx context.Context, routeService RouteService) (AdminRouteService, error) {
	return &AdminRouteServiceImpl{routeService: routeService}, nil
}

func (a *AdminRouteServiceImpl) AddRoute(ctx context.Context, info RouteInfo) (Route, error) {
	return a.routeService.CreateAndModify(ctx, info)
}

func (a *AdminRouteServiceImpl) DeleteRoute(ctx context.Context, id string) error {
	return a.routeService.DeleteRoute(ctx, id)
}

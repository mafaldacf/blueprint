package train_ticket2

import (
	"context"
)

type AdminRouteService interface {
	DeleteRoute(ctx context.Context, id string) error
}

type AdminRouteServiceImpl struct {
	routeService RouteService
}

func NewAdminRouteServiceImpl(ctx context.Context, routeService RouteService) (AdminRouteService, error) {
	return &AdminRouteServiceImpl{routeService: routeService}, nil
}

func (a *AdminRouteServiceImpl) DeleteRoute(ctx context.Context, id string) error {
	return a.routeService.DeleteRoute(ctx, id)
}

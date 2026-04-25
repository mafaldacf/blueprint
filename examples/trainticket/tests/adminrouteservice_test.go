package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var adminRouteServiceRegistry = registry.NewServiceRegistry[trainticket.AdminRouteService]("admin_route_service")

func init() {
	adminRouteServiceRegistry.Register("local", func(ctx context.Context) (trainticket.AdminRouteService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		routeService, err := trainticket.NewRouteServiceImpl(ctx, db)
		if err != nil {
			return nil, err
		}
		return trainticket.NewAdminRouteServiceImpl(ctx, routeService)
	})
}

func TestAdminRouteServiceAddAndGetAll(t *testing.T) {
	ctx := context.Background()
	service, err := adminRouteServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.RouteInfo{
		ID:           "adm_route001",
		StartStation: "shanghai",
		EndStation:   "beijing",
		StationList:  "shanghai,nanjing,beijing",
		DistanceList: "0,300,1200",
	}
	route, err := service.AddRoute(ctx, info)
	assert.NoError(t, err)
	assert.Equal(t, "shanghai", route.StartStation)

	all, err := service.GetAllRoutes(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestAdminRouteServiceDeleteRoute(t *testing.T) {
	ctx := context.Background()
	service, err := adminRouteServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.RouteInfo{
		ID:           "adm_route002",
		StartStation: "nanjing",
		EndStation:   "guangzhou",
		StationList:  "nanjing,wuhan,guangzhou",
		DistanceList: "0,500,1200",
	}
	route, err := service.AddRoute(ctx, info)
	assert.NoError(t, err)

	err = service.DeleteRoute(ctx, route.ID)
	assert.NoError(t, err)
}

package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var routeServiceRegistry = registry.NewServiceRegistry[trainticket.RouteService]("route_service")

func init() {
	routeServiceRegistry.Register("local", func(ctx context.Context) (trainticket.RouteService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewRouteServiceImpl(ctx, db)
	})
}

func TestRouteServiceCreateAndFind(t *testing.T) {
	ctx := context.Background()
	service, err := routeServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.RouteInfo{
		ID:           "route001",
		StartStation: "shanghai",
		EndStation:   "beijing",
		StationList:  "shanghai,nanjing,beijing",
		DistanceList: "0,300,1200",
	}
	route, err := service.CreateAndModify(ctx, info)
	assert.NoError(t, err)
	assert.NotEmpty(t, route.ID)
	assert.Equal(t, "shanghai", route.StartStation)
	assert.Equal(t, "beijing", route.EndStation)
	assert.Len(t, route.Stations, 3)
}

func TestRouteServiceGetRouteById(t *testing.T) {
	ctx := context.Background()
	service, err := routeServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.RouteInfo{
		StartStation: "shanghai",
		EndStation:   "hangzhou",
		StationList:  "shanghai,hangzhou",
		DistanceList: "0,200",
	}
	created, err := service.CreateAndModify(ctx, info)
	assert.NoError(t, err)

	route, err := service.GetRouteById(ctx, created.ID)
	assert.NoError(t, err)
	assert.Equal(t, created.ID, route.ID)
	assert.Equal(t, "hangzhou", route.EndStation)
}

func TestRouteServiceGetAllRoutes(t *testing.T) {
	ctx := context.Background()
	service, err := routeServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.RouteInfo{
		ID:           "route003",
		StartStation: "beijing",
		EndStation:   "guangzhou",
		StationList:  "beijing,wuhan,guangzhou",
		DistanceList: "0,1100,2300",
	}
	_, err = service.CreateAndModify(ctx, info)
	assert.NoError(t, err)

	routes, err := service.GetAllRoutes(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, routes)
}

func TestRouteServiceGetRouteByStartAndEnd(t *testing.T) {
	ctx := context.Background()
	service, err := routeServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	info := trainticket.RouteInfo{
		ID:           "route004",
		StartStation: "suzhou",
		EndStation:   "nanjing",
		StationList:  "suzhou,nanjing",
		DistanceList: "0,100",
	}
	_, err = service.CreateAndModify(ctx, info)
	assert.NoError(t, err)

	route, err := service.GetRouteByStartAndEnd(ctx, "suzhou", "nanjing")
	assert.NoError(t, err)
	assert.Equal(t, "suzhou", route.StartStation)
	assert.Equal(t, "nanjing", route.EndStation)
}

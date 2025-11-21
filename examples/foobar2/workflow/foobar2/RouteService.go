package foobar2

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Route struct {
	RouteID string
}

type RouteService interface {
	ReadRoute(ctx context.Context, routeID string) (Route, error)
}

type RouteServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewRouteServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (RouteService, error) {
	d := &RouteServiceImpl{db: db}
	return d, nil
}

func (s *RouteServiceImpl) ReadRoute(ctx context.Context, routeID string) (Route, error) {
	var route Route

	collection, err := s.db.GetCollection(ctx, "route_db", "route")
	if err != nil {
		return Route{}, err
	}

	query := bson.D{{Key: "RouteID", Value: routeID}}
	cursor, err := collection.FindOne(ctx, query)
	if err != nil {
		return Route{}, err
	}

	res, err := cursor.One(ctx, &route)
	if !res || err != nil {
		return Route{}, err
	}

	return route, nil
}

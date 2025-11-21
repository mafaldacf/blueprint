package foobar2

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Price struct {
	PriceID string
	RouteID string
}

type PriceService interface {
	ReadPriceByRouteID(ctx context.Context, routeID string) (Price, error)
}

type PriceServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewPriceServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (PriceService, error) {
	d := &PriceServiceImpl{db: db}
	return d, nil
}

func (s *PriceServiceImpl) ReadPriceByRouteID(ctx context.Context, routeID string) (Price, error) {
	var price Price

	collection, err := s.db.GetCollection(ctx, "price_db", "price")
	if err != nil {
		return Price{}, err
	}

	query := bson.D{{Key: "RouteID", Value: routeID}}
	cursor, err := collection.FindOne(ctx, query)
	if err != nil {
		return Price{}, err
	}

	res, err := cursor.One(ctx, &price)
	if !res || err != nil {
		return Price{}, err
	}

	return price, nil
}

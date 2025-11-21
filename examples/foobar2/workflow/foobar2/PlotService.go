package foobar2

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Plot struct {
	PriceID string
	RouteID string
}

type PlotService interface {
	ReadPlot(ctx context.Context, plotID string) (Plot, error)
}

type PlotServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewPlotServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (PlotService, error) {
	d := &PlotServiceImpl{db: db}
	return d, nil
}

func (s *PlotServiceImpl) ReadPlot(ctx context.Context, plotID string) (Plot, error) {
	var plot Plot

	collection, err := s.db.GetCollection(ctx, "plot_db", "plot")
	if err != nil {
		return Plot{}, err
	}

	query := bson.D{{Key: "PlotID", Value: plotID}}
	cursor, err := collection.FindOne(ctx, query)
	if err != nil {
		return Plot{}, err
	}

	res, err := cursor.One(ctx, &plot)
	if !res || err != nil {
		return Plot{}, err
	}

	return plot, nil
}

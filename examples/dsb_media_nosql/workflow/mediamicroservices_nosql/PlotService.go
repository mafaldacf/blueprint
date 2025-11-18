package mediamicroservices_nosql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Plot struct {
	PlotID string `bson:"_id"`
	Plot   string
}

type PlotService interface {
	WritePlot(ctx context.Context, reqID int64, plotID string, plotText string) (Plot, error)
	ReadPlot(ctx context.Context, reqID int64, plotID string) (Plot, error)
}

type PlotServiceImpl struct {
	database backend.NoSQLDatabase
}

func NewPlotServiceImpl(ctx context.Context, database backend.NoSQLDatabase) (PlotService, error) {
	s := &PlotServiceImpl{database: database}
	return s, nil
}

func (s *PlotServiceImpl) WritePlot(ctx context.Context, reqID int64, plotID string, plotText string) (Plot, error) {
	plot := Plot{
		PlotID: plotID,
		Plot:   plotText,
	}
	collection, err := s.database.GetCollection(ctx, "plot_db", "plot")
	if err != nil {
		return Plot{}, err
	}
	err = collection.InsertOne(ctx, plot)
	if err != nil {
		return Plot{}, err
	}

	return plot, err
}

func (s *PlotServiceImpl) ReadPlot(ctx context.Context, reqID int64, plotID string) (Plot, error) {
	var plot Plot
	collection, err := s.database.GetCollection(ctx, "plot_db", "plot")
	if err != nil {
		return plot, err
	}
	query := bson.D{{Key: "PlotID", Value: plotID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return plot, err
	}
	res, err := result.One(ctx, &plot)
	if !res || err != nil {
		return Plot{}, err
	}
	return plot, err
}

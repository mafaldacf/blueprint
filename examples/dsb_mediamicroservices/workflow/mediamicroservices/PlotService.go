package mediamicroservices

import (
	"context"
	"fmt"
	"strconv"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Plot struct {
	PlotID int64 `bson:"_id"`
	Plot   string
}

type PlotService interface {
	WritePlot(ctx context.Context, reqID int64, plotID int64, plotText string) (Plot, error)
	ReadPlot(ctx context.Context, reqID int64, plotID int64) (Plot, error)
}

type PlotServiceImpl struct {
	database backend.NoSQLDatabase
	cache    backend.Cache
}

func NewPlotServiceImpl(ctx context.Context, database backend.NoSQLDatabase, cache backend.Cache) (PlotService, error) {
	s := &PlotServiceImpl{database: database, cache: cache}
	return s, nil
}

func (s *PlotServiceImpl) WritePlot(ctx context.Context, reqID int64, plotID int64, plotText string) (Plot, error) {
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

func (s *PlotServiceImpl) ReadPlot(ctx context.Context, reqID int64, plotID int64) (Plot, error) {
	var plot Plot

	var cachedPlot interface{}
	plotIDStr := strconv.FormatInt(plotID, 10)
	found, err := s.cache.Get(ctx, plotIDStr, &cachedPlot)
	if err != nil {
		return Plot{}, err
	}

	if found {
		// if cached in memcached
		plot = cachedPlot.(Plot)
		return plot, nil
	}
	
	// if not cached in memcached
	collection, err := s.database.GetCollection(ctx, "plot_db", "plot")
	if err != nil {
		return plot, err
	}
	query := bson.D{{Key: "PlotID", Value: plotID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return plot, err
	}
	found, err = result.One(ctx, &plot)
	if err != nil {
		return Plot{}, err
	}
	if !found {
		return Plot{}, fmt.Errorf("plot %d not found", plotID)
	}
	return plot, err
}

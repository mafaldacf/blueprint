package mediamicroservices_sql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type Plot struct {
	plotid string `bson:"_id"`
	plot   string `bson:"plot"`
}

type PlotService interface {
	WritePlot(ctx context.Context, reqID int64, plotID string, plotText string) (Plot, error)
	ReadPlot(ctx context.Context, reqID int64, plotID string) (Plot, error)
}

type PlotServiceImpl struct {
	PlotDB backend.RelationalDB
}

func NewPlotServiceImpl(ctx context.Context, PlotDB backend.RelationalDB) (PlotService, error) {
	m := &PlotServiceImpl{PlotDB: PlotDB}
	return m, nil
}

func (m *PlotServiceImpl) WritePlot(ctx context.Context, reqID int64, plotID string, plotText string) (Plot, error) {
	plot := Plot{
		plotid: plotID,
		plot:   plotText,
	}
	_, err := m.PlotDB.Exec(ctx, "INSERT INTO plot(plotid, plot) VALUES (?, ?);", plotID, plotText)
	return plot, err
}

func (m *PlotServiceImpl) ReadPlot(ctx context.Context, reqID int64, plotID string) (Plot, error) {
	var Plot Plot
	err := m.PlotDB.Select(ctx, &Plot, "SELECT * FROM plot WHERE plotid = ?", plotID)
	return Plot, err
}

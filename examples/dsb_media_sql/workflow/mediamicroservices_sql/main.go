package mediamicroservices_sql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var movieIdDB backend.RelationalDB
	var movieInfoDB backend.RelationalDB
	var castInfoDB backend.RelationalDB
	var plotDB backend.RelationalDB

	movieIdService, _ := NewMovieIdServiceImpl(ctx, movieIdDB)
	movieInfoService, _ := NewMovieInfoServiceImpl(ctx, movieInfoDB)
	castInfoService, _ := NewCastInfoServiceImpl(ctx, castInfoDB)
	plotService, _ := NewPlotServiceImpl(ctx, plotDB)
	api, _ := NewAPIServiceImpl(ctx, movieIdService, movieInfoService, castInfoService, plotService)

	var reqID int64
	var title string
	api.ReadPage(ctx, reqID, title)
}

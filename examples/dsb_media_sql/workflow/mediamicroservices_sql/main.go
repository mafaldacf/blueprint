package mediamicroservices_sql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var movieIdDB backend.RelationalDB
	var movieInfoDB backend.RelationalDB

	movieIdService, _ := NewMovieIdServiceImpl(ctx, movieIdDB)
	movieInfoService, _ := NewMovieInfoServiceImpl(ctx, movieInfoDB)
	api, _ := NewAPIServiceImpl(ctx, movieIdService, movieInfoService)

	var reqID int64
	var movieID string
	api.ReadMovie(ctx, reqID, movieID)
}

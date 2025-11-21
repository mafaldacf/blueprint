package mediamicroservices_nosql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()
	var database backend.NoSQLDatabase
	var cache backend.Cache
	service, _ := NewPlotServiceImpl(ctx, database, cache)
	var reqID int64
	var plotID int64
	var text string
	service.WritePlot(ctx, reqID, plotID, text)
}

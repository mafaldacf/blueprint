package mediamicroservices_nosql

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()
	var db backend.NoSQLDatabase
	service, _ := NewPlotServiceImpl(ctx, db)
	var reqID int64
	var plotID string
	var text string
	service.WritePlot(ctx, reqID, plotID, text)
}

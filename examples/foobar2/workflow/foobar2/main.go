package foobar2

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var db backend.NoSQLDatabase
	service, _ := NewRouteServiceImpl(ctx, db)

	var id string
	service.ReadRoute(ctx, id)
}

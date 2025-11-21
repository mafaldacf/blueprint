package foobar

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var barDB backend.NoSQLDatabase
	barService, _ := NewBarServiceImpl(ctx, barDB)

	var id, text, barID string
	barService.WriteBar(ctx, id, text, barID)
}

package digota

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var skusDB backend.NoSQLDatabase
	var queue backend.Queue
	skuService, _ := NewSkuServiceImpl(ctx, skusDB, queue)

	var id string
	skuService.Get(ctx, id)
}

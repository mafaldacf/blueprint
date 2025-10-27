package large_scale_app

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var service100DB backend.NoSQLDatabase
	service100, _ := NewService500Impl(ctx, service100DB)

	var id string
	var data string
	var datatwo string
	service100.Method500(ctx, id, data, datatwo)
}

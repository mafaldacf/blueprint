package large_scale_app

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var db backend.NoSQLDatabase
	service, _ := NewService100Impl(ctx, db)

	var id string
	var data string
	service.Method100(ctx, id, data, "")
}

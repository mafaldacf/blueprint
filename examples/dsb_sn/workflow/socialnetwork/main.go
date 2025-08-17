package socialnetwork

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var postCache backend.Cache
	var postDB backend.NoSQLDatabase
	postStorageService, _ := NewPostStorageServiceImpl(ctx, postCache, postDB)

	var reqID int64
	var post Post
	postStorageService.StorePost(ctx, reqID, post)
}

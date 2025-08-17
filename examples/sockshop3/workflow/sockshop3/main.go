package sockshop3

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

func main() {
	ctx := context.Background()

	var userDB backend.NoSQLDatabase
	userService, _ := NewUserServiceImpl(ctx, userDB)

	var username, password string
	userService.Login(ctx, username, password)
}

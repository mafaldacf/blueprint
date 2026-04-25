package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_mediamicroservices/workflow/mediamicroservices"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var userServiceRegistry = registry.NewServiceRegistry[mediamicroservices.UserService]("user_service")

func init() {
	userServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.UserService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		composeReviewService, err := composeReviewServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewUserServiceImpl(ctx, db, cache, composeReviewService)
	})
}

func TestUserServiceRegister(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user, err := service.RegisterUser(ctx, "req001", "John", "Doe", "johndoe", "password123")
	assert.NoError(t, err)
	assert.Equal(t, "johndoe", user.Username)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.NotEmpty(t, user.Salt)
	assert.NotEmpty(t, user.Password)
	// password is stored hashed, not in plaintext
	assert.NotEqual(t, "password123", user.Password)
}

func TestUserServiceRegisterWithId(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user, err := service.RegisterUserWithId(ctx, "req002", "Jane", "Smith", "janesmith", "securepass", 42)
	assert.NoError(t, err)
	assert.Equal(t, int64(42), user.UserID)
	assert.Equal(t, "janesmith", user.Username)
	assert.Equal(t, "Jane", user.FirstName)
	assert.Equal(t, "Smith", user.LastName)
}

func TestUserServiceLogin(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.RegisterUser(ctx, "req003", "Alice", "Wonder", "alicew", "mypassword")
	assert.NoError(t, err)

	err = service.Login(ctx, 0, "alicew", "mypassword")
	assert.NoError(t, err)
}

func TestUserServiceRegisterDifferentUsers(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user1, err := service.RegisterUser(ctx, "req004", "Bob", "Builder", "bobbuilder", "pass1")
	assert.NoError(t, err)
	assert.Equal(t, "bobbuilder", user1.Username)

	user2, err := service.RegisterUser(ctx, "req005", "Carol", "Singer", "carolsinger", "pass2")
	assert.NoError(t, err)
	assert.Equal(t, "carolsinger", user2.Username)

	assert.NotEqual(t, user1.UserID, user2.UserID)
}

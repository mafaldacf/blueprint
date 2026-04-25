package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_socialnetwork/workflow/socialnetwork"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var userServiceRegistry = registry.NewServiceRegistry[socialnetwork.UserService]("user_service")

func init() {
	userServiceRegistry.Register("local", func(ctx context.Context) (socialnetwork.UserService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		socialGraphService, err := socialGraphServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return socialnetwork.NewUserServiceImpl(ctx, cache, db, socialGraphService, "test_secret")
	})
}

func TestUserServiceRegister(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.RegisterUser(ctx, 0, "Alice", "Smith", "alicesmith", "password123")
	assert.NoError(t, err)
}

func TestUserServiceRegisterWithId(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.RegisterUserWithId(ctx, 0, "Bob", "Jones", "bobjones", "pass456", 100)
	assert.NoError(t, err)
}

func TestUserServiceRegisterDuplicateUsername(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.RegisterUserWithId(ctx, 0, "Carol", "White", "carolwhite", "mypass", 200)
	assert.NoError(t, err)

	err = service.RegisterUserWithId(ctx, 0, "Carol", "White", "carolwhite", "mypass", 201)
	assert.Error(t, err)
}

func TestUserServiceLogin(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.RegisterUserWithId(ctx, 0, "Dave", "Brown", "davebrown", "securepass", 300)
	assert.NoError(t, err)

	token, err := service.Login(ctx, 0, "davebrown", "securepass")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestUserServiceLoginWrongPassword(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.RegisterUserWithId(ctx, 0, "Eve", "Green", "evegreen", "correctpass", 400)
	assert.NoError(t, err)

	_, err = service.Login(ctx, 0, "evegreen", "wrongpass")
	assert.Error(t, err)
}

func TestUserServiceComposeCreatorWithUserId(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	creator, err := service.ComposeCreatorWithUserId(ctx, 0, 42, "frank")
	assert.NoError(t, err)
	assert.Equal(t, int64(42), creator.UserID)
	assert.Equal(t, "frank", creator.Username)
}

func TestUserServiceComposeCreatorWithUsername(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	err = service.RegisterUserWithId(ctx, 0, "Grace", "Hall", "gracehall", "gracepass", 500)
	assert.NoError(t, err)

	creator, err := service.ComposeCreatorWithUsername(ctx, 0, "gracehall")
	assert.NoError(t, err)
	assert.Equal(t, int64(500), creator.UserID)
	assert.Equal(t, "gracehall", creator.Username)
}

func TestUserServiceComposeCreatorWithUsernameNotFound(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.ComposeCreatorWithUsername(ctx, 0, "nonexistentuser")
	assert.Error(t, err)
}

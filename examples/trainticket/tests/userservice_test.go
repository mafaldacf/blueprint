package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var userServiceRegistry = registry.NewServiceRegistry[trainticket.UserService]("user_service")

func init() {
	userServiceRegistry.Register("local", func(ctx context.Context) (trainticket.UserService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewUserServiceImpl(ctx, db)
	})
}

func TestUserServiceSaveAndFind(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user := trainticket.User{
		UserID:   "user001",
		Username: "johndoe",
		Password: "password123",
		Gender:   1,
		Email:    "johndoe@example.com",
	}
	err = service.SaveUser(ctx, user)
	assert.NoError(t, err)

	found, err := service.FindByUsername(ctx, "johndoe")
	assert.NoError(t, err)
	assert.Equal(t, "user001", found.UserID)
	assert.Equal(t, "johndoe", found.Username)
	assert.Equal(t, "johndoe@example.com", found.Email)
}

func TestUserServiceFindByUserID(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user := trainticket.User{
		UserID:   "user002",
		Username: "janesmith",
		Password: "securepass",
		Gender:   0,
		Email:    "jane@example.com",
	}
	err = service.SaveUser(ctx, user)
	assert.NoError(t, err)

	found, err := service.FindByUserID(ctx, "user002")
	assert.NoError(t, err)
	assert.Equal(t, "janesmith", found.Username)
	assert.Equal(t, "jane@example.com", found.Email)
}

func TestUserServiceGetAllUsers(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user1 := trainticket.User{UserID: "user003", Username: "alice", Password: "pass1", Email: "alice@example.com"}
	user2 := trainticket.User{UserID: "user004", Username: "bob", Password: "pass2", Email: "bob@example.com"}
	err = service.SaveUser(ctx, user1)
	assert.NoError(t, err)
	err = service.SaveUser(ctx, user2)
	assert.NoError(t, err)

	users, err := service.GetAllUsers(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}

func TestUserServiceUpdateUser(t *testing.T) {
	ctx := context.Background()
	service, err := userServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user := trainticket.User{
		UserID:   "user005",
		Username: "charlie",
		Password: "oldpass",
		Email:    "charlie@example.com",
	}
	err = service.SaveUser(ctx, user)
	assert.NoError(t, err)

	user.Email = "charlie_new@example.com"
	updated, err := service.UpdateUser(ctx, user)
	assert.NoError(t, err)
	assert.True(t, updated)
}

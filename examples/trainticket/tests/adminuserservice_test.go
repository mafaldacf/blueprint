package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var adminUserServiceRegistry = registry.NewServiceRegistry[trainticket.AdminUserService]("admin_user_service")

func init() {
	adminUserServiceRegistry.Register("local", func(ctx context.Context) (trainticket.AdminUserService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		userService, err := trainticket.NewUserServiceImpl(ctx, db)
		if err != nil {
			return nil, err
		}
		return trainticket.NewAdminUserServiceImpl(ctx, userService)
	})
}

func TestAdminUserServiceAddAndGetAll(t *testing.T) {
	ctx := context.Background()
	service, err := adminUserServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user := trainticket.User{
		UserID:   "adm_user001",
		Username: "alice",
		Email:    "alice@example.com",
	}
	err = service.AddUser(ctx, user)
	assert.NoError(t, err)

	all, err := service.GetAllUsers(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestAdminUserServiceUpdateUser(t *testing.T) {
	ctx := context.Background()
	service, err := adminUserServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user := trainticket.User{
		UserID:   "adm_user002",
		Username: "bob",
		Email:    "bob@example.com",
	}
	err = service.AddUser(ctx, user)
	assert.NoError(t, err)

	user.Email = "bob_new@example.com"
	updated, err := service.UpdateUser(ctx, user)
	assert.NoError(t, err)
	assert.True(t, updated)
}

func TestAdminUserServiceDeleteUser(t *testing.T) {
	ctx := context.Background()
	service, err := adminUserServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	user := trainticket.User{
		UserID:   "adm_user003",
		Username: "carol",
	}
	err = service.AddUser(ctx, user)
	assert.NoError(t, err)

	err = service.DeleteUser(ctx, "adm_user003")
	assert.NoError(t, err)
}

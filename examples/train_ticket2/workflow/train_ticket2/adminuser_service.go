package train_ticket2

import (
	"context"
)

type AdminUserService interface {
	AddUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, userID string) error
}

type AdminUserServiceImpl struct {
	userService UserService
}

func NewAdminUserServiceImpl(ctx context.Context, userService UserService) (AdminUserService, error) {
	return &AdminUserServiceImpl{userService: userService}, nil
}

func (a *AdminUserServiceImpl) AddUser(ctx context.Context, user User) error {
	return a.userService.SaveUser(ctx, user)
}

func (a *AdminUserServiceImpl) DeleteUser(ctx context.Context, userID string) error {
	return a.userService.DeleteUser(ctx, userID)
}

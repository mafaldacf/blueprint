package app_constraints_referential_integrity

import (
	"context"
)

type Frontend interface {
	CreateAccount(ctx context.Context, reqID string, accountID string, text string) (Account, error)
	ReadAccount(ctx context.Context, accountID string) (Account, error)
	DeleteAccount(ctx context.Context, accountID string) (bool, error)
	ReadAccountUser(ctx context.Context, accountID string) (User, error)
	DeleteUser(ctx context.Context, username string) (bool, error)
}

type FrontendImpl struct {
	accountService AccountService
	userService    UserService
}

func NewFrontendImpl(ctx context.Context, accountService AccountService, userService UserService) (Frontend, error) {
	return &FrontendImpl{accountService: accountService, userService: userService}, nil
}

func (u *FrontendImpl) CreateAccount(ctx context.Context, reqID string, accountID string, username string) (Account, error) {
	account, err := u.accountService.CreateAccount(ctx, reqID, accountID, username)
	return account, err
}

func (u *FrontendImpl) ReadAccount(ctx context.Context, accountID string) (Account, error) {
	account, err := u.accountService.GetAccount(ctx, accountID)
	return account, err
}

func (u *FrontendImpl) DeleteAccount(ctx context.Context, accountID string) (bool, error) {
	ok, err := u.accountService.DeleteAccount(ctx, accountID)
	return ok, err
}

func (u *FrontendImpl) DeleteUser(ctx context.Context, username string) (bool, error) {
	ok, err := u.userService.DeleteUser(ctx, username)
	return ok, err
}

func (u *FrontendImpl) ReadAccountUser(ctx context.Context, accountID string) (User, error) {
	/* account, err := u.accountService.GetAccount(ctx, accountID)
	if err != nil {
		return User{}, err
	}
	username := account.Username
	user, err := u.userService.GetUser(ctx, username)
	return user, err */
	return u.accountService.ReadAccountUser(ctx, accountID)
}

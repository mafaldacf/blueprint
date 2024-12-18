package app_constraints_referential_integrity

import (
	"context"
	"math/rand"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type AccountService interface {
	CreateAccount(ctx context.Context, reqID string, accountID string, username string) (Account, error)
	GetAccount(ctx context.Context, accountID string) (Account, error)
	DeleteAccount(ctx context.Context, accountID string) (bool, error)
	//AddAccountUser(ctx context.Context, accountID string, username string) (AccountUsers, error)
	ReadAccountUser(ctx context.Context, accountID string) (User, error)
}

type AccountServiceImpl struct {
	userService UserService
	accountsDB  backend.NoSQLDatabase
}

func NewAccountServiceImpl(ctx context.Context, userService UserService, accountsDB backend.NoSQLDatabase) (AccountService, error) {
	s := &AccountServiceImpl{userService: userService, accountsDB: accountsDB}
	return s, nil
}

func (s *AccountServiceImpl) CreateAccount(ctx context.Context, reqID string, accountID string, username string) (Account, error) {
	_, err := s.userService.CreateUser(ctx, reqID, username)
	if err != nil {
		return Account{}, err
	}

	timestamp := rand.Int63()
	account := Account{
		ReqID:     reqID,
		AccountID: accountID,
		Username:  username,
		Timestamp: timestamp,
	}
	collection, err := s.accountsDB.GetCollection(ctx, "accounts", "accounts")
	if err != nil {
		return account, err
	}
	err = collection.InsertOne(ctx, account)
	if err != nil {
		return account, err
	}

	/* accountUsers := AccountUsers{
		AccountID: accountID,
		Usernames: []string{username},
	}
	collection, err = s.accountsDB.GetCollection(ctx, "accounts", "accounts_users")
	if err != nil {
		return account, err
	}
	err = collection.InsertOne(ctx, accountUsers) */
	return account, err
}

/* func (s *AccountServiceImpl) AddAccountUser(ctx context.Context, accountID string, username string) (AccountUsers, error) {
	var accountUsers AccountUsers
	collection, err := s.accountsDB.GetCollection(ctx, "accounts", "accounts_users")
	if err != nil {
		return accountUsers, err
	}
	query := bson.D{{Key: "accountID", Value: accountID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return accountUsers, err
	}
	res, err := result.One(ctx, &accountUsers)
	if !res || err != nil {
		return accountUsers, err
	}

	accountUsers.Usernames = append(accountUsers.Usernames, username)
	err = collection.InsertOne(ctx, accountUsers)
	return accountUsers, err
} */

func (s *AccountServiceImpl) GetAccount(ctx context.Context, accountID string) (Account, error) {
	var account Account
	collection, err := s.accountsDB.GetCollection(ctx, "accounts", "accounts")
	if err != nil {
		return account, err
	}
	query := bson.D{{Key: "accountID", Value: accountID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return account, err
	}
	res, err := result.One(ctx, &account)
	if !res || err != nil {
		return account, err
	}
	return account, err
}

func (s *AccountServiceImpl) DeleteAccount(ctx context.Context, accountID string) (bool, error) {
	collection, err := s.accountsDB.GetCollection(ctx, "accounts", "accounts")
	if err != nil {
		return false, err
	}
	filter := bson.D{{Key: "accountID", Value: accountID}}
	err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err
}

func (s *AccountServiceImpl) ReadAccountUser(ctx context.Context, accountID string) (User, error) {
	var account Account
	collection, err := s.accountsDB.GetCollection(ctx, "accounts", "accounts")
	if err != nil {
		return User{}, err
	}
	query := bson.D{{Key: "accountID", Value: accountID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return User{}, err
	}
	res, err := result.One(ctx, &account)
	if !res || err != nil {
		return User{}, err
	}
	return s.userService.GetUser(ctx, account.Username)
}

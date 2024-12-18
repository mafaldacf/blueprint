package app_constraints_referential_integrity

import (
	"context"
	"math/rand"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService interface {
	CreateUser(ctx context.Context, reqID string, username string) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
	DeleteUser(ctx context.Context, username string) (bool, error)
}

type UserServiceImpl struct {
	usersDB backend.NoSQLDatabase
}

func NewUserServiceImpl(ctx context.Context, usersDB backend.NoSQLDatabase) (UserService, error) {
	s := &UserServiceImpl{usersDB: usersDB}
	return s, nil
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, reqID string, username string) (User, error) {
	timestamp := rand.Int63()
	user := User{
		ReqID:     reqID,
		Username:  username,
		Timestamp: timestamp,
	}

	collection, err := s.usersDB.GetCollection(ctx, "users", "users")
	if err != nil {
		return user, err
	}
	err = collection.InsertOne(ctx, user)
	return user, err
}

func (s *UserServiceImpl) GetUser(ctx context.Context, username string) (User, error) {
	var user User
	collection, err := s.usersDB.GetCollection(ctx, "users", "users")
	if err != nil {
		return user, err
	}
	query := bson.D{{Key: "username", Value: username}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return user, err
	}
	res, err := result.One(ctx, &user)
	if !res || err != nil {
		return user, err
	}
	return user, err
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, username string) (bool, error) {
	collection, err := s.usersDB.GetCollection(ctx, "users", "users")
	if err != nil {
		return false, err
	}
	filter := bson.D{{Key: "username", Value: username}}
	err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err
}

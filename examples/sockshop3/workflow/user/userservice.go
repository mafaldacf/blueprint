// Package user implements the SockShop user microservice.
//
// The service stores three kinds of information:
//   - user accounts
//   - addresses
//   - credit cards
//
// The sock shop allows customers to check out without creating a user
// account; in this case the customer's address and credit card data
// will be stored without a user accont.
//
// The UserService thus uses three collections for the above information.
// To get the data for a user also means more than one database call.
package user

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type UserService interface {
	Login(ctx context.Context, username, password string) (User, error)
	Register(ctx context.Context, username, password, email, first, last string) (string, error)
	GetUsers(ctx context.Context, id string) ([]User, error)
	PostUser(ctx context.Context, user User) (string, error)
	GetAddresses(ctx context.Context, id string) ([]Address, error)
	PostAddress(ctx context.Context, userid string, address Address) (string, error)
	GetCards(ctx context.Context, cardid string) ([]Card, error)
	PostCard(ctx context.Context, userid string, card Card) (string, error)
	Delete(ctx context.Context, entity string, id string) error
}
type UserServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewUserServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (UserService, error) {
	return &UserServiceImpl{db: db}, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, username, password string) (User, error) {
	// Load the u from the DB
	u, err := s.userdb_GetUserByName(ctx, username)
	if err != nil {
		return u, err
	}

	// Check the password
	if u.Password != calculatePassHash(password, u.Salt) {
		return User{}, errors.New("Unauthorized")
	}

	// Fetch user's card and address data, mask out CC numbers
	err = s.userdb_GetUserAttributes(ctx, &u)
	if err != nil {
		return u, err
	}
	u.maskCCs()
	return u, nil
}

func (s *UserServiceImpl) Register(ctx context.Context, username, password, email, first, last string) (string, error) {
	// Create the public user info
	u := User{}
	u.Username = username
	u.Password = calculatePassHash(password, u.Salt)
	u.Email = email
	u.FirstName = first
	u.LastName = last
	u.Addresses = []Address{}
	u.Cards = []Card{}

	// Save the user in the DB
	err := s.userdb_CreateUser(ctx, &u)
	return u.UserID, err
}

func (s *UserServiceImpl) GetUsers(ctx context.Context, userid string) ([]User, error) {
	if userid == "" {
		return s.userdb_GetUsers(ctx)
	}
	user, err := s.userdb_GetUser(ctx, userid)
	return []User{user}, err

}

func (s *UserServiceImpl) PostUser(ctx context.Context, u User) (string, error) {
	u.newSalt()
	u.Password = calculatePassHash(u.Password, u.Salt)
	err := s.userdb_CreateUser(ctx, &u)
	return u.UserID, err
}

func (s *UserServiceImpl) GetAddresses(ctx context.Context, addressid string) ([]Address, error) {
	if addressid == "" {
		return s.addressdb_GetAllAddresses(ctx)
	}
	address, err := s.addressdb_GetAddress(ctx, addressid)
	return []Address{address}, err
}

func (s *UserServiceImpl) PostAddress(ctx context.Context, userid string, address Address) (string, error) {
	err := s.userdb_CreateAddress(ctx, userid, &address)
	return address.ID, err
}

func (s *UserServiceImpl) GetCards(ctx context.Context, cardid string) ([]Card, error) {
	if cardid == "" {
		return s.carddb_GetAllCards(ctx)
	}
	card, err := s.carddb_GetCard(ctx, cardid)
	return []Card{card}, err
}

func (s *UserServiceImpl) PostCard(ctx context.Context, userid string, card Card) (string, error) {
	err := s.userdb_CreateCard(ctx, userid, &card)
	return card.ID, err
}

func (s *UserServiceImpl) Delete(ctx context.Context, entity string, id string) error {
	return s.userdb_Delete(ctx, entity, id)
}

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}

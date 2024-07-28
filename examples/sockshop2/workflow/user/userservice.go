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
	"errors"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService interface {
	// Log in to an existing user account.  Returns an error if the password
	// doesn't match the registered password
	Login(ctx context.Context, username, password string) (User, error)

	// Register a new user account.
	// Returns the user ID
	Register(ctx context.Context, username, password, email, first, last string) (string, error)

	// Look up a user by id.  If id is the empty string, returns all users.
	GetUsers(ctx context.Context, id string) ([]User, error)

	// Insert a (possibly new) user into the DB.  Returns the user's ID
	PostUser(ctx context.Context, user User) (string, error)

	// Look up an address by id.  If id is the empty string, returns all addresses.
	GetAddresses(ctx context.Context, id string) ([]Address, error)

	// Insert a (possibly new) address into the DB.  Returns the address ID
	PostAddress(ctx context.Context, userid string, address Address) (string, error)

	// Look up a card by id.  If id is the empty string, returns all cards.
	GetCards(ctx context.Context, cardid string) ([]Card, error)

	// Insert a (possibly new) card into the DB.  Returns the card ID
	PostCard(ctx context.Context, userid string, card Card) (string, error)

	// Deletes an entity with ID id from the DB.
	//
	// entity can be one of "customers", "addresses", or "cards".
	// ID should be the id of the entity to delete
	Delete(ctx context.Context, entity string, id string) error
}

// An implementation of the UserService that stores information in a NoSQLDatabase.
// It uses three collections within the database: users, addresses, and cards.
// Addresses and cards are stored separately from user information, because having
// a user account is optional when placing an order.
type userServiceImpl struct {
	UserService
	db backend.NoSQLDatabase
}

// Creates a UserService implementation that stores user, address, and credit card
// information in a NoSQLDatabase.
//
// Returns an error if unable to get the users, addresses, or cards collection from the DB
func NewUserServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (UserService, error) {
	return &userServiceImpl{db: db}, nil
}

func (s *userServiceImpl) Login(ctx context.Context, username, password string) (User, error) {
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	query := bson.D{{Key: "username", Value: username}}
	projection := bson.D{{Key: "password", Value: true}}
	var user User
	result, _ := collection.FindOne(ctx, query, projection)
	result.One(ctx, &user)

	// Check the password
	if user.Password != password {
		return User{}, errors.New("Unauthorized")
	}
	return user, nil
}

func (s *userServiceImpl) Register(ctx context.Context, username, password, email, first, last string) (string, error) {
	// Create the public user info
	u := User{}
	u.Username = username
	u.Password = password
	u.Email = email
	u.FirstName = first
	u.LastName = last
	u.Addresses = Address{}
	u.Cards = Card{}

	// Save the user in the DB
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	collection.InsertOne(ctx, u)
	return u.UserID, nil
}

func (s *userServiceImpl) GetUsers(ctx context.Context, userid string) ([]User, error) {
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	var users []User
	filter := bson.D{{Key: "userid", Value: userid}}
	result, _ := collection.FindMany(ctx, filter)
	result.All(ctx, &users)
	return users, nil
}

func (s *userServiceImpl) PostUser(ctx context.Context, u User) (string, error) {
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	collection.InsertOne(ctx, u)
	return u.UserID, nil
}

func (s *userServiceImpl) GetAddresses(ctx context.Context, addressid string) ([]Address, error) {
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	var addresses []Address
	filter := bson.D{{Key: "addressid", Value: addressid}}
	projection := bson.D{{Key: "addressid", Value: true}}
	result, _ := collection.FindMany(ctx, filter, projection)
	result.All(ctx, &addresses)
	return addresses, nil
}

func (s *userServiceImpl) PostAddress(ctx context.Context, userid string, address Address) (string, error) {
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	filter := bson.D{{Key: "userid", Value: userid}}
	update := bson.D{{Key: "address", Value: address}}
	collection.Upsert(ctx, filter, update)
	return userid, nil
}

func (s *userServiceImpl) GetCards(ctx context.Context, cardid string) ([]Card, error) {
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	var cards []Card
	filter := bson.D{{Key: "addressid", Value: cardid}}
	projection := bson.D{{Key: "addressid", Value: true}}
	result, _ := collection.FindMany(ctx, filter, projection)
	result.All(ctx, &cards)
	return cards, nil
}

func (s *userServiceImpl) PostCard(ctx context.Context, userid string, card Card) (string, error) {
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	filter := bson.D{{Key: "userid", Value: userid}}
	update := bson.D{{Key: "card", Value: card}}
	collection.Upsert(ctx, filter, update)
	return userid, nil
}

func (s *userServiceImpl) Delete(ctx context.Context, entity string, id string) error {
	collection, _ := s.db.GetCollection(ctx, "users", "users")
	query := bson.D{{Key: "id", Value: id}}
	collection.DeleteOne(ctx, query)
	return nil
}

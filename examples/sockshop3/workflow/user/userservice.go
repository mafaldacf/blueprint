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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	collection, err := s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return User{}, err
	}
	
	// get user by name
	query := bson.D{{Key: "Username", Value: username}}
	projection := bson.D{{Key: "Password", Value: true}}
	var user User
	result, err := collection.FindOne(ctx, query, projection)
	if err != nil {
		return User{}, err
	}
	_, err  = result.One(ctx, &user)
	if err != nil {
		return User{}, err
	}

	// Check the password
	if user.Password != calculatePassHash(password, user.Salt) {
		return User{}, errors.New("Unauthorized")
	}

	// get user attributes
	// TODO

	return user, nil
}

func (s *UserServiceImpl) Register(ctx context.Context, username, password, email, first, last string) (string, error) {
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
	collection, err := s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return "", err
	}
	err = collection.InsertOne(ctx, u)
	if err != nil {
		return "", err
	}
	return u.UserID, nil
}

func (s *UserServiceImpl) GetUsers(ctx context.Context, userid string) ([]User, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return nil, err
	}
	
	var users []User
	filter := bson.D{{Key: "UserID", Value: userid}}
	result, _ := collection.FindMany(ctx, filter)
	result.All(ctx, &users)
	return users, nil
}

func (s *UserServiceImpl) PostUser(ctx context.Context, u User) (string, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return "", err
	}

	collection.InsertOne(ctx, u)
	return u.UserID, nil
}

func (s *UserServiceImpl) GetAddresses(ctx context.Context, addressid string) ([]Address, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return nil, err
	}

	var addresses []Address
	filter := bson.D{{Key: "Addresses", Value: addressid}}
	projection := bson.D{{Key: "Addresses", Value: true}}
	result, _ := collection.FindMany(ctx, filter, projection)
	result.All(ctx, &addresses)
	return addresses, nil
}

func (s *UserServiceImpl) PostAddress(ctx context.Context, userid string, address Address) (string, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return "", err
	}

	filter := bson.D{{Key: "UserID", Value: userid}}
	update := bson.D{{Key: "Address", Value: address}}
	collection.Upsert(ctx, filter, update)
	return userid, nil
}

func (s *UserServiceImpl) GetCards(ctx context.Context, cardid string) ([]Card, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return nil, err
	}

	var cards []Card
	filter := bson.D{{Key: "Cards", Value: cardid}}
	projection := bson.D{{Key: "Cards", Value: true}}
	result, _ := collection.FindMany(ctx, filter, projection)
	result.All(ctx, &cards)
	return cards, nil
}

func (s *UserServiceImpl) PostCard(ctx context.Context, userid string, card Card) (string, error) {
	if userid == "" {
		// An anonymous user; simply insert the card to the DB
		dbcard := dbCard{Card: card, ID: primitive.NewObjectID()}
		collection, err := s.db.GetCollection(ctx, "user_db", "cards")
		if err != nil {
			return "", err
		}
		if _, err := collection.UpsertID(ctx, dbcard.ID, card); err != nil {
			return dbcard.ID.String(), err
		}

		// Update the provided card
		dbcard.Card.ID = dbcard.ID.Hex()
		card = dbcard.Card
		return card.ID, err
	}

	// A userid is provided; first check it's valid
	id, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return "", errors.New("invalid ID Hex")
	}

	// Insert the card to the DB
	collection, err := s.db.GetCollection(ctx, "user_db", "cards")
	if err != nil {
		return "", err
	}
	dbcard := dbCard{Card: card, ID: primitive.NewObjectID()}
	collection.UpsertID(ctx, dbcard.ID, dbcard)

	// Update the provided card
	dbcard.Card.ID = dbcard.ID.Hex()
	card = dbcard.Card

	// Update the user
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$addToSet", Value: bson.D{{"cards", card.ID}}}}
	collection, err = s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return "", err
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}

	return card.ID, nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, entity string, id string) error {
	collection, err := s.db.GetCollection(ctx, "user_db", "users")
	if err != nil {
		return err
	}

	switch entity {
	case "customers":
		//TODO
	case "addresses":
		//TODO
	case "cards":
		//TODO
	default:
		return errors.New("Invalid entity " + entity)
	}

	query := bson.D{{Key: "UserID", Value: id}}
	return collection.DeleteOne(ctx, query)
}

func calculatePassHash(pass, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}

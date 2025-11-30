package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dbCard struct {
	Card `bson:",inline"`
	ID   primitive.ObjectID `bson:"_id"`
}

// Gets card by objects Id
func (s *UserServiceImpl) carddb_GetCard(ctx context.Context, cardid string) (Card, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "card")
	if err != nil {
		return Card{}, err
	}

	// Convert the card ID
	id, err := primitive.ObjectIDFromHex(cardid)
	if err != nil {
		return Card{}, errors.New("invalid ID Hex")
	}

	// Run the query
	cursor, err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return Card{}, err
	}
	card := dbCard{}
	_, err = cursor.One(ctx, &card)

	// Convert from DB card data to Card object
	card.Card.ID = card.ID.Hex()
	return card.Card, err
}

// Gets cards from the card store
func (s *UserServiceImpl) carddb_GetCards(ctx context.Context, cardIds []string) ([]Card, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "card")
	if err != nil {
		return nil, err
	}

	if len(cardIds) == 0 {
		return nil, nil
	}

	// Convert the card IDs from hex strings to objects
	ids, err := hexToObjectIds(cardIds)
	if err != nil {
		return nil, err
	}

	// Run the query
	cursor, err := collection.FindMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}})
	if err != nil {
		return nil, err
	}
	dbCards := make([]dbCard, 0, len(cardIds))
	err = cursor.All(ctx, &dbCards)

	// Convert from DB card data to Card objects
	cards := make([]Card, 0, len(dbCards))
	for _, card := range dbCards {
		card.Card.ID = card.ID.Hex()
		cards = append(cards, card.Card)
	}

	return cards, err
}

func (s *UserServiceImpl) carddb_GetAllCards(ctx context.Context) ([]Card, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "card")
	if err != nil {
		return nil, err
	}

	// Run the query
	cursor, err := collection.FindMany(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	dbCards := make([]dbCard, 0)
	err = cursor.All(ctx, &dbCards)

	// Convert from DB card data to Card objects
	cards := make([]Card, 0, len(dbCards))
	for _, card := range dbCards {
		card.Card.ID = card.ID.Hex()
		cards = append(cards, card.Card)
	}

	return cards, err
}

// Adds a card to the cards DB
func (s *UserServiceImpl) carddb_CreateCard(ctx context.Context, card *Card) (primitive.ObjectID, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "card")
	if err != nil {
		return primitive.ObjectID{}, err
	}

	// Create and insert to DB
	dbcard := dbCard{Card: *card, ID: primitive.NewObjectID()}
	if _, err := collection.UpsertID(ctx, dbcard.ID, dbcard); err != nil {
		return dbcard.ID, err
	}

	// Update the provided card
	dbcard.Card.ID = dbcard.ID.Hex()
	*card = dbcard.Card
	return dbcard.ID, nil
}

// Creates or updates the provided cards in the cardStore.
func (s *UserServiceImpl) carddb_CreateCards(ctx context.Context, cards []Card) ([]primitive.ObjectID, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "card")
	if err != nil {
		return nil, err
	}

	if len(cards) == 0 {
		return []primitive.ObjectID{}, nil
	}
	createdIds := make([]primitive.ObjectID, 0)
	for _, card := range cards {
		toInsert := dbCard{
			Card: card,
			ID:   primitive.NewObjectID(),
		}
		_, err := collection.UpsertID(ctx, toInsert.ID, toInsert)
		if err != nil {
			return createdIds, err
		}
		createdIds = append(createdIds, toInsert.ID)
	}

	return createdIds, nil
}

func (s *UserServiceImpl) carddb_RemoveCard(ctx context.Context, cardid string) error {
	// Convert the card ID
	id, err := primitive.ObjectIDFromHex(cardid)
	if err != nil {
		return errors.New("invalid ID Hex")
	}
	return s.carddb_RemoveCards(ctx, []primitive.ObjectID{id})
}

// Removes all specified cards from the DB
func (s *UserServiceImpl) carddb_RemoveCards(ctx context.Context, ids []primitive.ObjectID) error {
	collection, err := s.db.GetCollection(ctx, "user_db", "card")
	if err != nil {
		return nil
	}
	return collection.DeleteMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}})
}

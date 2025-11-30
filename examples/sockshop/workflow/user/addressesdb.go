package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type dbAddress struct {
	Address `bson:",inline"`
	ID      primitive.ObjectID `bson:"_id"`
}

// Gets an address by object Id
func (s *UserServiceImpl) addressdb_GetAddress(ctx context.Context, addressid string) (Address, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "address")
	if err != nil {
		return Address{}, err
	}

	// Convert the address ID
	id, err := primitive.ObjectIDFromHex(addressid)
	if err != nil {
		return Address{}, errors.New("invalid ID Hex")
	}

	// Run the query
	cursor, err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return Address{}, err
	}
	address := dbAddress{}
	_, err = cursor.One(ctx, &address)

	// Convert from DB address data to Address object
	address.Address.ID = address.ID.Hex()
	return address.Address, err
}

// Gets addresses from the address store
func (s *UserServiceImpl) addressdb_GetAddresses(ctx context.Context, addressIds []string) ([]Address, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "address")
	if err != nil {
		return nil, err
	}

	if len(addressIds) == 0 {
		return nil, nil
	}

	// Convert the address IDs from hex strings to objects
	ids, err := hexToObjectIds(addressIds)
	if err != nil {
		return nil, err
	}

	// Run the query
	cursor, err := collection.FindMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}})
	if err != nil {
		return nil, err
	}
	dbAddresses := make([]dbAddress, 0, len(addressIds))
	err = cursor.All(ctx, &dbAddresses)

	// Convert from DB address data to Address objects
	addresses := make([]Address, 0, len(dbAddresses))
	for _, address := range dbAddresses {
		address.Address.ID = address.ID.Hex()
		addresses = append(addresses, address.Address)
	}

	return addresses, err
}

func (s *UserServiceImpl) addressdb_GetAllAddresses(ctx context.Context) ([]Address, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "address")
	if err != nil {
		return nil, err
	}

	// Run the query
	cursor, err := collection.FindMany(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	dbAddresses := make([]dbAddress, 0)
	err = cursor.All(ctx, &dbAddresses)

	// Convert from DB address data to Address objects
	addresses := make([]Address, 0, len(dbAddresses))
	for _, address := range dbAddresses {
		address.Address.ID = address.ID.Hex()
		addresses = append(addresses, address.Address)
	}

	return addresses, err
}

// Adds an address to the address DB
func (s *UserServiceImpl) addressdb_CreateAddress(ctx context.Context, address *Address) (primitive.ObjectID, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "address")
	if err != nil {
		return primitive.ObjectID{}, err
	}

	// Create and insert to DB
	dbaddress := dbAddress{Address: *address, ID: primitive.NewObjectID()}
	if _, err := collection.UpsertID(ctx, dbaddress.ID, dbaddress); err != nil {
		return dbaddress.ID, err
	}

	// Update the provided address
	dbaddress.Address.ID = dbaddress.ID.Hex()
	*address = dbaddress.Address
	return dbaddress.ID, nil
}

// Creates or updates the provided addresses in the addressStore.
func (s *UserServiceImpl) addressdb_CreateAddresses(ctx context.Context, addresses []Address) ([]primitive.ObjectID, error) {
	collection, err := s.db.GetCollection(ctx, "user_db", "address")
	if err != nil {
		return nil, err
	}

	if len(addresses) == 0 {
		return []primitive.ObjectID{}, nil
	}
	createdIds := make([]primitive.ObjectID, 0)
	for _, address := range addresses {
		toInsert := dbAddress{
			Address: address,
			ID:      primitive.NewObjectID(),
		}
		_, err := collection.UpsertID(ctx, toInsert.ID, toInsert)
		if err != nil {
			return createdIds, err
		}
		createdIds = append(createdIds, toInsert.ID)
	}

	return createdIds, nil
}

func (s *UserServiceImpl) addressdb_RemoveAddress(ctx context.Context, addressid string) error {
	// Convert the address ID
	id, err := primitive.ObjectIDFromHex(addressid)
	if err != nil {
		return errors.New("invalid ID Hex")
	}
	return s.addressdb_RemoveAddresses(ctx, []primitive.ObjectID{id})
}

func (s *UserServiceImpl) addressdb_RemoveAddresses(ctx context.Context, ids []primitive.ObjectID) error {
	collection, err := s.db.GetCollection(ctx, "user_db", "address")
	if err != nil {
		return err
	}
	return collection.DeleteMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}})
}

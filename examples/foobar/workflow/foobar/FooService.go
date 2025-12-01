package foobar

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Foo struct {
	FooID string
	Text  string
}

type FooService interface {
	WriteFoo(ctx context.Context, id string, text string, barID string) (Foo, Bar, error)
	ReadFoo(ctx context.Context, id string) (Foo, error)
	//UpdateFoo(ctx context.Context, id string, timestamp string) error
}

type FooServiceImpl struct {
	database   backend.NoSQLDatabase
	barService BarService
}

func NewFooServiceImpl(ctx context.Context, database backend.NoSQLDatabase, barService BarService) (FooService, error) {
	d := &FooServiceImpl{database: database, barService: barService}
	return d, nil
}

func (s *FooServiceImpl) writeToDatabase(ctx context.Context, foo2 Foo) error {
	collection, err := s.database.GetCollection(ctx, "foo_db", "foo")
	if err != nil {
		return err
	}

	err = collection.InsertOne(ctx, foo2)
	if err != nil {
		return err
	}

	return nil
}

func (s *FooServiceImpl) invokeService(ctx context.Context, barID2 string, id2 string, text2 string) error {
	_, err := s.barService.WriteBar(ctx, barID2, text2, id2)
	if err != nil {
		return err
	}
	return nil
}

func (s *FooServiceImpl) WriteFoo(ctx context.Context, id string, text string, barID string) (Foo, Bar, error) {
	foo := Foo{
		FooID: id,
		Text:  text,
	}

	// write to database

	err := s.writeToDatabase(ctx, foo)
	if err != nil {
		return Foo{}, Bar{}, err
	}

	/* collection, err := s.database.GetCollection(ctx, "foo_db", "foo")
	if err != nil {
		return Foo{}, err
	}

	err = collection.InsertOne(ctx, foo)
	if err != nil {
		return Foo{}, err
	} */

	// invoke service

	//s.barService.WriteBar(ctx, barID, text, id)

	err = s.invokeService(ctx, barID, id, text)
	if err != nil {
		return Foo{}, Bar{}, err
	}

	/* var bar Bar
	go func() {
		bar, err = s.barService.WriteBar(ctx, barID, text, id)
	}() */

	return foo, Bar{}, nil
}

func (s *FooServiceImpl) ReadFoo(ctx context.Context, id string) (Foo, error) {
	var bar Foo

	collection, err := s.database.GetCollection(ctx, "foo_db", "foo")
	if err != nil {
		return Foo{}, err
	}

	/* query := bson.D{{Key: "FooID", Value: id}}
	query2 := &query
	fmt.Print(query2)
	query = append(query, bson.E{Key: "AT", Value: 1})
	cursor, err := collection.FindOne(ctx, *query2) */

	/* query_d := bson.D{
		{Key: "Username", Value: bson.D{
			{Key: "$in", Value: id},
		}},
	}

	query_d = append(query_d, bson.E{Key: "Age", Value: bson.D{{Key: "$gt", Value: 18}}}) */

	query := bson.D{{Key: "FooID", Value: id}}
	cursor, err := collection.FindOne(ctx, query)
	if err != nil {
		return Foo{}, err
	}

	res, err := cursor.One(ctx, &bar)
	if !res || err != nil {
		return Foo{}, err
	}

	return bar, nil
}

/* func (s *FooServiceImpl) UpdateFoo(ctx context.Context, id string, timestamp string) error {
	collection, err := s.database.GetCollection(ctx, "foo_db", "foo")
	if err != nil {
		return err
	}

	query := bson.D{{Key: "FooID", Value: id}}
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "Reviews", Value: bson.D{
				{Key: "$each", Value: bson.A{
					bson.D{
						{Key: "ID", Value: id},
						{Key: "Timestamp", Value: timestamp},
					},
				}},
				{Key: "$position", Value: 0},
			}},
		}},
	}
	_, err = collection.UpdateMany(ctx, query, update)
	return err
}
*/

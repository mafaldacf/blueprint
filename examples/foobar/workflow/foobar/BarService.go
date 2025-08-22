package foobar

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type Bar struct {
	BarID string
	Text  string
}

type BarService interface {
	WriteBar(ctx context.Context, id string, text string) (Bar, error)
	ReadBar(ctx context.Context, id string) (Bar, error)
}

type BarServiceImpl struct {
	barDb backend.NoSQLDatabase
}

func NewBarServiceImpl(ctx context.Context, barDb backend.NoSQLDatabase) (BarService, error) {
	d := &BarServiceImpl{barDb: barDb}
	return d, nil
}

func (s *BarServiceImpl) WriteBar(ctx context.Context, id string, text string) (Bar, error) {
	bar := Bar{
		BarID: id,
		Text:  text,
	}

	collection, err := s.barDb.GetCollection(ctx, "bar_db", "bar")
	if err != nil {
		return Bar{}, err
	}

	err = collection.InsertOne(ctx, bar)
	if err != nil {
		return Bar{}, err
	}

	return bar, nil
}

func (s *BarServiceImpl) ReadBar(ctx context.Context, id string) (Bar, error) {
	var bar Bar

	collection, err := s.barDb.GetCollection(ctx, "bar_db", "bar")
	if err != nil {
		return Bar{}, err
	}

	query := bson.D{{Key: "BarID", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Bar{}, err
	}

	res, err := result.One(ctx, &bar)
	if !res || err != nil {
		return Bar{}, err
	}

	return bar, nil
}

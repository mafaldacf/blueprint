package foobar

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type Bar struct {
	BarID string
	Text  string
}

type BarService interface {
	WriteBar(ctx context.Context, id string, text string) (Bar, error)
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

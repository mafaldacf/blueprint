package bar

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

var HelloBarVariable = "Hello Bar!"

const HELLO_BAR_CONST = "Hello Bar (CONST)!"

type Bar struct {
	ID   string
	Text string
	Flag bool
}

func (b *Bar) GetID() string {
	return b.ID
}

func (b *Bar) GetText() string {
	return b.Text
}

func (b *Bar) Test() int {
	ID0 := 0
	if b.Flag {
		ID1 := 1
		ID0 += ID1
	} else {
		ID2 := 2
		ID0 += ID2
	}
	return ID0
}

type BarService interface {
	Bar(ctx context.Context, text string) (Bar, error)
}

type BarServiceImpl struct {
	barDb backend.NoSQLDatabase
}

func NewBarServiceImpl(ctx context.Context, barDb backend.NoSQLDatabase) (BarService, error) {
	d := &BarServiceImpl{barDb: barDb}
	return d, nil
}

func (s *BarServiceImpl) Bar(ctx context.Context, text string) (Bar, error) {
	newText := text

	bar := Bar{
		ID:   "id",
		Text: newText,
	}

	collection, err := s.barDb.GetCollection(ctx, "bar", "bar")
	if err != nil {
		return Bar{}, err
	}

	err = collection.InsertOne(ctx, bar)
	if err != nil {
		return Bar{}, err
	}

	return bar, nil
}

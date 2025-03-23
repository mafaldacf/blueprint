package foo

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

var HelloFooVariable = "Hello Foo!"

const HELLO_FOO_CONST = "Hello Foo (CONST)!"

type Foo struct {
	ID   string
	Text string
}

func (f *Foo) GetID() string {
	return f.ID
}

func (f *Foo) GetText() string {
	return f.Text
}
type FooService interface {
	Foo(ctx context.Context, text string) (Foo, error)
}

type FooServiceImpl struct {
	fooDb backend.NoSQLDatabase
}

func NewFooServiceImpl(ctx context.Context, fooDb backend.NoSQLDatabase) (FooService, error) {
	d := &FooServiceImpl{fooDb: fooDb}
	return d, nil
}

func (s *FooServiceImpl) Foo(ctx context.Context, text string) (Foo, error) {
	foo := Foo{
		ID:   "id",
		Text: text,
	}

	collection, err := s.fooDb.GetCollection(ctx, "foo", "foo")
	if err != nil {
		return Foo{}, err
	}

	err = collection.InsertOne(ctx, foo)
	if err != nil {
		return Foo{}, err
	}

	return foo, nil
}

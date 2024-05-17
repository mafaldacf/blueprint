package foo

import (
	"context"
)

type FooService interface {
	Foo(ctx context.Context, text string) (string, error)
}

type FooServiceImpl struct {
	
}

func NewFooServiceImpl(ctx context.Context) (FooService, error) {
	d := &FooServiceImpl{}
	return d, nil
}

func (d *FooServiceImpl) Foo(ctx context.Context, text string) (string, error) {
	return "Foo: " + text, nil
}

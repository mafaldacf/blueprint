package foo

import (
	"context"
)

type FooStruct struct {
	Value string
}

var HelloFooVariable = "Hello Foo!"
const HELLO_FOO_CONST = "Hello Foo (CONST)!"

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

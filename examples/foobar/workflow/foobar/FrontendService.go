package foobar

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/examples/foobar/workflow/foobar/bar"
	"github.com/blueprint-uservices/blueprint/examples/foobar/workflow/foobar/foo"
)

var HelloWorldVariable = "Hello World!"

const HELLO_WORLD_CONST = "Hello World (CONST)!"

type FrontendService interface {
	Frontend(ctx context.Context) (string, error)
}

type FrontendServiceImpl struct {
	barService bar.BarService
	fooService foo.FooService
}

func NewFrontendServiceImpl(ctx context.Context, fooService foo.FooService, barService bar.BarService) (FrontendService, error) {
	d := &FrontendServiceImpl{fooService: fooService, barService: barService}
	return d, nil
}

func (d *FrontendServiceImpl) Frontend(ctx context.Context) (string, error) {
	foo, err1 := d.fooService.Foo(ctx, "Frontend")
	bar, err2 := d.barService.Bar(ctx, foo.Text)
	if err1 != nil {
		return "", err1
	}
	if err2 != nil {
		return "", err2
	}
	out := fmt.Sprintf("%s, %s", foo.Text, bar.Text)
	return out, nil
}

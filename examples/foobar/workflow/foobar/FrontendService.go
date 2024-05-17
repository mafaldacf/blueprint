package foobar

import (
	"context"
	"github.com/blueprint-uservices/blueprint/examples/foobar/workflow/foobar/bar"
	"github.com/blueprint-uservices/blueprint/examples/foobar/workflow/foobar/foo"
)

type FrontendService interface {
	Frontend(ctx context.Context) (string, error)
}

type FrontendServiceImpl struct {
	barService bar.BarService
	fooService foo.FooService
}

func NewFrontendServiceImpl(ctx context.Context, fooService foo.FooService, barService bar.BarService) (FrontendService, error) {
	d := &FrontendServiceImpl{ fooService: fooService, barService: barService }
	return d, nil
}

func (d *FrontendServiceImpl) Frontend(ctx context.Context) (string, error) {
	f, err1 := d.fooService.Foo(ctx, "Frontend")
	b, err2 := d.barService.Bar(ctx, "Frontend")
	if err1 != nil {
		return "", err1
	}
	if err2 != nil {
		return "", err2
	}
	return f + ", " + b, nil
}

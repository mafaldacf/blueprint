package foobar

import (
	"context"
	"fmt"
)

type Frontend interface {
	Frontend(ctx context.Context, id string, fooText string, barText string) (string, error)
}

type FrontendImpl struct {
	barService BarService
	fooService FooService
}

func NewFrontendImpl(ctx context.Context, fooService FooService, barService BarService) (Frontend, error) {
	d := &FrontendImpl{fooService: fooService, barService: barService}
	return d, nil
}

func (d *FrontendImpl) Frontend(ctx context.Context, id string, fooText string, barText string) (string, error) {
	foo, err1 := d.fooService.WriteFoo(ctx, id, fooText)
	bar, err2 := d.barService.WriteBar(ctx, id, barText)
	if err1 != nil {
		return "", err1
	}
	if err2 != nil {
		return "", err2
	}
	out := fmt.Sprintf("%s, %s", foo.Text, bar.Text)
	return out, nil
}

package foobar

import (
	"context"
	"fmt"
)

type Frontend interface {
	WriteFooBar(ctx context.Context, id string, fooText string, bars []Bar, barText string, key string, val string) (string, error)
	//ReadFooBar(ctx context.Context, id string) (Foo, Bar, error)
}

type FrontendImpl struct {
	barService BarService
	fooService FooService
}

func NewFrontendImpl(ctx context.Context, fooService FooService, barService BarService) (Frontend, error) {
	d := &FrontendImpl{fooService: fooService, barService: barService}
	return d, nil
}

func (d *FrontendImpl) WriteFooBar(ctx context.Context, id string, fooText string, bars []Bar, barText string, key string, val string) (string, error) {
	foo, err1 := d.fooService.WriteFoo(ctx, id, fooText, key, val, bars)
	/* bar, err2 := d.barService.WriteBar(ctx, id, barText) */
	if err1 != nil {
		return "", err1
	}
	/* if err2 != nil {
		return "", err2
	} */
	out := fmt.Sprintf("%s, %s", foo.Text, foo.BarItems[0].Text)
	return out, nil
}

/* func (d *FrontendImpl) ReadFooBar(ctx context.Context, id string)  (Foo, Bar, error) {
	foo, err1 := d.fooService.ReadFoo(ctx, id)
	bar, err2 := d.barService.ReadBar(ctx, foo.FooID)
	if err1 != nil {
		return Foo{}, Bar{}, err1
	}
	if err2 != nil {
		return Foo{}, Bar{}, err2
	}
	return foo, bar, nil
} */

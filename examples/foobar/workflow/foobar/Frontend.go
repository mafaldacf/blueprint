package foobar

import (
	"context"
)

type Frontend interface {
	WriteFoo(ctx context.Context, id string, text string, barID string) (Foo, Bar, error)
	ReadBarFoo(ctx context.Context, barID string) (Bar, Foo, error)
	//UpdateFoo(ctx context.Context, id string, timestamp string) error
}

type FrontendImpl struct {
	fooService FooService
	barService BarService
}

func NewFrontendImpl(ctx context.Context, fooService FooService, barService BarService) (Frontend, error) {
	f := &FrontendImpl{fooService: fooService, barService: barService}
	return f, nil
}

func (s *FrontendImpl) WriteFoo(ctx context.Context, id string, text string, barID string) (Foo, Bar, error) {
	return s.fooService.WriteFoo(ctx, id, text, barID)
}

func (s *FrontendImpl) ReadBarFoo(ctx context.Context, barID string) (Bar, Foo, error) {
	bar, err := s.barService.ReadBar(ctx, barID)
	if err != nil {
		return Bar{}, Foo{}, err
	}
	foo, err := s.fooService.ReadFoo(ctx, bar.FooID)
	if err != nil {
		return Bar{}, Foo{}, err
	}
	return bar, foo, nil
}

/* func (s *FrontendImpl) UpdateFoo(ctx context.Context, id string, timestamp string) error {
	return s.fooService.UpdateFoo(ctx, id, timestamp)
} */

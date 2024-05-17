package bar

import (
	"context"
)

type BarService interface {
	Bar(ctx context.Context, text string) (string, error)
}

type BarServiceImpl struct {
	
}

func NewBarServiceImpl(ctx context.Context) (BarService, error) {
	d := &BarServiceImpl{}
	return d, nil
}

func (d *BarServiceImpl) Bar(ctx context.Context, text string) (string, error) {
	return "Bar: " + text, nil
}

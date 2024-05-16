package dummy

import (
	"context"
)

type DummyService2 interface {
	Dummy(ctx context.Context, postID int64) error
}

type DummyServiceImpl struct {
	
}

func NewDummyServiceImpl(ctx context.Context) (DummyService2, error) {
	d := &DummyServiceImpl{}
	return d, nil
}

func (d *DummyServiceImpl) Dummy(ctx context.Context, postID int64) error {
	return nil
}

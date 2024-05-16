package postnotification

import (
	"context"
)

type DummyService interface {
	Dummy(ctx context.Context, postID int64) error
}

type DummyServiceImpl struct {
	
}

func NewDummyServiceImpl(ctx context.Context) (DummyService, error) {
	d := &DummyServiceImpl{}
	return d, nil
}

func (d *DummyServiceImpl) Dummy(ctx context.Context, postID int64) error {
	return nil
}

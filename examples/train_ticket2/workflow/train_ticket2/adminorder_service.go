package train_ticket2

import (
	"context"
)

type AdminOrderService interface {
	DeleteOrder(ctx context.Context, id string) error
}

type AdminOrderServiceImpl struct {
	orderService OrderService
}

func NewAdminOrderServiceImpl(ctx context.Context, orderService OrderService) (AdminOrderService, error) {
	return &AdminOrderServiceImpl{orderService: orderService}, nil
}

func (a *AdminOrderServiceImpl) DeleteOrder(ctx context.Context, id string) error {
	return a.orderService.Delete(ctx, id)
}

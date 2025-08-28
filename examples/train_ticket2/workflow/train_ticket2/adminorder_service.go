package train_ticket2

import (
	"context"
)

type AdminOrderService interface {
	AddOrder(ctx context.Context, order Order) (Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type AdminOrderServiceImpl struct {
	orderService OrderService
}

func NewAdminOrderServiceImpl(ctx context.Context, orderService OrderService) (AdminOrderService, error) {
	return &AdminOrderServiceImpl{orderService: orderService}, nil
}

func (a *AdminOrderServiceImpl) AddOrder(ctx context.Context, order Order) (Order, error) {
	return a.orderService.Create(ctx, order)
}

func (a *AdminOrderServiceImpl) DeleteOrder(ctx context.Context, id string) error {
	return a.orderService.Delete(ctx, id)
}

package trainticket

import (
	"context"
	"fmt"
)

type ExecuteService interface {
	TicketCollect(ctx context.Context, orderID string) error
	TicketExecute(ctx context.Context, orderID string) error
}

type ExecuteServiceImpl struct {
	orderService OrderService
}

func NewExecuteServiceImpl(ctx context.Context, orderService OrderService) (ExecuteService, error) {
	return &ExecuteServiceImpl{orderService: orderService}, nil
}

func (e *ExecuteServiceImpl) TicketCollect(ctx context.Context, orderID string) error {
	// 1. get order information
	order, err := e.orderService.GetOrderById(ctx, orderID)
	if err != nil {
		return err
	}
	// 2. check if the order came in
	if order.Status != ORDER_STATUS_PAID && order.Status != ORDER_STATUS_CHANGE {
		return fmt.Errorf("ticket not paid or changed")
	}
	// 3. confirm inbound, request change order info
	return e.orderService.ModifyOrder(ctx, orderID, ORDER_STATUS_COLLECTED)
}

func (e *ExecuteServiceImpl) TicketExecute(ctx context.Context, orderID string) error {
	// 1. get order information
	order, err := e.orderService.GetOrderById(ctx, orderID)
	if err != nil {
		return err
	}
	// 2. check if the order came in
	if order.Status != ORDER_STATUS_COLLECTED {
		return fmt.Errorf("ticket not collected")
	}
	// 3. confirm inbound, request change order info
	return e.orderService.ModifyOrder(ctx, orderID, ORDER_STATUS_USED)
}

package train_ticket2

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type CancelService interface {
	CalculateRefund(ctx context.Context, orderID string) (string, error)
	CancelOrder(ctx context.Context, orderID string, loginID string) error
}

type CancelServiceImpl struct {
	orderService         OrderService
	userService          UserService
	insidePaymentService InsidePaymentService
	emailQueue           backend.Queue
}

func NewCancelServiceImpl(ctx context.Context, orderService OrderService, userService UserService, insidePaymentService InsidePaymentService, emailQueue backend.Queue) (CancelService, error) {
	return &CancelServiceImpl{orderService: orderService, userService: userService, insidePaymentService: insidePaymentService, emailQueue: emailQueue}, nil
}

func (c *CancelServiceImpl) CalculateRefund(ctx context.Context, orderID string) (string, error) {
	order, err := c.orderService.GetOrderById(ctx, orderID)
	if err != nil {
		return "", err
	}

	if order.Status == ORDER_STATUS_NOT_PAID || order.Status == ORDER_STATUS_PAID {
		if order.Status == ORDER_STATUS_NOT_PAID {
			return "not paid", nil
		}
		return calculateRefund(order)
	}
	return "", fmt.Errorf("not permitted")
}

func (c *CancelServiceImpl) CancelOrder(ctx context.Context, orderID string, loginID string) error {
	order, err := c.orderService.GetOrderById(ctx, orderID)
	if err != nil {
		return err
	}

	if order.Status == ORDER_STATUS_NOT_PAID || order.Status == ORDER_STATUS_PAID || order.Status == ORDER_STATUS_CHANGE {
		order.Status = ORDER_STATUS_CANCELED
		_, err := c.orderService.SaveOrderInfo(ctx, order)
		if err != nil {
			return err
		}

		money, err := calculateRefund(order)
		if err != nil {
			return err
		}

		err = c.insidePaymentService.Drawback(ctx, loginID, money)
		if err != nil {
			return err
		}

		user, err := c.userService.FindByUserID(ctx, order.AccountID)
		if err != nil {
			return err
		}

		notifyInfo := NotifyInfo{
			Date:        dateToString(),
			Email:       user.Email,
			StartPlace:  order.FromStation,
			EndPlace:    order.ToStation,
			Username:    user.Username,
			SeatNumber:  order.SeatNumber,
			OrderNumber: order.ID,
			Price:       order.Price,
			SeatClass:   strconv.Itoa(order.SeatClass),
			StartTime:   order.TravelTime,
		}
		fmt.Printf("[CANCEL] notify info: %v\n", notifyInfo)

		/* _, err = c.emailQueue.Push(ctx, notifyInfo)
		if err != nil {
			return err
		} */
	}

	return nil
}

func calculateRefund(order Order) (string, error) {
	if order.Status == ORDER_STATUS_NOT_PAID {
		return "0.00", nil
	}

	datePart, err := time.ParseInLocation(CALENDAR_LAYOUT, order.TravelDate, time.Local)
	if err != nil {
		return "", err
	}
	timePart, err := time.ParseInLocation(CALENDAR_LAYOUT, order.TravelTime, time.Local)
	if err != nil {
		return "", err
	}

	year, month, day := datePart.Date()
	hour, minute, second := timePart.Clock()

	startTime := time.Date(year, month, day, hour, minute, second, 0, time.Local)
	now := time.Now()

	if now.After(startTime) {
		return "0", nil
	}

	// calculate 80% refund
	totalPrice, err := strconv.ParseFloat(order.Price, 64)
	if err != nil {
		return "", err
	}
	price := totalPrice * 0.8
	refund := fmt.Sprintf("%.2f", price)

	return refund, nil
}

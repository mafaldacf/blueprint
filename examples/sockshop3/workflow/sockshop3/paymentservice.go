// Package payment implements the SockShop payment microservice.
//
// The service fakes payments, implementing simple logic whereby payments
// are authorized when they're below a predefined threshold, and rejected
// when they are above that threshold.
package sockshop3

import (
	"context"
	"errors"
	"fmt"
)
type PaymentService interface {
	Authorise(ctx context.Context, amount float32) (Authorisation, error)
}

func NewPaymentServiceImpl(ctx context.Context, declineOverAmount string) (PaymentService, error) {
	return &PaymentServiceImpl{
		declineOverAmount: float32(50),
	}, nil
}

type PaymentServiceImpl struct {
	declineOverAmount float32
}

var ErrInvalidPaymentAmount = errors.New("invalid payment amount")

func (s *PaymentServiceImpl) Authorise(ctx context.Context, amount float32) (Authorisation, error) {
	if amount == 0 {
		return Authorisation{}, ErrInvalidPaymentAmount
	}
	if amount < 0 {
		return Authorisation{}, ErrInvalidPaymentAmount
	}
	if amount <= s.declineOverAmount {
		return Authorisation{
			Authorised: true,
			Message:    "Payment authorised",
		}, nil
	}
	return Authorisation{
		Authorised: false,
		Message:    fmt.Sprintf("Payment declined: amount exceeds %.2f", s.declineOverAmount),
	}, nil
}

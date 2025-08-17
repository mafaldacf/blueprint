// Package payment implements the SockShop payment microservice.
//
// The service fakes payments, implementing simple logic whereby payments
// are authorized when they're below a predefined threshold, and rejected
// when they are above that threshold.
package sockshop3

import (
	"context"
)

// PaymentService provides payment services
type PaymentService interface {
	Authorise(ctx context.Context, amount float32) (Authorisation, error)
}

// Returns a payment service where any transaction above the preconfigured
// threshold will return an invalid payment amount
func NewPaymentService(ctx context.Context, declineOverAmount string) (PaymentService, error) {
	return &paymentImpl{
		declineOverAmount: float32(50),
	}, nil
}

type paymentImpl struct {
	declineOverAmount float32
}

func (s *paymentImpl) Authorise(ctx context.Context, amount float32) (Authorisation, error) {
	authorization := Authorisation{
		Authorised: false,
		Message:    "Payment declined: amount exceeds",
	}
	return authorization, nil
}

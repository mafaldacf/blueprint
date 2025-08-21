package digota

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"

	//"github.com/blueprint-uservices/blueprint/examples/digota/workflow/digota/validation"
)

type PaymentService interface {
	NewCharge(ctx context.Context, currency int32, total uint64, card *Card, email string, statement string, paymentProviderId int32, metadata map[string]string) (*Charge, error)
	Get(ctx context.Context, id string) (*Charge, error)
	List(ctx context.Context, page int64, limit int64, sort int32) (*ChargeList, error)
	RefundCharge(ctx context.Context, id string, amount uint64, reason int32) (*Charge, error)
}

type PaymentServiceImpl struct {
	db backend.NoSQLDatabase
}

func NewPaymentServiceImpl(ctx context.Context, db backend.NoSQLDatabase) (PaymentService, error) {
	s := &PaymentServiceImpl{db: db}
	return s, nil
}

func (s *PaymentServiceImpl) NewCharge(ctx context.Context, currency int32, total uint64, card *Card, email string, statement string, paymentProviderId int32, metadata map[string]string) (*Charge, error) {
	charge := &Charge{
		Currency:     currency,
		ChargeAmount: total,
		Email:        email,
		Statement:    statement,
	}

	/* err := validation.Validate(charge)
	if err != nil {
		return nil, err
	} */

	collection, err := s.db.GetCollection(ctx, "payments_db", "payments")
	if err != nil {
		return nil, err
	}
	err = collection.InsertOne(ctx, *charge)
	return charge, err
}

func (s *PaymentServiceImpl) Get(ctx context.Context, id string) (*Charge, error) {
	collection, err := s.db.GetCollection(ctx, "payments_db", "payments")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "Id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	var charge *Charge
	found, err := result.One(ctx, charge)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("charge not found for id (%s)", id)
	}

	return charge, nil
}

func (s *PaymentServiceImpl) List(ctx context.Context, page int64, limit int64, sort int32) (*ChargeList, error) {
	collection, err := s.db.GetCollection(ctx, "payments_db", "payments")
	if err != nil {
		return nil, err
	}

	result, err := collection.FindMany(ctx, nil)
	if err != nil {
		return nil, err
	}

	var charges []*Charge
	err = result.All(ctx, charges)
	if err != nil {
		return nil, err
	}

	chargeList := &ChargeList{
		Charges: charges,
		Total:   int32(len(charges)),
	}

	return chargeList, nil
}

func (s *PaymentServiceImpl) RefundCharge(ctx context.Context, id string, amount uint64, reason int32) (*Charge, error) {
	collection, err := s.db.GetCollection(ctx, "payments_db", "payments")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "Id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	var charge *Charge
	found, err := result.One(ctx, charge)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("charge not found for id (%s)", id)
	}

	return charge, nil
}

package digota

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
	return &PaymentServiceImpl{db: db}, nil
}

func (s *PaymentServiceImpl) NewCharge(ctx context.Context, currency int32, total uint64, card *Card, email string, statement string, paymentProviderId int32, metadata map[string]string) (*Charge, error) {
	charge := &Charge{
		Id:           uuid.NewString(),
		Currency:     currency,
		ChargeAmount: total,
		Email:        email,
		Statement:    statement,
	}

	collection, err := s.db.GetCollection(ctx, "payments_db", "payments")
	if err != nil {
		return nil, err
	}
	return charge, collection.InsertOne(ctx, *charge)
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

	var charge Charge
	found, err := result.One(ctx, &charge)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("charge not found for id (%s)", id)
	}

	return &charge, nil
}

func (s *PaymentServiceImpl) List(ctx context.Context, page int64, limit int64, sort int32) (*ChargeList, error) {
	collection, err := s.db.GetCollection(ctx, "payments_db", "payments")
	if err != nil {
		return nil, err
	}

	result, err := collection.FindMany(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var charges []Charge
	err = result.All(ctx, &charges)
	if err != nil {
		return nil, err
	}

	return &ChargeList{Charges: charges, Total: int32(len(charges))}, nil
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

	var charge Charge
	found, err := result.One(ctx, &charge)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("charge not found for id (%s)", id)
	}

	return &charge, nil
}

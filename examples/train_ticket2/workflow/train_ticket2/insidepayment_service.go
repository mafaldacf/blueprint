package train_ticket2

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type InsidePaymentService interface {
	// TODO:
	// - Pay
	// - CreateAccount
	// - AddMoney
	// - QueryPayment
	// - QueryAccount
	// - DrawBack
	// - PayDifference
	// - QueryAddMoney
	Drawback(ctx context.Context, userID string, money string) error
}

type InsidePaymentServiceImpl struct {
	insidePaymentDB backend.NoSQLDatabase
}

func NewInsidePaymentServiceImpl(ctx context.Context, orderDB backend.NoSQLDatabase) (InsidePaymentService, error) {
	return &InsidePaymentServiceImpl{insidePaymentDB: orderDB}, nil
}

func (o *InsidePaymentServiceImpl) Drawback(ctx context.Context, userID string, money string) error {
	collection, err := o.insidePaymentDB.GetCollection(ctx, "inside_payment_db", "inside_money")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "UserID", Value: userID}}
	update := bson.D{{Key: "Money", Value: money}, {Key: "Type", Value: INSIDE_MONEY_TYPE_DRAWBACK}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

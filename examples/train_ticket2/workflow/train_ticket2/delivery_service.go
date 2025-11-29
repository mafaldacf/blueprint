// Package delivery implements ts-delivery service from the original train ticket application
package train_ticket2

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slog"
)

// DeliveryService implements the Delivery microservice.
//
// It is not a service that can be called; instead it pulls deliveries from
// the delivery queue
type DeliveryService interface {
	Run(ctx context.Context) error
	FindDelivery(ctx context.Context, orderID string) (Delivery, error)
}

type DeliveryServiceImpl struct {
	db   backend.NoSQLDatabase
	delQ backend.Queue
}

func NewDeliveryServiceImpl(ctx context.Context, queue backend.Queue, db backend.NoSQLDatabase) (DeliveryService, error) {
	return &DeliveryServiceImpl{db: db, delQ: queue}, nil
}

func (d *DeliveryServiceImpl) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			var delivery *Delivery
			ok, err := d.delQ.Pop(ctx, delivery)
			if err != nil {
				slog.Error(fmt.Sprintf("DeliveryService unable to pull delivery info from deliver queue due to %v", err))
			}
			if ok {
				coll, err := d.db.GetCollection(ctx, "delivery_db", "delivery")
				if err != nil {
					slog.Error(fmt.Sprintf("DeliveryService unable to obtain a collection to delivery database due to %v", err))
				}
				if delivery.getId() == "" {
					delivery.setId(uuid.NewString())
				}
				err = coll.InsertOne(ctx, delivery)
				if err != nil {
					slog.Error(fmt.Sprintf("DeliveryService unable to add a delivery to the database due to %v", err))
				}
			}
		}
	}
}

func (d *DeliveryServiceImpl) FindDelivery(ctx context.Context, orderID string) (Delivery, error) {
	collection, err := d.db.GetCollection(ctx, "delivery_db", "delivery")
	if err != nil {
		return Delivery{}, err
	}

	filter := bson.D{{Key: "OrderID", Value: orderID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Delivery{}, err
	}

	var delivery Delivery
	ok, err := cursor.One(ctx, &delivery)
	if err != nil {
		return Delivery{}, err
	}
	if !ok {
		return Delivery{}, fmt.Errorf("delivery (%s) not found", orderID)
	}
	return delivery, nil
}

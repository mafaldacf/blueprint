// Package shipping implements the SockShop shipping microservice.
//
// All the shipping microservice does is push the shipment to a queue.
// The queue-master service pulls shipments from the queue and "processes"
// them.
package sockshop3

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// ShippingService implements the SockShop shipping microservice
type ShippingService interface {
	PostShipping(ctx context.Context, shipment Shipment) (Shipment, error)
	GetShipment(ctx context.Context, id string) (Shipment, error)
	UpdateStatus(ctx context.Context, id, status string) error
}

func NewShippingServiceImpl(ctx context.Context, queue backend.Queue, db backend.NoSQLDatabase) (ShippingService, error) {
	return &ShippingServiceImpl{
		q:  queue,
		db: db,
	}, nil
}

type ShippingServiceImpl struct {
	q  backend.Queue
	db backend.NoSQLDatabase
}

func (service *ShippingServiceImpl) PostShipping(ctx context.Context, shipment Shipment) (Shipment, error) {
	// Push to the queue to be shipped
	shipped, err := service.q.Push(ctx, shipment)
	if err != nil {
		return shipment, err
	} else if !shipped {
		return shipment, errors.Errorf("Unable to submit shipment %v %v to the shipping queue", shipment.ID, shipment.Name)
	}

	collection, err := service.db.GetCollection(ctx, "ship_db", "shipments")
	if err != nil {
		return Shipment{}, err
	}
	return shipment, collection.InsertOne(ctx, shipment)
}

// GetShipment implements ShippingService.
func (s *ShippingServiceImpl) GetShipment(ctx context.Context, id string) (Shipment, error) {
	collection, err := s.db.GetCollection(ctx, "ship_db", "shipments")
	if err != nil {
		return Shipment{}, err
	}

	cursor, err := collection.FindOne(ctx, bson.D{{Key: "ID", Value: id}})
	if err != nil {
		return Shipment{}, err
	}

	var shipment Shipment
	shipmentExists, err := cursor.One(ctx, &shipment)
	if err != nil {
		return Shipment{}, err
	} else if !shipmentExists {
		return Shipment{}, errors.Errorf("unknown shipment %v", id)
	}
	return shipment, nil
}

// UpdateStatus implements ShippingService.
func (s *ShippingServiceImpl) UpdateStatus(ctx context.Context, id string, status string) error {
	collection, err := s.db.GetCollection(ctx, "ship_db", "shipments")
	if err != nil {
		return err
	}

	updated, err := collection.UpdateOne(ctx, bson.D{{Key: "ID", Value: id}}, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: status}}}})
	if err != nil {
		return err
	} else if updated == 0 {
		return errors.Errorf("unknown shipment %v", id)
	}
	
	return nil
}

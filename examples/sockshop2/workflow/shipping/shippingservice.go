// Package shipping implements the SockShop shipping microservice.
//
// All the shipping microservice does is push the shipment to a queue.
// The queue-master service pulls shipments from the queue and "processes"
// them.
package shipping

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

// ShippingService implements the SockShop shipping microservice
type ShippingService interface {
	// Submit a shipment to be shipped.  The actual handling of the
	// shipment will happen asynchronously by the queue-master service.
	//
	// Returns the submitted shipment or an error
	PostShipping(ctx context.Context, shipment Shipment) (Shipment, error)

	// Get a shipment's status
	GetShipment(ctx context.Context, id string) (Shipment, error)

	// Update a shipment's status; called by the queue master
	UpdateStatus(ctx context.Context, id, status string) error
}

// Instantiates a shipping service that submits all shipments to a queue for asynchronous background processing
func NewShippingService(ctx context.Context, queue backend.Queue, db backend.NoSQLDatabase) (ShippingService, error) {
	return &shippingImpl{
		q:  queue,
		db: db,
	}, nil
}

type shippingImpl struct {
	q  backend.Queue
	db backend.NoSQLDatabase
}

// PostShipping implements ShippingService.
func (service *shippingImpl) PostShipping(ctx context.Context, shipment Shipment) (Shipment, error) {
	collection, _ := service.db.GetCollection(ctx, "shipping_service", "shipments")
	service.q.Push(ctx, shipment)
	return shipment, collection.InsertOne(ctx, shipment)
}

// GetShipment implements ShippingService.
func (s *shippingImpl) GetShipment(ctx context.Context, id string) (Shipment, error) {
	collection, _ := s.db.GetCollection(ctx, "shipping_service", "shipments")
	cursor, err := collection.FindOne(ctx, bson.D{{"id", id}})
	if err != nil {
		return Shipment{}, err
	}
	var shipment Shipment
	cursor.One(ctx, &shipment)
	return shipment, nil
}

// UpdateStatus implements ShippingService.
func (s *shippingImpl) UpdateStatus(ctx context.Context, id string, status string) error {
	collection, _ := s.db.GetCollection(ctx, "shipping_service", "shipments")
	collection.UpdateOne(ctx, bson.D{{"id", id}}, bson.D{{"$set", bson.D{{"status", status}}}})
	return nil
}

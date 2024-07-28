// Package queuemaster implements the queue-master SockShop service, responsible for
// pulling and "processing" shipments from the shipment queue.
package queuemaster

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"

	"github.com/blueprint-uservices/blueprint/examples/sockshop2/workflow/shipping"
)

// QueueMaster implements the SockShop queue-master microservice.
//
// It is not a service that can be called; instead it pulls shipments from
// the shipments queue
type QueueMaster interface {
	// Runs the background goroutine that continually pulls elements from
	// the queue.  Does not return until ctx is cancelled or an error is
	// encountered
	Run(ctx context.Context) error
}

// Creates a new QueueMaster service.
//
// New: once an order is shipped, it will update the order status in the orderservice.
func NewQueueMaster(ctx context.Context, queue backend.Queue, shipping shipping.ShippingService) (QueueMaster, error) {
	return &queueMasterImpl{
		q:           queue,
		shipping:    shipping,
		exitOnError: false,
		processed:   0,
	}, nil
}

type queueMasterImpl struct {
	q           backend.Queue
	shipping    shipping.ShippingService
	exitOnError bool
	processed   int32
}

// Starts a processing loop that continually pulls elements from the queue.
// Does not exit when an error is encountered; only when ctx is cancelled
func (q *queueMasterImpl) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			var shipment shipping.Shipment
			q.q.Pop(ctx, &shipment)
			q.shipping.UpdateStatus(ctx, shipment.ID, "shipped")
		}
	}
}

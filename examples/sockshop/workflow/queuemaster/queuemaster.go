// Package queuemaster implements the queue-master SockShop service, responsible for
// pulling and "processing" shipments from the shipment queue.
package queuemaster

import (
	"context"
	"fmt"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"

	"github.com/blueprint-uservices/blueprint/examples/sockshop/workflow/shipping"
)

type QueueMaster interface {
	Run(ctx context.Context) error
}

func NewQueueMasterImpl(ctx context.Context, queue backend.Queue, shipping shipping.ShippingService) (QueueMaster, error) {
	return &QueueMasterImpl{
		q:           queue,
		shipping:    shipping,
		exitOnError: false,
		processed:   0,
	}, nil
}

type QueueMasterImpl struct {
	q           backend.Queue
	shipping    shipping.ShippingService
	exitOnError bool
	processed   int32
}

// Starts a processing loop that continually pulls elements from the queue.
// Does not exit when an error is encountered; only when ctx is cancelled
func (q *QueueMasterImpl) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			var shipment shipping.Shipment
			didPop, err := q.q.Pop(ctx, &shipment)
			if err != nil {
				if q.exitOnError {
					return err
				} else {
					slog.Error(fmt.Sprintf("QueueMaster unable to pull order from shipping queue due to %v", err))
					continue
				}
			}
			if didPop {
				msgNumber := atomic.AddInt32(&q.processed, 1)
				slog.Info(fmt.Sprintf("Received shipment task %v %v: %v", msgNumber, shipment.ID, shipment.Name))

				// Keep attempting to update shipping status
				for {
					err := q.shipping.UpdateStatus(ctx, shipment.ID, "shipped")
					if err != nil {
						if q.exitOnError {
							return err
						} else {
							slog.Error(fmt.Sprintf("Unable to send shipment %v due to %v; waiting 1 second then retrying", shipment.ID, err))
							time.Sleep(1 * time.Second)
						}
					} else {
						break
					}
				}
			}
		}
	}
}

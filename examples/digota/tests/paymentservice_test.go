package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/digota/workflow/digota"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var paymentServiceRegistry = registry.NewServiceRegistry[digota.PaymentService]("payment_service")

func init() {
	paymentServiceRegistry.Register("local", func(ctx context.Context) (digota.PaymentService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}

		return digota.NewPaymentServiceImpl(ctx, db)
	})
}

func TestPaymentServiceNewCharge(t *testing.T) {
	ctx := context.Background()
	paymentService, err := paymentServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	charge, err := paymentService.NewCharge(ctx, 1, 1000, &digota.Card{Number: "4242424242424242", ExpireMonth: "12", ExpireYear: "2025", CVC: "123"}, "test@example.com", "Test charge", 1, map[string]string{"key": "value"})
	assert.NoError(t, err)
	assert.NotNil(t, charge)
	assert.Equal(t, uint64(1000), charge.ChargeAmount)
	assert.Equal(t, "test@example.com", charge.Email)
}

func TestPaymentServiceGet(t *testing.T) {
	ctx := context.Background()
	paymentService, err := paymentServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a charge
	charge, err := paymentService.NewCharge(ctx, 1, 500, &digota.Card{Number: "4242424242424242", ExpireMonth: "12", ExpireYear: "2025", CVC: "123"}, "get@example.com", "Get test", 1, map[string]string{})
	assert.NoError(t, err)

	// Get it
	retrieved, err := paymentService.Get(ctx, charge.Id)
	assert.NoError(t, err)
	assert.Equal(t, charge.Id, retrieved.Id)
	assert.Equal(t, uint64(500), retrieved.ChargeAmount)
}

func TestPaymentServiceList(t *testing.T) {
	ctx := context.Background()
	paymentService, err := paymentServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create some charges
	_, err = paymentService.NewCharge(ctx, 1, 100, &digota.Card{Number: "4242424242424242", ExpireMonth: "12", ExpireYear: "2025", CVC: "123"}, "list1@example.com", "List 1", 1, map[string]string{})
	assert.NoError(t, err)
	_, err = paymentService.NewCharge(ctx, 1, 200, &digota.Card{Number: "4242424242424242", ExpireMonth: "12", ExpireYear: "2025", CVC: "123"}, "list2@example.com", "List 2", 1, map[string]string{})
	assert.NoError(t, err)

	list, err := paymentService.List(ctx, 0, 10, 1)
	assert.NoError(t, err)
	assert.True(t, list.Total >= 2)
}

func TestPaymentServiceRefundCharge(t *testing.T) {
	ctx := context.Background()
	paymentService, err := paymentServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a charge
	charge, err := paymentService.NewCharge(ctx, 1, 1000, &digota.Card{Number: "4242424242424242", ExpireMonth: "12", ExpireYear: "2025", CVC: "123"}, "refund@example.com", "Refund test", 1, map[string]string{})
	assert.NoError(t, err)

	// Refund it
	refunded, err := paymentService.RefundCharge(ctx, charge.Id, 500, 1)
	assert.NoError(t, err)
	assert.Equal(t, charge.Id, refunded.Id)
	// Assuming refund updates the charge
}

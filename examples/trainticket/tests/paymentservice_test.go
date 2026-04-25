package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var paymentServiceRegistry = registry.NewServiceRegistry[trainticket.PaymentService]("payments_service")

func init() {
	paymentServiceRegistry.Register("local", func(ctx context.Context) (trainticket.PaymentService, error) {
		payDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		moneyDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewPaymentServiceImpl(ctx, payDB, moneyDB)
	})
}

func TestPaymentServiceInitPayment(t *testing.T) {
	ctx := context.Background()
	service, err := paymentServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	payment := trainticket.Payment{
		ID:          "pay001",
		OrderID:     "order001",
		UserID:      "user001",
		Price:       "120.50",
		PaymentType: trainticket.PaymentType_P,
	}
	err = service.InitPayment(ctx, payment)
	assert.NoError(t, err)
}

func TestPaymentServicePay(t *testing.T) {
	ctx := context.Background()
	service, err := paymentServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	payment := trainticket.Payment{
		ID:          "pay002",
		OrderID:     "order002",
		UserID:      "user002",
		Price:       "85.00",
		PaymentType: trainticket.PaymentType_P,
	}
	err = service.InitPayment(ctx, payment)
	assert.NoError(t, err)

	err = service.Pay(ctx, payment)
	assert.NoError(t, err)
}

func TestPaymentServiceQuery(t *testing.T) {
	ctx := context.Background()
	service, err := paymentServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	payment := trainticket.Payment{
		ID:          "pay003",
		OrderID:     "order003",
		UserID:      "user003",
		Price:       "200.00",
		PaymentType: trainticket.PaymentType_D,
	}
	err = service.InitPayment(ctx, payment)
	assert.NoError(t, err)

	payments, err := service.Query(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, payments)
}

func TestPaymentServiceAddMoney(t *testing.T) {
	ctx := context.Background()
	service, err := paymentServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	moneyPayment := trainticket.Payment{
		ID:          "pay004",
		UserID:      "user004",
		Price:       "500.00",
		PaymentType: trainticket.PaymentType_O,
	}
	err = service.AddMoney(ctx, moneyPayment)
	assert.NoError(t, err)
}

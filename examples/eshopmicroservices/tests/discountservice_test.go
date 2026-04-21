package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/discount"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var discountServiceRegistry = registry.NewServiceRegistry[discount.DiscountService]("discount_service")

func init() {
	discountServiceRegistry.Register("local", func(ctx context.Context) (discount.DiscountService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return discount.NewDiscountServiceImpl(ctx, db)
	})
}

func TestDiscountServiceCreate(t *testing.T) {
	ctx := context.Background()
	discountService, err := discountServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	coupon, err := discountService.CreateDiscount(ctx, discount.CreateDiscountRequest{
		Coupon: discount.Coupon{Id: 1, ProductName: "Widget", Description: "10% off", Amount: 5.0},
	})
	assert.NoError(t, err)
	assert.Equal(t, "Widget", coupon.ProductName)
	assert.Equal(t, 5.0, coupon.Amount)
}

func TestDiscountServiceGet(t *testing.T) {
	ctx := context.Background()
	discountService, err := discountServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = discountService.CreateDiscount(ctx, discount.CreateDiscountRequest{
		Coupon: discount.Coupon{Id: 2, ProductName: "Gadget", Description: "5% off", Amount: 2.5},
	})
	assert.NoError(t, err)

	coupon, err := discountService.GetDiscount(ctx, discount.GetDiscountRequest{ProductName: "Gadget"})
	assert.NoError(t, err)
	assert.Equal(t, "Gadget", coupon.ProductName)
	assert.Equal(t, 2.5, coupon.Amount)
}

func TestDiscountServiceUpdate(t *testing.T) {
	ctx := context.Background()
	discountService, err := discountServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = discountService.CreateDiscount(ctx, discount.CreateDiscountRequest{
		Coupon: discount.Coupon{Id: 3, ProductName: "Gizmo", Description: "Original", Amount: 3.0},
	})
	assert.NoError(t, err)

	updated, err := discountService.UpdateDiscount(ctx, discount.UpdateDiscountRequest{
		Coupon: discount.Coupon{Id: 3, ProductName: "Gizmo", Description: "Updated", Amount: 6.0},
	})
	assert.NoError(t, err)
	assert.Equal(t, "Updated", updated.Description)
	assert.Equal(t, 6.0, updated.Amount)
}

func TestDiscountServiceDelete(t *testing.T) {
	ctx := context.Background()
	discountService, err := discountServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = discountService.CreateDiscount(ctx, discount.CreateDiscountRequest{
		Coupon: discount.Coupon{Id: 4, ProductName: "Doomed", Description: "Will be deleted", Amount: 1.0},
	})
	assert.NoError(t, err)

	resp, err := discountService.DeleteDiscount(ctx, discount.DeleteDiscountRequest{ProductName: "Doomed"})
	assert.NoError(t, err)
	assert.True(t, resp.Success)

	_, err = discountService.GetDiscount(ctx, discount.GetDiscountRequest{ProductName: "Doomed"})
	assert.Error(t, err)
}

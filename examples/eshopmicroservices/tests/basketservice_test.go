package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/basket"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/discount"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplequeue"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var basketServiceRegistry = registry.NewServiceRegistry[basket.BasketService]("basket_service")

func init() {
	basketServiceRegistry.Register("local", func(ctx context.Context) (basket.BasketService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		queue, err := simplequeue.NewSimpleQueue(ctx)
		if err != nil {
			return nil, err
		}
		discountSvc, err := discountServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		return basket.NewBasketServiceImpl(ctx, db, queue, discountSvc)
	})
}

func TestBasketServiceStoreBasket(t *testing.T) {
	ctx := context.Background()
	basketService, err := basketServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	discountSvc, err := discountServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	_, err = discountSvc.CreateDiscount(ctx, discount.CreateDiscountRequest{
		Coupon: discount.Coupon{Id: 100, ProductName: "Laptop", Amount: 50.0},
	})
	assert.NoError(t, err)

	cart := basket.ShoppingCart{
		UserName: "storeuser",
		Items: []basket.ShoppingCartItem{
			{Quantity: 1, Color: "silver", Price: 999.99, ProductId: uuid.NewString(), ProductName: "Laptop"},
		},
		TotalPrice: 999.99,
	}
	resp, err := basketService.StoreBasket(ctx, basket.StoreBasketRequest{Cart: cart})
	assert.NoError(t, err)
	assert.Equal(t, "storeuser", resp.UserName)
}

func TestBasketServiceGetBasket(t *testing.T) {
	ctx := context.Background()
	basketService, err := basketServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	discountSvc, err := discountServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	_, err = discountSvc.CreateDiscount(ctx, discount.CreateDiscountRequest{
		Coupon: discount.Coupon{Id: 101, ProductName: "Phone", Amount: 20.0},
	})
	assert.NoError(t, err)

	cart := basket.ShoppingCart{
		UserName: "getuser",
		Items: []basket.ShoppingCartItem{
			{Quantity: 2, Color: "black", Price: 499.99, ProductId: uuid.NewString(), ProductName: "Phone"},
		},
		TotalPrice: 999.98,
	}
	_, err = basketService.StoreBasket(ctx, basket.StoreBasketRequest{Cart: cart})
	assert.NoError(t, err)

	result, err := basketService.GetBasket(ctx, basket.GetBasketQuery{UserName: "getuser"})
	assert.NoError(t, err)
	assert.Equal(t, "getuser", result.Cart.UserName)
}

func TestBasketServiceDeleteBasket(t *testing.T) {
	ctx := context.Background()
	basketService, err := basketServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	cart := basket.ShoppingCart{
		UserName:   "deleteuser",
		Items:      []basket.ShoppingCartItem{},
		TotalPrice: 0,
	}
	_, err = basketService.StoreBasket(ctx, basket.StoreBasketRequest{Cart: cart})
	assert.NoError(t, err)

	result, err := basketService.DeleteBasket(ctx, basket.DeleteBasketCommand{UserName: "deleteuser"})
	assert.NoError(t, err)
	assert.True(t, result.IsSuccess)

	_, err = basketService.GetBasket(ctx, basket.GetBasketQuery{UserName: "deleteuser"})
	assert.Error(t, err)
}

func TestBasketServiceCheckoutBasket(t *testing.T) {
	ctx := context.Background()
	basketService, err := basketServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	discountSvc, err := discountServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	_, err = discountSvc.CreateDiscount(ctx, discount.CreateDiscountRequest{
		Coupon: discount.Coupon{Id: 102, ProductName: "Tablet", Amount: 30.0},
	})
	assert.NoError(t, err)

	cart := basket.ShoppingCart{
		UserName: "checkoutuser",
		Items: []basket.ShoppingCartItem{
			{Quantity: 1, Color: "white", Price: 299.99, ProductId: uuid.NewString(), ProductName: "Tablet"},
		},
		TotalPrice: 299.99,
	}
	_, err = basketService.StoreBasket(ctx, basket.StoreBasketRequest{Cart: cart})
	assert.NoError(t, err)

	resp, err := basketService.CheckoutBasket(ctx, basket.CheckoutBasketCommand{
		BasketCheckoutDto: basket.BasketCheckoutDto{
			UserName:      "checkoutuser",
			CustomerId:    uuid.NewString(),
			FirstName:     "John",
			LastName:      "Doe",
			EmailAddress:  "john@example.com",
			AddressLine:   "123 Main St",
			Country:       "US",
			State:         "NY",
			ZipCode:       "10001",
			CardName:      "John Doe",
			CardNumber:    "4242424242424242",
			Expiration:    "12/25",
			CVV:           "123",
			PaymentMethod: 1,
		},
	})
	assert.NoError(t, err)
	assert.True(t, resp.IsSuccess)
}

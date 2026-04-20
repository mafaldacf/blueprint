package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/digota/workflow/digota"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var orderServiceRegistry = registry.NewServiceRegistry[digota.OrderService]("order_service")

func init() {
	orderServiceRegistry.Register("local", func(ctx context.Context) (digota.OrderService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		skuService, err := skuServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}
		paymentService, err := paymentServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}

		return digota.NewOrderServiceImpl(ctx, skuService, paymentService, db)
	})
}

func TestOrderServiceNew(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a product and SKU first
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Order Product", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	sku, err := skuService.New(ctx, "Order SKU", 0, true, 1000, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 10, Type: int32(digota.Inventory_Finite)}, map[string]string{})
	assert.NoError(t, err)

	order, err := orderService.New(ctx, 0, []*digota.OrderItem{{Parent: sku.Id, Quantity: 2, Currency: 0}}, map[string]string{"order": "test"}, "order@example.com", &digota.Shipping{Name: "Test User", Address: &digota.Shipping_Address{Line1: "123 Test St", City: "Test City", State: "TS", PostalCode: "12345", Country: "US"}})
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, "order@example.com", order.Email)
	assert.Len(t, order.Items, 1)
}

func TestOrderServiceGet(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create an order
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Get Product", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	sku, err := skuService.New(ctx, "Get SKU", 0, true, 500, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 5, Type: int32(digota.Inventory_Finite)}, map[string]string{})
	assert.NoError(t, err)

	order, err := orderService.New(ctx, 0, []*digota.OrderItem{{Parent: sku.Id, Quantity: 1}}, map[string]string{}, "get@example.com", &digota.Shipping{Name: "Get User", Address: &digota.Shipping_Address{Line1: "456 Get St", City: "Get City", State: "GS", PostalCode: "67890", Country: "US"}})
	assert.NoError(t, err)

	// Get it
	retrieved, err := orderService.Get(ctx, order.Id)
	assert.NoError(t, err)
	assert.Equal(t, order.Id, retrieved.Id)
	assert.Equal(t, "get@example.com", retrieved.Email)
}

func TestOrderServiceList(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a product and SKU
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "List Product", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	sku, err := skuService.New(ctx, "List SKU", 1, true, 200, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 20, Type: int32(digota.Inventory_Finite)}, map[string]string{})
	assert.NoError(t, err)

	// Create some orders
	_, err = orderService.New(ctx, 0, []*digota.OrderItem{{Parent: sku.Id, Quantity: 1}}, map[string]string{}, "list1@example.com", &digota.Shipping{Name: "List1", Address: &digota.Shipping_Address{Line1: "1 List St", City: "List City", State: "LS", PostalCode: "11111", Country: "US"}})
	assert.NoError(t, err)
	_, err = orderService.New(ctx, 0, []*digota.OrderItem{{Parent: sku.Id, Quantity: 1}}, map[string]string{}, "list2@example.com", &digota.Shipping{Name: "List2", Address: &digota.Shipping_Address{Line1: "2 List St", City: "List City", State: "LS", PostalCode: "22222", Country: "US"}})
	assert.NoError(t, err)

	list, err := orderService.List(ctx, 0, 10, 1)
	assert.NoError(t, err)
	assert.True(t, list.Total >= 2)
}

func TestOrderServicePay(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create an order
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Pay Product", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	sku, err := skuService.New(ctx, "Pay SKU", 0, true, 1000, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 1, Type: int32(digota.Inventory_Finite)}, map[string]string{})
	assert.NoError(t, err)

	order, err := orderService.New(ctx, 0, []*digota.OrderItem{{Parent: sku.Id, Quantity: 1, Type: 1}}, map[string]string{}, "pay@example.com", &digota.Shipping{Name: "Pay User", Address: &digota.Shipping_Address{Line1: "789 Pay St", City: "Pay City", State: "PS", PostalCode: "99999", Country: "US"}})
	assert.NoError(t, err)

	// Pay it
	paid, err := orderService.Pay(ctx, order.Id, &digota.Card{Number: "4242424242424242", ExpireMonth: "12", ExpireYear: "2025", CVC: "123"}, 1)
	assert.NoError(t, err)
	assert.Equal(t, order.Id, paid.Id)
	// Assuming payment status changes
}

func TestOrderServiceReturn(t *testing.T) {
	ctx := context.Background()
	orderService, err := orderServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create and pay an order
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Return Product", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	sku, err := skuService.New(ctx, "Return SKU", 0, true, 800, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 2, Type: int32(digota.Inventory_Finite)}, map[string]string{})
	assert.NoError(t, err)

	order, err := orderService.New(ctx, 0, []*digota.OrderItem{{Parent: sku.Id, Quantity: 1, Type: 1}}, map[string]string{}, "return@example.com", &digota.Shipping{Name: "Return User", Address: &digota.Shipping_Address{Line1: "000 Return St", City: "Return City", State: "RS", PostalCode: "00000", Country: "US"}})
	assert.NoError(t, err)

	_, err = orderService.Pay(ctx, order.Id, &digota.Card{Number: "4242424242424242", ExpireMonth: "12", ExpireYear: "2025", CVC: "123"}, 1)
	assert.NoError(t, err)

	// Return it
	returned, err := orderService.Return(ctx, order.Id)
	assert.NoError(t, err)
	assert.Equal(t, order.Id, returned.Id)
}

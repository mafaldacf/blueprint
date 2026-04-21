package tests

import (
	"context"
	"testing"
	"time"

	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/basket"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/catalog"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/discount"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/order"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/web"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplequeue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testWebApp struct {
	webapp     web.WebApp
	catalogSvc catalog.CatalogService
	orderSvc   order.OrderService
}

func newTestWebApp(t *testing.T, ctx context.Context) testWebApp {
	basketDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
	require.NoError(t, err)
	orderDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
	require.NoError(t, err)
	catalogDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
	require.NoError(t, err)
	discountDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
	require.NoError(t, err)
	queue, err := simplequeue.NewSimpleQueue(ctx)
	require.NoError(t, err)

	discountSvc, err := discount.NewDiscountServiceImpl(ctx, discountDB)
	require.NoError(t, err)
	basketSvc, err := basket.NewBasketServiceImpl(ctx, basketDB, queue, discountSvc)
	require.NoError(t, err)
	catalogSvc, err := catalog.NewCatalogServiceImpl(ctx, catalogDB)
	require.NoError(t, err)
	orderSvc, err := order.NewOrderServiceImpl(ctx, orderDB, queue)
	require.NoError(t, err)

	webapp, err := web.NewWebAppImpl(ctx, basketSvc, catalogSvc, discountSvc, orderSvc)
	require.NoError(t, err)

	_, err = basketSvc.StoreBasket(ctx, basket.StoreBasketRequest{Cart: basket.ShoppingCart{UserName: "swn"}})
	require.NoError(t, err)

	return testWebApp{webapp: webapp, catalogSvc: catalogSvc, orderSvc: orderSvc}
}

func TestWebAppGetProducts(t *testing.T) {
	ctx := context.Background()
	tw := newTestWebApp(t, ctx)

	_, err := tw.catalogSvc.CreateProduct(ctx, catalog.CreateProductCommand{
		Name: "Laptop", Category: []string{"electronics"}, Price: 999.99,
	})
	require.NoError(t, err)
	_, err = tw.catalogSvc.CreateProduct(ctx, catalog.CreateProductCommand{
		Name: "Shirt", Category: []string{"clothing"}, Price: 29.99,
	})
	require.NoError(t, err)

	products, categories, _, err := tw.webapp.OnGetProductsAsync(ctx, "")
	assert.NoError(t, err)
	assert.True(t, len(products) >= 2)
	assert.True(t, len(categories) >= 2)
}

func TestWebAppGetProductsByCategory(t *testing.T) {
	ctx := context.Background()
	tw := newTestWebApp(t, ctx)

	_, err := tw.catalogSvc.CreateProduct(ctx, catalog.CreateProductCommand{
		Name: "Phone", Category: []string{"electronics"}, Price: 499.99,
	})
	require.NoError(t, err)
	_, err = tw.catalogSvc.CreateProduct(ctx, catalog.CreateProductCommand{
		Name: "Book", Category: []string{"books"}, Price: 14.99,
	})
	require.NoError(t, err)

	products, _, selectedCategory, err := tw.webapp.OnGetProductsAsync(ctx, "electronics")
	assert.NoError(t, err)
	assert.Equal(t, "electronics", selectedCategory)
	for _, p := range products {
		assert.Contains(t, p.Category, "electronics")
	}
}

func TestWebAppAddToCart(t *testing.T) {
	ctx := context.Background()
	tw := newTestWebApp(t, ctx)

	created, err := tw.catalogSvc.CreateProduct(ctx, catalog.CreateProductCommand{
		Name: "Headphones", Category: []string{"electronics"}, Price: 79.99,
	})
	require.NoError(t, err)

	err = tw.webapp.OnPostAddToCartAsync(ctx, created.Product.Id)
	assert.NoError(t, err)
}

func TestWebAppRemoveFromCart(t *testing.T) {
	ctx := context.Background()
	tw := newTestWebApp(t, ctx)

	created, err := tw.catalogSvc.CreateProduct(ctx, catalog.CreateProductCommand{
		Name: "Keyboard", Category: []string{"electronics"}, Price: 49.99,
	})
	require.NoError(t, err)

	err = tw.webapp.OnPostAddToCartAsync(ctx, created.Product.Id)
	require.NoError(t, err)

	err = tw.webapp.OnPostRemoveToCartAsync(ctx, created.Product.Id)
	assert.NoError(t, err)
}

func TestWebAppCheckoutAndGetOrders(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tw := newTestWebApp(t, ctx)

	// Start queue consumer so checkout events create orders
	go tw.orderSvc.Init(ctx)

	created, err := tw.catalogSvc.CreateProduct(ctx, catalog.CreateProductCommand{
		Name: "Monitor", Category: []string{"electronics"}, Price: 299.99,
	})
	require.NoError(t, err)

	err = tw.webapp.OnPostAddToCartAsync(ctx, created.Product.Id)
	require.NoError(t, err)

	err = tw.webapp.OnPostCheckoutAsync(ctx)
	require.NoError(t, err)

	// Give the queue consumer time to process the checkout event
	time.Sleep(100 * time.Millisecond)

	orders, err := tw.webapp.OnGetOrdersAsync(ctx)
	assert.NoError(t, err)
	assert.True(t, len(orders) >= 1)
}

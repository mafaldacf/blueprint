package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/catalog"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var catalogServiceRegistry = registry.NewServiceRegistry[catalog.CatalogService]("catalog_service")

func init() {
	catalogServiceRegistry.Register("local", func(ctx context.Context) (catalog.CatalogService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return catalog.NewCatalogServiceImpl(ctx, db)
	})
}

func TestCatalogServiceCreateProduct(t *testing.T) {
	ctx := context.Background()
	catalogService, err := catalogServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	resp, err := catalogService.CreateProduct(ctx, catalog.CreateProductCommand{
		Name:        "Test Product",
		Category:    []string{"electronics"},
		Description: "A test product",
		ImageFile:   "test.jpg",
		Price:       19.99,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp.Product)
	assert.Equal(t, "Test Product", resp.Product.Name)
	assert.Equal(t, []string{"electronics"}, resp.Product.Category)
	assert.Equal(t, 19.99, resp.Product.Price)
}

func TestCatalogServiceGetProductById(t *testing.T) {
	ctx := context.Background()
	catalogService, err := catalogServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	created, err := catalogService.CreateProduct(ctx, catalog.CreateProductCommand{
		Name:     "Get Product",
		Category: []string{"books"},
		Price:    9.99,
	})
	assert.NoError(t, err)

	resp, err := catalogService.GetProductById(ctx, catalog.GetProductByIdQuery{Id: created.Product.Id})
	assert.NoError(t, err)
	assert.Equal(t, created.Product.Id, resp.Product.Id)
	assert.Equal(t, "Get Product", resp.Product.Name)
}

func TestCatalogServiceGetProductByCategory(t *testing.T) {
	ctx := context.Background()
	catalogService, err := catalogServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = catalogService.CreateProduct(ctx, catalog.CreateProductCommand{
		Name:     "Category Product",
		Category: []string{"toys"},
		Price:    14.99,
	})
	assert.NoError(t, err)

	resp, err := catalogService.GetProductByCategory(ctx, catalog.GetProductByCategoryQuery{Category: "toys"})
	assert.NoError(t, err)
	assert.Contains(t, resp.Product.Category, "toys")
}

func TestCatalogServiceGetProducts(t *testing.T) {
	ctx := context.Background()
	catalogService, err := catalogServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = catalogService.CreateProduct(ctx, catalog.CreateProductCommand{Name: "Product A", Category: []string{"food"}, Price: 5.99})
	assert.NoError(t, err)
	_, err = catalogService.CreateProduct(ctx, catalog.CreateProductCommand{Name: "Product B", Category: []string{"food"}, Price: 6.99})
	assert.NoError(t, err)

	resp, err := catalogService.GetProducts(ctx)
	assert.NoError(t, err)
	assert.True(t, len(resp.Products) >= 2)
}

func TestCatalogServiceDeleteProduct(t *testing.T) {
	ctx := context.Background()
	catalogService, err := catalogServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	created, err := catalogService.CreateProduct(ctx, catalog.CreateProductCommand{
		Name:     "Delete Me",
		Category: []string{"misc"},
		Price:    1.99,
	})
	assert.NoError(t, err)

	err = catalogService.DeleteProduct(ctx, catalog.DeleteProductCommand{Id: created.Product.Id})
	assert.NoError(t, err)

	_, err = catalogService.GetProductById(ctx, catalog.GetProductByIdQuery{Id: created.Product.Id})
	assert.Error(t, err)
}

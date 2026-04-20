package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/digota/workflow/digota"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var skuServiceRegistry = registry.NewServiceRegistry[digota.SkuService]("sku_service")

func init() {
	skuServiceRegistry.Register("local", func(ctx context.Context) (digota.SkuService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		productService, err := productServiceRegistry.Get(ctx)
		if err != nil {
			return nil, err
		}

		return digota.NewSkuServiceImpl(ctx, productService, db)
	})
}

func TestSkuServiceNew(t *testing.T) {
	ctx := context.Background()
	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// First create a product
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Parent Product", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	sku, err := skuService.New(ctx, "Test SKU", 1, true, 1000, product.Id, map[string]string{"key": "value"}, "image.jpg", &digota.PackageDimensions{Weight: 10, Height: 5, Width: 5, Length: 5}, &digota.Inventory{Quantity: 100, Type: 1}, map[string]string{"attr": "val"})
	assert.NoError(t, err)
	assert.NotNil(t, sku)
	assert.Equal(t, "Test SKU", sku.Name)
	assert.True(t, sku.Active)
	assert.Equal(t, uint64(1000), sku.Price)
}

func TestSkuServiceGet(t *testing.T) {
	ctx := context.Background()
	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a product and SKU
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Parent", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	sku, err := skuService.New(ctx, "Test SKU", 1, true, 500, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 50, Type: 1}, map[string]string{})
	assert.NoError(t, err)

	// Get it
	retrieved, err := skuService.Get(ctx, sku.Id)
	assert.NoError(t, err)
	assert.Equal(t, sku.Id, retrieved.Id)
	assert.Equal(t, "Test SKU", retrieved.Name)
}

func TestSkuServiceList(t *testing.T) {
	ctx := context.Background()
	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a product
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Parent", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	// Create some SKUs
	_, err = skuService.New(ctx, "SKU 1", 1, true, 100, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 10, Type: 1}, map[string]string{})
	assert.NoError(t, err)
	_, err = skuService.New(ctx, "SKU 2", 1, true, 200, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 20, Type: 1}, map[string]string{})
	assert.NoError(t, err)

	list, err := skuService.List(ctx, 0, 10, 1)
	assert.NoError(t, err)
	assert.True(t, list.Total >= 2)
}

func TestSkuServiceUpdate(t *testing.T) {
	ctx := context.Background()
	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a product and SKU
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Parent", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	sku, err := skuService.New(ctx, "Original SKU", 1, true, 300, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 30, Type: 1}, map[string]string{})
	assert.NoError(t, err)

	// Update it
	updated, err := skuService.Update(ctx, sku.Id, "Updated SKU", 1, false, 400, product.Id, map[string]string{"updated": "yes"}, "newimage.jpg", &digota.PackageDimensions{Weight: 20}, &digota.Inventory{Quantity: 40, Type: 1}, map[string]string{"newattr": "newval"})
	assert.NoError(t, err)
	assert.Equal(t, "Updated SKU", updated.Name)
	assert.False(t, updated.Active)
	assert.Equal(t, uint64(400), updated.Price)
}

func TestSkuServiceDelete(t *testing.T) {
	ctx := context.Background()
	skuService, err := skuServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a product and SKU
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)
	product, err := productService.New(ctx, "Parent", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	sku, err := skuService.New(ctx, "To Delete", 1, true, 100, product.Id, map[string]string{}, "", nil, &digota.Inventory{Quantity: 5, Type: 1}, map[string]string{})
	assert.NoError(t, err)

	// Delete it
	err = skuService.Delete(ctx, sku.Id)
	assert.NoError(t, err)

	// Try to get it
	sku, err = skuService.Get(ctx, sku.Id)
	assert.Error(t, err)
}

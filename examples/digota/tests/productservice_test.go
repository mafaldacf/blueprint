package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/digota/workflow/digota"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var productServiceRegistry = registry.NewServiceRegistry[digota.ProductService]("product_service")

func init() {
	productServiceRegistry.Register("local", func(ctx context.Context) (digota.ProductService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}

		return digota.NewProductServiceImpl(ctx, db)
	})
}

func TestProductServiceNew(t *testing.T) {
	ctx := context.Background()
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	product, err := productService.New(ctx, "Test Product", true, []string{"attr1", "attr2"}, "Description", []string{"image1.jpg"}, map[string]string{"key": "value"}, true, "http://example.com")
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Test Product", product.Name)
	assert.True(t, product.Active)
}

func TestProductServiceGet(t *testing.T) {
	ctx := context.Background()
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// First create a product
	product, err := productService.New(ctx, "Test Product", true, []string{"attr1"}, "Description", []string{}, map[string]string{}, true, "http://example.com")
	assert.NoError(t, err)

	// Now get it
	retrieved, err := productService.Get(ctx, product.Id)
	assert.NoError(t, err)
	assert.Equal(t, product.Id, retrieved.Id)
	assert.Equal(t, "Test Product", retrieved.Name)
}

func TestProductServiceList(t *testing.T) {
	ctx := context.Background()
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create some products
	_, err = productService.New(ctx, "Product 1", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)
	_, err = productService.New(ctx, "Product 2", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	list, err := productService.List(ctx, 0, 10, 1)
	assert.NoError(t, err)
	assert.True(t, list.Total >= 2)
}

func TestProductServiceUpdate(t *testing.T) {
	ctx := context.Background()
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a product
	product, err := productService.New(ctx, "Original", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	// Update it
	updated, err := productService.Update(ctx, product.Id, "Updated", false, []string{"newattr"}, "New desc", []string{"newimage.jpg"}, map[string]string{"newkey": "newvalue"}, false, "http://updated.com")
	assert.NoError(t, err)
	assert.Equal(t, "Updated", updated.Name)
	assert.False(t, updated.Active)
}

func TestProductServiceDelete(t *testing.T) {
	ctx := context.Background()
	productService, err := productServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	// Create a product
	product, err := productService.New(ctx, "To Delete", true, []string{}, "", []string{}, map[string]string{}, true, "")
	assert.NoError(t, err)

	// Delete it
	err = productService.Delete(ctx, product.Id)
	assert.NoError(t, err)

	// Try to get it - should fail
	_, err = productService.Get(ctx, product.Id)
	assert.Error(t, err)
}

package shopping_simple

import (
	"context"
)

type Frontend interface {
	CreateProduct(ctx context.Context, productID string, description string, pricePerUnit int, category string) (Product, error)
	CreateCart(ctx context.Context, cartID string) (Cart, error)


	GetAllProducts(ctx context.Context) ([]Product, error)
	GetCart(ctx context.Context, cartID string) (Cart, error)
	GetProduct(ctx context.Context, productID string) (Product, error)

	
	AddProductToCart(ctx context.Context, cartID string, productID string) (CartProduct, error)
	DeleteProduct(ctx context.Context, productID string) (bool, error)
}

type FrontendImpl struct {
	product_service ProductService
	cart_service    CartService
}

func NewFrontendImpl(ctx context.Context, product_service ProductService, cart_service CartService) (Frontend, error) {
	return &FrontendImpl{product_service: product_service, cart_service: cart_service}, nil
}

func (f *FrontendImpl) CreateProduct(ctx context.Context, productID string, description string, pricePerUnit int, category string) (Product, error) {
	return f.product_service.CreateProduct(ctx, productID, description, pricePerUnit, category)
}

func (f *FrontendImpl) CreateCart(ctx context.Context, cartID string) (Cart, error) {
	return f.cart_service.CreateCart(ctx, cartID)
}

func (f *FrontendImpl) GetAllProducts(ctx context.Context) ([]Product, error) {
	return f.product_service.GetAllProducts(ctx)
}

func (f *FrontendImpl) GetProduct(ctx context.Context, productID string) (Product, error) {
	return f.product_service.GetProduct(ctx, productID)
}

func (f *FrontendImpl) GetCart(ctx context.Context, cartID string) (Cart, error) {
	return f.cart_service.GetCart(ctx, cartID)
}

func (f *FrontendImpl) AddProductToCart(ctx context.Context, cartID string, productID string) (CartProduct, error) {
	return f.cart_service.AddProductToCart(ctx, cartID, productID)
}

func (f *FrontendImpl) DeleteProduct(ctx context.Context, productID string) (bool, error) {
	return f.product_service.DeleteProduct(ctx, productID)
}

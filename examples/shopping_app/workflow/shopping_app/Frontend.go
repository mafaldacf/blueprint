package shopping_app

import (
	"context"
)

type Frontend interface {
	AddProductToCart(ctx context.Context, cartID string, productID string) error
	Checkout(ctx context.Context, cartID string, userID string, productID string, quantity int, price int) error
	ReadOrder(ctx context.Context, orderID string) (Order, error)
}

type FrontendImpl struct {
	cart_service  CartService
	order_service OrderService
}

func NewFrontendImpl(ctx context.Context, order_service OrderService, cart_service CartService) (Frontend, error) {
	return &FrontendImpl{order_service: order_service, cart_service: cart_service}, nil
}

func (f *FrontendImpl) ReadOrder(ctx context.Context, orderID string) (Order, error) {
	return f.order_service.ReadOrder(ctx, orderID)
}

func (f *FrontendImpl) Checkout(ctx context.Context, cartID string, userID string, productID string, quantity int, price int) error {
	return f.order_service.CreateOrder(ctx, cartID, userID, productID, quantity, price)
}

func (f *FrontendImpl) AddProductToCart(ctx context.Context, cartID string, productID string) error {
	return f.cart_service.AddProductToCart(ctx, cartID, productID)
}

package eshopmicroservices

import (
	"context"

	"github.com/google/uuid"
)

type WebApp interface {
	OnPostAddToCartAsync(ctx context.Context, productId uuid.UUID) error
}

type WebAppImpl struct {
	basketService   BasketService
	catalogService  CatalogService
	discountService DiscountService
	orderService    OrderService
}

func NewWebAppImpl(ctx context.Context, basketService BasketService, catalogService CatalogService, discountService DiscountService, orderService OrderService) (WebApp, error) {
	s := &WebAppImpl{
		basketService:   basketService,
		catalogService:  catalogService,
		discountService: discountService,
		orderService:    orderService,
	}
	return s, nil
}

var quantity int
var color string

func (webapp *WebAppImpl) OnPostAddToCartAsync(ctx context.Context, productId uuid.UUID) error {
	productResponse, err := webapp.catalogService.GetProductById(ctx, GetProductByIdQuery{Id: productId})
	if err != nil {
		return err
	}

	basketResponse, err := webapp.basketService.GetBasket(ctx, GetBasketQuery{UserName: "swn"})
	if err != nil {
		return err
	}
	basket := basketResponse.Cart
	basket.Items = append(basket.Items, ShoppingCartItem{
		ProductId:   productId,
		ProductName: productResponse.Product.Name,
		Price:       productResponse.Product.Price,
		Quantity:    quantity,
		Color:       color,
	})

	_, err = webapp.basketService.StoreBasket(ctx, StoreBasketRequest{Cart: basket})
	if err != nil {
		return err
	}

	return nil
}

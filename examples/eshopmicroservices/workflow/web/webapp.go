package web

import (
	"context"

	"github.com/google/uuid"

	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/basket"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/catalog"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/discount"
	"github.com/blueprint-uservices/blueprint/examples/eshopmicroservices/workflow/order"
)

type WebApp interface {
	OnPostRemoveToCartAsync(ctx context.Context, productId uuid.UUID) error
	OnPostCheckoutAsync(ctx context.Context) error
	OnGetOrdersAsync(ctx context.Context) ([]order.OrderDto, error)
	OnGetProductsAsync(ctx context.Context, categoryName string) ([]catalog.Product, []string, string, error)
	OnPostAddToCartAsync(ctx context.Context, productId uuid.UUID) error
}

type WebAppImpl struct {
	basketService   basket.BasketService
	catalogService  catalog.CatalogService
	discountService discount.DiscountService
	orderService    order.OrderService
	customerId      uuid.UUID
}

func NewWebAppImpl(ctx context.Context, basketService basket.BasketService, catalogService catalog.CatalogService, discountService discount.DiscountService, orderService order.OrderService) (WebApp, error) {
	s := &WebAppImpl{
		basketService:   basketService,
		catalogService:  catalogService,
		discountService: discountService,
		orderService:    orderService,
		customerId:      uuid.MustParse("5334c996-8457-4cf0-815c-ed2b77c4ff61"),
	}
	return s, nil
}

var quantity int
var color string

func (webapp *WebAppImpl) OnPostRemoveToCartAsync(ctx context.Context, productId uuid.UUID) error {
	basketResponse, err := webapp.basketService.GetBasket(ctx, basket.GetBasketQuery{UserName: "swn"})
	if err != nil {
		return err
	}
	cart := basketResponse.Cart

	removeAll(&cart, productId)

	_, err = webapp.basketService.StoreBasket(ctx, basket.StoreBasketRequest{Cart: cart})
	if err != nil {
		return err
	}

	return nil
}

func removeAll(cart *basket.ShoppingCart, productId uuid.UUID) []basket.ShoppingCartItem {
	remainingItems := cart.Items[:0]
	for _, item := range cart.Items {
		if item.ProductId != productId {
			remainingItems = append(remainingItems, item)
		}
	}
	return remainingItems
}

func (webapp *WebAppImpl) OnPostCheckoutAsync(ctx context.Context) error {
	basketResponse, err := webapp.basketService.GetBasket(ctx, basket.GetBasketQuery{UserName: "swn"})
	if err != nil {
		return err
	}
	cart := basketResponse.Cart

	// assumption customerId is passed in from the UI authenticated user swn
	var order basket.BasketCheckoutDto
	order.CustomerId = webapp.customerId
	order.UserName = cart.UserName
	order.TotalPrice = cart.TotalPrice

	webapp.basketService.CheckoutBasket(ctx, basket.CheckoutBasketCommand{order})

	return nil
}

func (webapp *WebAppImpl) OnGetOrdersAsync(ctx context.Context) ([]order.OrderDto, error) {
	// assumption customerId is passed in from the UI authenticated user swn
	customerId := webapp.customerId
	response, err := webapp.orderService.GetOrdersByCustomer(ctx, order.GetOrdersByCustomerQuery{customerId})
	if err != nil {
		return nil, err
	}
	return response.Orders, nil
}

func (webapp *WebAppImpl) OnGetProductsAsync(ctx context.Context, categoryName string) ([]catalog.Product, []string, string, error) {
	response, err := webapp.catalogService.GetProducts(ctx)
	if err != nil {
		return nil, nil, "", err
	}

	categorySet := make(map[string]bool)
	for _, p := range response.Products {
		for _, c := range p.Category {
			categorySet[c] = true
		}
	}
	var categoryList []string
	for c := range categorySet {
		categoryList = append(categoryList, c)
	}

	var productList []catalog.Product
	var selectedCategory string

	if categoryName != "" {
		for _, p := range response.Products {
			for _, c := range p.Category {
				if c == categoryName {
					productList = append(productList, p)
					break
				}
			}
		}
		selectedCategory = categoryName
	} else {
		productList = response.Products
	}
	return productList, categoryList, selectedCategory, nil

}

func (webapp *WebAppImpl) OnPostAddToCartAsync(ctx context.Context, productId uuid.UUID) error {
	productResponse, err := webapp.catalogService.GetProductById(ctx, catalog.GetProductByIdQuery{Id: productId})
	if err != nil {
		return err
	}

	basketResponse, err := webapp.basketService.GetBasket(ctx, basket.GetBasketQuery{UserName: "swn"})
	if err != nil {
		return err
	}
	retBasket := basketResponse.Cart
	retBasket.Items = append(retBasket.Items, basket.ShoppingCartItem{
		ProductId:   productId,
		ProductName: productResponse.Product.Name,
		Price:       productResponse.Product.Price,
		Quantity:    quantity,
		Color:       color,
	})

	_, err = webapp.basketService.StoreBasket(ctx, basket.StoreBasketRequest{Cart: retBasket})
	if err != nil {
		return err
	}

	return nil
}

package eshopmicroservices

type CheckoutBasketCommand struct {
	BasketCheckoutDto BasketCheckoutDto
}

type CheckoutBasketResponse struct {
	IsSuccess bool
}

type StoreBasketRequest struct {
	Cart ShoppingCart
}

type StoreBasketResponse struct {
	UserName string
}

type GetBasketQuery struct {
	UserName string
}

type GetBasketResult struct {
	Cart ShoppingCart
}

type DeleteBasketCommand struct {
	UserName string
}

type DeleteBasketResult struct {
	IsSuccess bool
}

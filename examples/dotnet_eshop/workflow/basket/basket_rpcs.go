package basket

type GetBasketRequest struct {
	UserName string
}

type UpdateBasketRequest struct {
	Cart CustomerBasket
}

type DeleteBasketRequest struct {
	UserName string
}

type CustomerBasketResponse struct {
	BasketItems []BasketItem
}

type DeleteBasketResponse struct {
}

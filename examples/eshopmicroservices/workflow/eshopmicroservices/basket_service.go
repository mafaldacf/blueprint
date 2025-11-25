package eshopmicroservices

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type BasketService interface {
	CheckoutBasket(ctx context.Context, request CheckoutBasketCommand) (CheckoutBasketResponse, error)
	StoreBasket(ctx context.Context, request StoreBasketRequest) (StoreBasketResponse, error)
	GetBasket(ctx context.Context, query GetBasketQuery) (GetBasketResult, error)
	DeleteBasket(ctx context.Context, query DeleteBasketCommand) (DeleteBasketResult, error)
}

type BasketServiceImpl struct {
	database        backend.NoSQLDatabase
	queue           backend.Queue
	discountService DiscountService
}

func NewBasketServiceImpl(ctx context.Context, database backend.NoSQLDatabase, queue backend.Queue, discountService DiscountService) (BasketService, error) {
	s := &BasketServiceImpl{
		database:        database,
		queue:           queue,
		discountService: discountService,
	}
	return s, nil
}

func (s *BasketServiceImpl) CheckoutBasket(ctx context.Context, command CheckoutBasketCommand) (CheckoutBasketResponse, error) {
	basket, err := s.getBasket(ctx, command.BasketCheckoutDto.UserName)
	if err != nil {
		return CheckoutBasketResponse{IsSuccess: false}, err
	}

	eventMessage := adapt(command.BasketCheckoutDto)
	eventMessage.TotalPrice = basket.TotalPrice

	ok, err := s.queue.Push(ctx, eventMessage)
	if err != nil || !ok {
		return CheckoutBasketResponse{IsSuccess: false}, err
	}

	err = s.deleteBasket(ctx, command.BasketCheckoutDto.UserName)
	if err != nil {
		return CheckoutBasketResponse{IsSuccess: false}, err
	}
	return CheckoutBasketResponse{IsSuccess: true}, nil
}

func (s *BasketServiceImpl) StoreBasket(ctx context.Context, command StoreBasketRequest) (StoreBasketResponse, error) {
	// deduct discount
	for _, item := range command.Cart.Items {
		coupon, err := s.discountService.GetDiscount(ctx, GetDiscountRequest{ProductName: item.ProductName})
		if err != nil {
			return StoreBasketResponse{}, err
		}
		item.Price -= coupon.Amount
	}

	err := s.storeBasket(ctx, command.Cart)
	if err != nil {
		return StoreBasketResponse{}, err
	}
	return StoreBasketResponse{UserName: command.Cart.UserName}, nil
}

func (s *BasketServiceImpl) GetBasket(ctx context.Context, query GetBasketQuery) (GetBasketResult, error) {
	basket, err := s.getBasket(ctx, query.UserName)
	if err != nil {
		return GetBasketResult{}, err
	}
	return GetBasketResult{Cart: basket}, nil
}

func (s *BasketServiceImpl) DeleteBasket(ctx context.Context, query DeleteBasketCommand) (DeleteBasketResult, error) {
	err := s.deleteBasket(ctx, query.UserName)
	if err != nil {
		return DeleteBasketResult{IsSuccess: false}, err
	}
	return DeleteBasketResult{IsSuccess: true}, nil
}

func (s *BasketServiceImpl) getBasket(ctx context.Context, username string) (ShoppingCart, error) {
	collection, err := s.database.GetCollection(ctx, "basket_db", "basket")
	if err != nil {
		return ShoppingCart{}, err
	}
	filter := bson.D{{Key: "UserName", Value: username}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return ShoppingCart{}, err
	}
	var cart ShoppingCart
	ok, err := cursor.One(ctx, &cart)
	if err != nil {
		return ShoppingCart{}, err
	}
	if !ok {
		return ShoppingCart{}, fmt.Errorf("basket not found for username (%s)", username)
	}
	return cart, nil
}

func (s *BasketServiceImpl) storeBasket(ctx context.Context, basket ShoppingCart) error {
	collection, err := s.database.GetCollection(ctx, "basket_db", "basket")
	if err != nil {
		return err
	}
	return collection.InsertOne(ctx, basket)
}

func (s *BasketServiceImpl) deleteBasket(ctx context.Context, username string) error {
	collection, err := s.database.GetCollection(ctx, "basket_db", "basket")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "UserName", Value: username}}
	return collection.DeleteOne(ctx, filter)
}

func adapt(basket BasketCheckoutDto) BasketChekoutEvent {
	return BasketChekoutEvent{
		basket.UserName,
		basket.CustomerId,
		basket.TotalPrice,
		basket.FirstName,
		basket.LastName,
		basket.EmailAddress,
		basket.AddressLine,
		basket.Country,
		basket.State,
		basket.ZipCode,
		basket.CardName,
		basket.CardNumber,
		basket.Expiration,
		basket.CVV,
		basket.PaymentMethod,
	}
}

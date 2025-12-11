package basket

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type BasketService interface {
	Run(ctx context.Context) error
	Get(ctx context.Context, query GetBasketRequest) (CustomerBasketResponse, error)
	Update(ctx context.Context, request UpdateBasketRequest) (CustomerBasketResponse, error)
	Delete(ctx context.Context, query DeleteBasketRequest) (DeleteBasketResponse, error)
}

type BasketServiceImpl struct {
	database      backend.NoSQLDatabase
	queue         backend.Queue
	exit_on_error bool
}

func NewBasketServiceImpl(ctx context.Context, database backend.NoSQLDatabase, queue backend.Queue) (BasketService, error) {
	s := &BasketServiceImpl{
		database:      database,
		queue:         queue,
		exit_on_error: false,
	}
	return s, nil
}

func (s *BasketServiceImpl) Update(ctx context.Context, command UpdateBasketRequest) (CustomerBasketResponse, error) {
	err := s.update(ctx, command.Cart)
	if err != nil {
		return CustomerBasketResponse{}, err
	}
	return CustomerBasketResponse{BasketItems: command.Cart.Items}, nil
}

func (s *BasketServiceImpl) Get(ctx context.Context, query GetBasketRequest) (CustomerBasketResponse, error) {
	basket, err := s.getBasket(ctx, query.UserName)
	if err != nil {
		return CustomerBasketResponse{}, err
	}
	return CustomerBasketResponse{BasketItems: basket.Items}, nil
}

func (s *BasketServiceImpl) Delete(ctx context.Context, query DeleteBasketRequest) (DeleteBasketResponse, error) {
	err := s.deleteBasket(ctx, query.UserName)
	if err != nil {
		return DeleteBasketResponse{}, err
	}
	return DeleteBasketResponse{}, nil
}

func (s *BasketServiceImpl) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			var event OrderStartedEvent
			ok, err := s.queue.Pop(ctx, &event)
			if err != nil && s.exit_on_error {
				return err
			}
			if !ok {
				continue
			}
			s.deleteBasket(ctx, event.UserID)
		}
	}
	return nil
}

func (s *BasketServiceImpl) getBasket(ctx context.Context, username string) (CustomerBasket, error) {
	collection, err := s.database.GetCollection(ctx, "basket_db", "basket")
	if err != nil {
		return CustomerBasket{}, err
	}
	filter := bson.D{{Key: "UserName", Value: username}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return CustomerBasket{}, err
	}
	var cart CustomerBasket
	ok, err := cursor.One(ctx, &cart)
	if err != nil {
		return CustomerBasket{}, err
	}
	if !ok {
		return CustomerBasket{}, fmt.Errorf("basket not found for username (%s)", username)
	}
	return cart, nil
}

func (s *BasketServiceImpl) update(ctx context.Context, basket CustomerBasket) error {
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

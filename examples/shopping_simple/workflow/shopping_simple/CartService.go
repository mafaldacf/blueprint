package shopping_simple

import (
	"context"
	"fmt"
	"sync"

	backend "github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

const DECLARED_CONSTANT = "THIS IS A CONSTANT!"

type CartService interface {
	AddProductToCart(ctx context.Context, cartID string, productID string) (CartProduct, error)
	GetCart(ctx context.Context, cartID string) (Cart, error)
	CreateCart(ctx context.Context, cartID string) (Cart, error)
	Run(ctx context.Context) error
}

type CartServiceImpl struct {
	product_service ProductService
	cart_db         backend.NoSQLDatabase
	product_queue   backend.Queue
	num_workers     int
}

func NewCartServiceImpl(ctx context.Context, product_service ProductService, cart_db backend.NoSQLDatabase, product_queue backend.Queue) (CartService, error) {
	return &CartServiceImpl{product_service: product_service, cart_db: cart_db, product_queue: product_queue, num_workers: 4}, nil
}

func (s *CartServiceImpl) CreateCart(ctx context.Context, cartID string) (Cart, error) {
	collection, _ := s.cart_db.GetCollection(ctx, "cart_database", "cart_database")
	cart := Cart{
		CartID: cartID,
	}
	err := collection.InsertOne(ctx, cart)
	return cart, err
}

func (s *CartServiceImpl) GetCart(ctx context.Context, cartID string) (Cart, error) {
	var cart Cart
	collection, _ := s.cart_db.GetCollection(ctx, "cart_database", "cart_database")
	filter := bson.D{{Key: "cartid", Value: cartID}}
	result, err := collection.FindOne(ctx, filter)
	if err != nil {
		return cart, fmt.Errorf("no cart found for id '%s': %s", cartID, err.Error())
	}
	exists, err := result.One(ctx, &cart)

	for _, p := range cart.Products {
		s.product_service.GetProduct(ctx, p)
	}
	

	if !exists {
		return cart, fmt.Errorf("no cart found for id '%s'", cartID)
	}
	return cart, err
}

func (s *CartServiceImpl) AddProductToCart(ctx context.Context, cartID string, productID string) (CartProduct, error) {
	var product Product
	product, _ = s.product_service.GetProduct(ctx, productID)

	var cart Cart
	var cartProduct CartProduct
	collection, _ := s.cart_db.GetCollection(ctx, "cart_database", "cart_collection")
	filter := bson.D{{Key: "cartid", Value: cartID}}

	result, err := collection.FindOne(ctx, filter)
	if err != nil {
		return cartProduct, err
	}

	cartProduct = CartProduct{
		CartID:       cartID,
		ProductID:    productID,
		PricePerUnit: product.PricePerUnit,
		Quantity:     1,
	}

	exists, _ := result.One(ctx, &cart)
	if !exists { // creates if it does not exist yet
		cart = Cart{
			CartID:        cartID,
			TotalQuantity: 1,
			LastProductID: productID,
			Products:      []string{productID},
		}
		collection.InsertOne(ctx, cart)
		return cartProduct, err
	}

	/* update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "totalquantity", Value: 1},
		}},
		{Key: "$set", Value: bson.D{
			{Key: "lastproductid", Value: product.ProductID},
		}},
		{Key: "$push", Value: bson.D{
			{Key: "products", Value: cartProduct.ProductID},
		}},
	}
	updated, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cartProduct, nil
	}
	if updated == 0 {
		return cartProduct, fmt.Errorf("no cart to update for id '%s'", cartID)
	} */
	return cartProduct, err
}

func (s *CartServiceImpl) removeProduct(ctx context.Context, message ProductQueueMessage) error {
	collection, _ := s.cart_db.GetCollection(ctx, "cart_database", "cart_database")
	filter := bson.D{}
	update := bson.D{
		{Key: "$dec", Value: bson.D{
			{Key: "totalquantity", Value: 1},
		}},
		{Key: "$pull", Value: bson.D{
			{Key: "products", Value: message.ProductID},
		}},
	}
	updated, err := collection.UpdateMany(ctx, filter, update)
	if updated == 0 {
		return fmt.Errorf("no cart to update for product id '%s'", message.ProductID)
	}
	return err
}

func (s *CartServiceImpl) workerThread(ctx context.Context) error {
	var forever chan struct{}
	go func() {
		var event map[string]interface{}
		s.product_queue.Pop(ctx, &event)
		workerMessage := ProductQueueMessage{
			ProductID: event["ProductID"].(string),
			Remove:    event["Remove"].(bool),
		}
		s.removeProduct(ctx, workerMessage)
	}()
	<-forever
	return nil
}

func (s *CartServiceImpl) Run(ctx context.Context) error {
	backend.GetLogger().Info(ctx, "initializing %d workers", s.num_workers)
	var wg sync.WaitGroup
	wg.Add(s.num_workers)
	for i := 1; i <= s.num_workers; i++ {
		go func(i int) {
			defer wg.Done()
			err := s.workerThread(ctx)
			if err != nil {
				backend.GetLogger().Error(ctx, "error in worker thread: %s", err.Error())
				panic(err)
			}
		}(i)
	}
	wg.Wait()
	backend.GetLogger().Info(ctx, "joining %d workers", s.num_workers)
	return nil
}


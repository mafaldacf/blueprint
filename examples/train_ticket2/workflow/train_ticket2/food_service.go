package train_ticket2

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type FoodService interface {
	CreateFoodOrder(ctx context.Context, addFoodOrder FoodOrder) (FoodOrder, error)
	// extra
	FindFoodOrder(ctx context.Context, orderID string) (FoodOrder, error)
}

type FoodServiceImpl struct {
	foodDB        backend.NoSQLDatabase
	deliveryQueue backend.Queue
}

func NewFoodServiceImpl(ctx context.Context, foodDB backend.NoSQLDatabase, foodOrderQueue backend.Queue) (FoodService, error) {
	return &FoodServiceImpl{foodDB: foodDB, deliveryQueue: foodOrderQueue}, nil
}

func (c *FoodServiceImpl) CreateFoodOrder(ctx context.Context, addFoodOrder FoodOrder) (FoodOrder, error) {
	foodOrder := FoodOrder{
		OrderID:     addFoodOrder.OrderID,
		FoodType:    addFoodOrder.FoodType,
		StationName: addFoodOrder.StationName,
		StoreName:   addFoodOrder.StoreName,
		FoodName:    addFoodOrder.FoodName,
		Price:       addFoodOrder.Price,
	}

	collection, err := c.foodDB.GetCollection(ctx, "food_db", "food_order")
	if err != nil {
		return FoodOrder{}, err
	}
	err = collection.InsertOne(ctx, foodOrder)
	if err != nil {
		return FoodOrder{}, err
	}

	delivery := Delivery{
		OrderID:     addFoodOrder.OrderID,
		FoodName:    addFoodOrder.FoodName,
		StationName: addFoodOrder.StationName,
		StoreName:   addFoodOrder.StoreName,
	}

	_, err = c.deliveryQueue.Push(ctx, delivery)
	if err != nil {
		return FoodOrder{}, err
	}

	return foodOrder, nil
}

func (c *FoodServiceImpl) FindFoodOrder(ctx context.Context, orderID string) (FoodOrder, error) {
	collection, err := c.foodDB.GetCollection(ctx, "food_db", "food_order")
	if err != nil {
		return FoodOrder{}, err
	}

	filter := bson.D{{Key: "OrderID", Value: orderID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return FoodOrder{}, err
	}

	var foodOrder FoodOrder
	ok, err := cursor.One(ctx, &foodOrder)
	if err != nil {
		return FoodOrder{}, err
	}
	if !ok {
		return FoodOrder{}, fmt.Errorf("food order (%s) not found", orderID)
	}
	return foodOrder, nil
}

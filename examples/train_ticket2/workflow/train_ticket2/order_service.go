package train_ticket2

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderService interface {
	Create(ctx context.Context, order Order) (Order, error)
	UpdateStatus(ctx context.Context, order Order) (Order, error)
	Find(ctx context.Context, orderID string) (Order, error)
	FindAll(ctx context.Context) ([]Order, error)
	Delete(ctx context.Context, id string) error
	ModifyOrder(ctx context.Context, id string, status int) error
	GetTicketListByDateAndTripID(ctx context.Context, seatRequest SeatRequest) (LeftTicketInfo, error)
}

type OrderServiceImpl struct {
	orderDB backend.NoSQLDatabase
}

func NewOrderServiceImpl(ctx context.Context, orderDB backend.NoSQLDatabase) (OrderService, error) {
	return &OrderServiceImpl{orderDB: orderDB}, nil
}

func (o *OrderServiceImpl) Create(ctx context.Context, order Order) (Order, error) {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return Order{}, err
	}

	filter := bson.D{{Key: "ID", Value: order.ID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Order{}, nil
	}

	var orderTmp Order
	ok, err := cursor.One(ctx, &orderTmp)
	if err != nil {
		return Order{}, err
	}
	if ok {
		return Order{}, fmt.Errorf("order (%s) already exists", order.ID)
	}

	err = collection.InsertOne(ctx, order)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (o *OrderServiceImpl) UpdateStatus(ctx context.Context, order Order) (Order, error) {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return Order{}, err
	}

	filter := bson.D{{Key: "ID", Value: order.ID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Order{}, nil
	}

	var oldOrder Order
	ok, err := cursor.One(ctx, &oldOrder)
	if err != nil {
		return Order{}, err
	}
	if !ok {
		return Order{}, fmt.Errorf("order (%s) does not exist", order.ID)
	}

	_, err = collection.ReplaceOne(ctx, filter, order)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (o *OrderServiceImpl) Find(ctx context.Context, orderID string) (Order, error) {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return Order{}, err
	}

	filter := bson.D{{Key: "ID", Value: orderID}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Order{}, err
	}

	var order Order
	ok, err := cursor.One(ctx, &order)
	if err != nil {
		return Order{}, err
	}
	if !ok {
		return Order{}, fmt.Errorf("order (%s) not found", orderID)
	}
	return order, nil
}

func (o *OrderServiceImpl) FindAll(ctx context.Context) ([]Order, error) {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return nil, err
	}
	cursor, err := collection.FindMany(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var orders []Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *OrderServiceImpl) Delete(ctx context.Context, id string) error {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return err
	}
	return collection.DeleteOne(ctx, bson.D{{Key: "ID", Value: id}})
}

func (o *OrderServiceImpl) ModifyOrder(ctx context.Context, id string, status int) error {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "ID", Value: id}}
	update := bson.D{{Key: "Status", Value: status}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderServiceImpl) GetTicketListByDateAndTripID(ctx context.Context, seatRequest SeatRequest) (LeftTicketInfo, error) {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return LeftTicketInfo{}, err
	}

	filter := bson.D{{Key: "TravelDate", Value: seatRequest.TravelDate}, {Key: "TrainNumber", Value: seatRequest.TrainNumber}}
	cursor, err := collection.FindMany(ctx, filter)
	if err != nil {
		return LeftTicketInfo{}, nil
	}

	var orders []Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		return LeftTicketInfo{}, nil
	}

	if len(orders) == 0 {
		return LeftTicketInfo{}, fmt.Errorf("left ticket info is empty")
	}

	var soldTickets = make([]Ticket, len(orders))
	for i, order := range orders {
		soldTickets[i] = Ticket{
			StartStation: order.FromStation,
			DestStation:  order.ToStation,
		}
	}
	leftTicketInfo := LeftTicketInfo{SoldTickets: soldTickets}
	return leftTicketInfo, nil
}

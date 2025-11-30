package trainticket

import (
	"context"
	"fmt"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderService interface {
	QueryOrders(ctx context.Context, qi OrderInfo, accountId string) ([]Order, error)
	CalculateSoldTicket(ctx context.Context, travelDate time.Time, trainNumber string) (SoldTicket, error)
	SecurityInfoCheck(ctx context.Context, dateFrom time.Time, accountId string) (OrderSecurity, error)
	GetOrderPrice(ctx context.Context, orderId string) (string, error)
	PayOrder(ctx context.Context, orderId string) (bool, error)
	AddCreateNewOrder(ctx context.Context, order Order) error
	CreateNewOrder(ctx context.Context, order Order) (Order, error)
	UpdateOrder(ctx context.Context, order Order) error
	SaveOrderInfo(ctx context.Context, order Order) (Order, error)
	GetOrderById(ctx context.Context, orderID string) (Order, error)
	FindAllOrder(ctx context.Context) ([]Order, error)
	DeleteOrder(ctx context.Context, id string) error
	ModifyOrder(ctx context.Context, id string, status int) error
	GetTicketListByDateAndTripID(ctx context.Context, seatRequest SeatRequest) (LeftTicketInfo, error)
}

type OrderServiceImpl struct {
	orderDB backend.NoSQLDatabase
}

func NewOrderServiceImpl(ctx context.Context, orderDB backend.NoSQLDatabase) (OrderService, error) {
	return &OrderServiceImpl{orderDB: orderDB}, nil
}

func (o *OrderServiceImpl) QueryOrders(ctx context.Context, qi OrderInfo, accountId string) ([]Order, error) {
	//1.Get all orders of the user
	orders, err := o.findByAccountId(ctx, accountId)
	if err != nil {
		return nil, err
	}
	//2.Check is these orders fit the requirement
	if qi.EnableStateQuery || qi.EnableBoughtDateQuery || qi.EnableTravelDateQuery {
		var finalList []Order
		for _, tmpOrder := range orders {
			statePassFlag := false
			boughtDatePassFlag := false
			travelDatePassFlag := false
			//3.Check order state requirement.
			if qi.EnableStateQuery {
				if tmpOrder.Status != qi.State {
					statePassFlag = false
				} else {
					statePassFlag = true
				}
			} else {
				statePassFlag = true
			}
			// 4.Check order travel date requirement.
			boughtDate, _ := time.Parse(time.DateTime, tmpOrder.BoughtDate)
			travelDate, _ := time.Parse(time.DateTime, tmpOrder.BoughtDate)
			travelDateEnd, _ := time.Parse(time.DateTime, qi.TravelDateEnd)
			boughDateEnd, _ := time.Parse(time.DateTime, qi.BoughtDateStart)
			boughDateStart, _ := time.Parse(time.DateTime, qi.BoughtDateEnd)
			if qi.EnableTravelDateQuery {
				if travelDate.Before(travelDateEnd) && travelDate.After(boughDateStart) {
					travelDatePassFlag = true
				} else {
					travelDatePassFlag = false
				}
			} else {
				travelDatePassFlag = true
			}
			// 5.Check order bought date requirement.
			if qi.EnableBoughtDateQuery {
				if boughtDate.Before(boughDateEnd) && boughtDate.After(boughDateStart) {
					boughtDatePassFlag = true
				} else {
					boughtDatePassFlag = false
				}
			} else {
				boughtDatePassFlag = true
			}
			// 6.Check if all requirement fits.
			if statePassFlag && boughtDatePassFlag && travelDatePassFlag {
				finalList = append(finalList, tmpOrder)
			}
		}
		return finalList, nil
	}
	return nil, nil
}

func (o *OrderServiceImpl) CalculateSoldTicket(ctx context.Context, travelDate time.Time, trainNumber string) (SoldTicket, error) {
	orders, err := o.findByTravelDateAndTrainNumber(ctx, travelDate, trainNumber)
	if err != nil {
		return SoldTicket{}, err
	}
	cstr := SoldTicket{
		TravelDate:  travelDate,
		TrainNumber: trainNumber,
	}
	for _, order := range orders {
		if order.Status == ORDER_STATUS_CHANGE {
			continue
		}
		if order.SeatClass == ORDER_SEAT_CLASS_NONE {
			cstr.NoSeat++
		} else if order.SeatClass == ORDER_SEAT_CLASS_BUSINESS {
			cstr.BusinessSeat++
		} else if order.SeatClass == ORDER_SEAT_CLASS_FIRSTCLASS {
			cstr.FirstClassSeat++
		} else if order.SeatClass == ORDER_SEAT_CLASS_SECONDCLASS {
			cstr.SecondClassSeat++
		} else if order.SeatClass == ORDER_SEAT_CLASS_HARDSEAT {
			cstr.HardSeat++
		} else if order.SeatClass == ORDER_SEAT_CLASS_SOFTSEAT {
			cstr.SoftSeat++
		} else if order.SeatClass == ORDER_SEAT_CLASS_HARDBED {
			cstr.HardBed++
		} else if order.SeatClass == ORDER_SEAT_CLASS_SOFTBED {
			cstr.SoftBed++
		} else if order.SeatClass == ORDER_SEAT_CLASS_HIGHSOFTBED {
			cstr.HighSoftBed++
		}
	}
	return cstr, nil
}

func (o *OrderServiceImpl) SecurityInfoCheck(ctx context.Context, dateFrom time.Time, accountId string) (OrderSecurity, error) {
	var result OrderSecurity
	orders, err := o.findByAccountId(ctx, accountId)
	if err != nil {
		return OrderSecurity{}, err
	}
	countOrderInOneHour := 0
	countTotalValidOrder := 0

	dateFrom = dateFrom.Add(-1 * time.Hour)

	for _, order := range orders {
		if order.Status == ORDER_STATUS_NOT_PAID || order.Status == ORDER_STATUS_PAID || order.Status == ORDER_STATUS_COLLECTED {
			countTotalValidOrder += 1
		}
		boughtDate, err := time.Parse(time.DateTime, order.BoughtDate)
		if err == nil && boughtDate.After(dateFrom) {
			countOrderInOneHour++
		}
	}

	result.orderNumInLastOneHour = countOrderInOneHour
	result.orderNumOfValidOrder = countTotalValidOrder
	return result, err
}

func (o *OrderServiceImpl) GetOrderPrice(ctx context.Context, orderId string) (string, error) {
	order, err := o.findById(ctx, orderId)
	if err != nil {
		return "", err
	}
	return order.Price, nil
}

func (o *OrderServiceImpl) PayOrder(ctx context.Context, orderId string) (bool, error) {
	order, err := o.findById(ctx, orderId)
	if err != nil {
		return false, err
	}
	order.Status = ORDER_STATUS_PAID
	err = o.update(ctx, order)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o *OrderServiceImpl) AddCreateNewOrder(ctx context.Context, order Order) error {
	accountOrders, err := o.findByAccountId(ctx, order.AccountID)
	if err != nil {
		return err
	}
	for _, accountOrder := range accountOrders {
		if accountOrder == order {
			return fmt.Errorf("order already exists")
		}
	}
	return o.save(ctx, order)
}

func (o *OrderServiceImpl) CreateNewOrder(ctx context.Context, order Order) (Order, error) {
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

func (o *OrderServiceImpl) UpdateOrder(ctx context.Context, order Order) error {
	return o.update(ctx, order)
}

func (o *OrderServiceImpl) SaveOrderInfo(ctx context.Context, order Order) (Order, error) {
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

	_, err = collection.Upsert(ctx, filter, order)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (o *OrderServiceImpl) GetOrderById(ctx context.Context, orderID string) (Order, error) {
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

func (o *OrderServiceImpl) FindAllOrder(ctx context.Context) ([]Order, error) {
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

func (o *OrderServiceImpl) DeleteOrder(ctx context.Context, id string) error {
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

func (o *OrderServiceImpl) findByTravelDateAndTrainNumber(ctx context.Context, travelDate time.Time, trainNumber string) ([]Order, error) {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "TravelDate", Value: travelDate}, {Key: "TrainNumber", Value: trainNumber}}
	cursor, err := collection.FindMany(ctx, filter)
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

func (o *OrderServiceImpl) findById(ctx context.Context, orderID string) (Order, error) {
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

func (o *OrderServiceImpl) findByAccountId(ctx context.Context, accountId string) ([]Order, error) {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "AccountID", Value: accountId}}
	cursor, err := collection.FindMany(ctx, filter)
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

func (o *OrderServiceImpl) update(ctx context.Context, order Order) error {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "ID", Value: order.ID}}
	updated, err := collection.Upsert(ctx, filter, order)
	if err != nil {
		return err
	}
	if updated {
		return fmt.Errorf("order not found for id (%s)", order.ID)
	}
	return nil
}

func (o *OrderServiceImpl) save(ctx context.Context, order Order) error {
	collection, err := o.orderDB.GetCollection(ctx, "order_db", "order")
	if err != nil {
		return err
	}

	return collection.InsertOne(ctx, order)
}

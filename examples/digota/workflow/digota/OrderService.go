package digota

import (
	"context"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	ns                         = "order"
	orderTTL                   = time.Minute * 2
	defaultTaxDescription      = "Tax"
	defaultDiscountDescription = "Discount"
	defaultShippingDescription = "Shipping"
)

type OrderService interface {
	//Run(ctx context.Context) error
	New(ctx context.Context, currency int32, items []*OrderItem, metadata map[string]string, email string, shipping *Shipping) (*Order, error)
	Get(ctx context.Context, id string) (*Order, error)
	List(ctx context.Context, page int64, limit int64, sort int32) (*OrderList, error)
	Pay(ctx context.Context, id string, card *Card, paymentProviderID int32) (*Order, error)
	Return(ctx context.Context, id string) (*Order, error)
}

type OrderServiceImpl struct {
	skuService SkuService
	db         backend.NoSQLDatabase
	/* queue      backend.Queue */
}

func NewOrderServiceImpl(ctx context.Context, skuService SkuService, db backend.NoSQLDatabase /* , queue backend.Queue */) (OrderService, error) {
	s := &OrderServiceImpl{skuService: skuService, db: db /* , queue: queue */}
	return s, nil
}

func (s *OrderServiceImpl) New(ctx context.Context, currency int32, items []*OrderItem, metadata map[string]string, email string, shipping *Shipping) (*Order, error) {
	order := &Order{
		Currency: currency,
		Items:    items,
		Metadata: metadata,
		Email:    email,
		Shipping: shipping,
	}

	// 1. get updated order items
	//var reqItems = items
	//var skuMap = make(map[string]*OrderItem)
	var orderItems []*OrderItem

	// 1.1. merge duplicated items
	/* for _, v := range reqItems {
		if v.IsTypeSku() {
			if skuItem, ok := skuMap[v.Parent]; ok {
				skuItem.Quantity += v.Quantity
				continue
			} else {
				skuMap[v.Parent] = v
			}
		} else if v.IsTypeDiscount() {
			if v.Quantity <= 0 {
				v.Quantity = 1
			}
		} else if v.IsTypeShipping() {
			if v.Quantity <= 0 {
				v.Quantity = 1
			}
		} else if v.IsTypeTax() {
			if v.Quantity <= 0 {
				v.Quantity = 1
			}
			orderItems = append(orderItems, v)
		}
	} */

	// 1.2. update order item data
	for _, myitem1 := range items {
		if myitem1.IsTypeTax() {
			if myitem1.Quantity <= 0 {
				myitem1.Quantity = 1
			}
			orderItems = append(orderItems, myitem1)
		}
	}
	for _, v := range orderItems {
		if v.IsTypeSku() {
			orderItem := v
			item, err := s.skuService.Get(ctx, v.Parent)
			if err != nil {
				return nil, err
			} else {
				orderItem.Amount = int64(item.Price)
				orderItem.Currency = item.Currency
				orderItem.Description = item.Name
			}
		} else if v.IsTypeDiscount() {
			// nothing to fetch yet
			if v.Description == "" {
				v.Description = defaultDiscountDescription
			}
		} else if v.IsTypeShipping() {
			// nothing to fetch yet
			if v.Description == "" {
				v.Description = defaultShippingDescription
			}
		} else if v.IsTypeTax() {
			// nothing to fetch yet
			if v.Description == "" {
				v.Description = defaultTaxDescription
			}
		}
	}

	// 2. calculate total and write to database

	order.Items = orderItems

	amount, err := calculateTotal(order.Currency, order.Items)
	if err != nil {
		return nil, err
	}
	order.Amount = amount

	collection, err := s.db.GetCollection(ctx, "orders_db", "orders")
	if err != nil {
		return nil, err
	}
	err = collection.InsertOne(ctx, order)
	return order, err
}

func (s *OrderServiceImpl) Get(ctx context.Context, id string) (*Order, error) {
	collection, err := s.db.GetCollection(ctx, "orders_db", "orders")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "Id", Value: id}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	var order *Order
	found, err := result.One(ctx, order)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("order not found for id (%s)", id)
	}

	return order, nil
}

func (s *OrderServiceImpl) List(ctx context.Context, page int64, limit int64, sort int32) (*OrderList, error) {
	collection, err := s.db.GetCollection(ctx, "orders_db", "orders")
	if err != nil {
		return nil, err
	}

	result, err := collection.FindMany(ctx, nil)
	if err != nil {
		return nil, err
	}

	var orders []*Order
	err = result.All(ctx, orders)
	if err != nil {
		return nil, err
	}

	orderList := &OrderList{
		Orders: orders,
		Total:  int32(len(orders)),
	}

	return orderList, nil
}

func (s *OrderServiceImpl) Pay(ctx context.Context, id string, card *Card, paymentProviderID int32) (*Order, error) {
	// TODO
	return nil, nil
}

func (s *OrderServiceImpl) Return(ctx context.Context, id string) (*Order, error) {
	// TODO
	return nil, nil
}

func calculateTotal(currency int32, orderItems []*OrderItem) (int64, error) {
	var err error
	currencyString := Currency_name[int32(currency)]
	m := money.New(0, currencyString)
	for _, v := range orderItems {
		if v.Quantity <= 0 {
			v.Quantity = 1
		}
		vCurrencyString := Currency_name[int32(v.Currency)]
		m, err = m.Add(money.New(v.Quantity*v.Amount, vCurrencyString))
		if err != nil {
			return 0, err
		}
	}
	return m.Amount(), nil
}

/* type QueueMessage struct {
	id string
}

func (s *OrderServiceImpl) Run(ctx context.Context) error {
	var message QueueMessage
	s.queue.Pop(ctx, &message)

	collection, err := s.db.GetCollection(ctx, "orders_db", "orders")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "Id", Value: message.id}}
	err = collection.DeleteOne(ctx, filter)
	return err
	return nil
} */

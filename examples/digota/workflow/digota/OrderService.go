package digota

import (
	"context"
	"sync"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

const (
	ns                         = "order"
	orderTTL                   = time.Minute * 2
	defaultTaxDescription      = "Tax"
	defaultDiscountDescription = "Discount"
	defaultShippingDescription = "Shipping"
)

type OrderService interface {
	New(ctx context.Context, currency int32, items []*OrderItem, metadata map[string]string, email string, shipping *Shipping) (*Order, error)
	/* Get(ctx context.Context, id string) (*Order, error)
	List(ctx context.Context, page int64, limit int64, sort int32) (*OrderList, error)
	Pay(ctx context.Context, id string, card *Card, paymentProviderID int32) (*Order, error)
	Return(ctx context.Context, id string) (*Order, error) */
}

type OrderServiceImpl struct {
	skuService SkuService
	db         backend.NoSQLDatabase
}

func NewOrderServiceImpl(ctx context.Context, skuService SkuService, db backend.NoSQLDatabase) (OrderService, error) {
	s := &OrderServiceImpl{skuService: skuService, db: db}
	return s, nil
}

/* func (s *OrderServiceImpl) getSkuObject(ctx context.Context, wg *sync.WaitGroup, orderItem *OrderItem) error {
	defer wg.Done()
	item, err := s.skuService.Get(ctx, orderItem.Parent)
	if err != nil {
		return err
	} else {
		orderItem.Amount = int64(item.Price)
		orderItem.Currency = item.Currency
		orderItem.Description = item.Name
	}
	return nil
} */

func (s *OrderServiceImpl) getUpdatedOrderItems(ctx context.Context, items []*OrderItem) ([]*OrderItem, error) {
	var skuMap = make(map[string]*OrderItem)
	var orderItems []*OrderItem
	//var mtx = sync.Mutex{}
	var errs []error
	var wg = sync.WaitGroup{}

	// get relevant order items
	for _, v := range items {
		if v.IsTypeSku() {
			if skuItem, ok := skuMap[v.Parent]; ok {
				skuItem.Quantity += v.Quantity
				continue
			} else {
				skuMap[v.Parent] = v
			}
			orderItems = append(orderItems, v)

		} else if v.IsTypeDiscount() {
			if v.Quantity <= 0 {
				v.Quantity = 1
			}
			orderItems = append(orderItems, v)

		} else if v.IsTypeShipping() {
			if v.Quantity <= 0 {
				v.Quantity = 1
			}
			orderItems = append(orderItems, v)

		} else if v.IsTypeTax() {
			if v.Quantity <= 0 {
				v.Quantity = 1
			}
			orderItems = append(orderItems, v)

		}
	}

	// update order item data
	for _, v := range orderItems {
		if v.IsTypeSku() {
			item, err := s.skuService.Get(ctx, v.Parent)
			if err != nil {
				return nil, err
			} else {
				v.Amount = int64(item.Price)
				v.Currency = item.Currency
				v.Description = item.Name
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
	wg.Wait()
	if errs != nil {
		return nil, errs[0]
	}

	return orderItems, nil
}

func (s *OrderServiceImpl) New(ctx context.Context, currency int32, items []*OrderItem, metadata map[string]string, email string, shipping *Shipping) (*Order, error) {
	order := &Order{
		Currency: currency,
		Items:    items,
		Metadata: metadata,
		Email:    email,
		Shipping: shipping,
	}

	orderItems, err := s.getUpdatedOrderItems(ctx, items)
	if err != nil {
		return nil, err
	}
	order.Items = orderItems

	amount, err := calculateTotal(order.Currency, orderItems)
	if err != nil {
		return nil, err
	}
	order.Amount = amount

	collection, err := s.db.GetCollection(ctx, "orders", "orders")
	if err != nil {
		return nil, err
	}
	err = collection.InsertOne(ctx, *order)
	return order, err
}

/* func (s *OrderServiceImpl) Get(ctx context.Context, id string) (*Order, error) {
	collection, err := s.db.GetCollection(ctx, "orders", "orders")
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "id", Value: id}}
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
	collection, err := s.db.GetCollection(ctx, "orders", "orders")
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
	return nil, nil
}

func (s *OrderServiceImpl) Return(ctx context.Context, id string) (*Order, error) {
	return nil, nil
}
*/

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

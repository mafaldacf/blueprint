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
	Run(ctx context.Context) error
	/* Get(ctx context.Context, id string) (*Order, error)
	List(ctx context.Context, page int64, limit int64, sort int32) (*OrderList, error)
	Pay(ctx context.Context, id string, card *Card, paymentProviderID int32) (*Order, error)
	Return(ctx context.Context, id string) (*Order, error) */
}

type OrderServiceImpl struct {
	skuService SkuService
	db         backend.NoSQLDatabase
	queue      backend.Queue
}

func NewOrderServiceImpl(ctx context.Context, skuService SkuService, db backend.NoSQLDatabase, queue backend.Queue) (OrderService, error) {
	s := &OrderServiceImpl{skuService: skuService, db: db, queue: queue}
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
	//var skuMap = make(map[string]*OrderItem)
	var orderItems []*OrderItem
	//var mtx = sync.Mutex{}
	var errs []error
	var wg = sync.WaitGroup{}

	// get relevant order items
	for _, myitem1 := range items {
		/* if myitem1.IsTypeSku() {
			if skuItem, ok := skuMap[myitem1.Parent]; ok {
				skuItem.Quantity += myitem1.Quantity
				continue
			} else {
				skuMap[myitem1.Parent] = myitem1
			}
			orderItems = append(orderItems, myitem1)

		} else if myitem1.IsTypeDiscount() {
			if myitem1.Quantity <= 0 {
				myitem1.Quantity = 1
			}
			orderItems = append(orderItems, myitem1)

		} else if myitem1.IsTypeShipping() {
			if myitem1.Quantity <= 0 {
				myitem1.Quantity = 1
			}
			orderItems = append(orderItems, myitem1)

		} else */if myitem1.IsTypeTax() {
			if myitem1.Quantity <= 0 {
				myitem1.Quantity = 1
			}
			orderItems = append(orderItems, myitem1)
		}
	}

	// update order item data
	for _, myitem2 := range orderItems {
		if myitem2.IsTypeSku() {
			item, err := s.skuService.Get(ctx, myitem2.Parent)
			if err != nil {
				return nil, err
			} else {
				myitem2.Amount = int64(item.Price)
				myitem2.Currency = item.Currency
				myitem2.Description = item.Name
			}
		} /* else if myitem2.IsTypeDiscount() {
			// nothing to fetch yet
			if myitem2.Description == "" {
				myitem2.Description = defaultDiscountDescription
			}
		} else if myitem2.IsTypeShipping() {
			// nothing to fetch yet
			if myitem2.Description == "" {
				myitem2.Description = defaultShippingDescription
			}
		} else if myitem2.IsTypeTax() {
			// nothing to fetch yet
			if myitem2.Description == "" {
				myitem2.Description = defaultTaxDescription
			}
		} */
	}
	wg.Wait()
	if errs != nil {
		return nil, errs[0]
	}

	return orderItems, nil
}

/* func (s *OrderServiceImpl) New2(ctx context.Context, currency int32, items []*OrderItem, amount int64, shipping *Shipping) (*Order, error) {
	order := &Order{
		Currency: currency,
		Items:    items,
		Shipping: shipping,
	}

	if amount > 10 {
		order.Currency = 2
		order.Shipping2 = Shipping{Name: "myname1"}
		order.Amount = amount
	} else {
		order.Currency = 3
		order.Shipping2 = Shipping{Name: "myname2"}
	}

	collection, _ := s.db.GetCollection(ctx, "orders", "orders")
	collection.InsertOne(ctx, order.Shipping2)


	return order, nil
} */

func (s *OrderServiceImpl) New2(ctx context.Context, items []*OrderItem, shipping1 *Shipping, shipping2 Shipping) (*Order, error) {
	order := &Order{
		Items: items,
	}

	for _, myitem1 := range items {
		if myitem1.Quantity <= 0 {
			shipping1.Name = "myname1"
		} else {
			shipping2.Name = "myname2"
		}
	}

	shipping2.Carrier = "mycarrier"
	order.Shipping = shipping1
	order.Shipping2 = shipping2

	collection, err := s.db.GetCollection(ctx, "orders_db", "orders")
	if err != nil {
		return nil, err
	}
	collection.InsertOne(ctx, order)

	return order, nil
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

	for _, myitem1 := range items {
		if myitem1.IsTypeTax() {
			if myitem1.Quantity <= 0 {
				myitem1.Quantity = 1
			}
			orderItems = append(orderItems, myitem1)
		}
	}
	for _, myitem2 := range orderItems {
		if myitem2.IsTypeSku() {
			item, err := s.skuService.Get(ctx, myitem2.Parent)
			if err != nil {
				return nil, err
			} else {
				myitem2.Amount = int64(item.Price)
				myitem2.Currency = item.Currency
				myitem2.Description = item.Name
			}
		} else if myitem2.IsTypeDiscount() {
			// nothing to fetch yet
			if myitem2.Description == "" {
				myitem2.Description = defaultDiscountDescription
			}
		} else if myitem2.IsTypeShipping() {
			// nothing to fetch yet
			if myitem2.Description == "" {
				myitem2.Description = defaultShippingDescription
			}
		} else if myitem2.IsTypeTax() {
			// nothing to fetch yet
			if myitem2.Description == "" {
				myitem2.Description = defaultTaxDescription
			}
		}
	}
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
	err = collection.InsertOne(ctx, *order)
	order.Items[0].Amount = 100
	order.Shipping.Address.City = "myaddress"
	return order, err
}

/* func (s *OrderServiceImpl) Get(ctx context.Context, id string) (*Order, error) {
	collection, err := s.db.GetCollection(ctx, "orders_db", "orders")
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

type QueueMessage struct {
	id string
}

func (s *OrderServiceImpl) Run(ctx context.Context) error {
	/* var message QueueMessage
	s.queue.Pop(ctx, &message)

	collection, err := s.db.GetCollection(ctx, "orders_db", "orders")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "id", Value: message.id}}
	err = collection.DeleteOne(ctx, filter)
	return err */
	return nil
}

package digota

import (
	"fmt"
	"time"
)

type Order struct {
	Id        string            `json:"id,omitempty" bson:"_id"`
	Amount    int64             `json:"amount,omitempty"`
	Currency  int32             `json:"currency,omitempty"`
	Items     []*OrderItem      `json:"items,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	Email     string            `json:"email,omitempty"`
	ChargeId  string            `json:"chargeId,omitempty"`
	Status    int32             `json:"Status,omitempty"`
	Shipping  *Shipping         `json:"shipping,omitempty"`
	Shipping2 Shipping          `json:"shipping2,omitempty"`
	Created   int64             `json:"created,omitempty"`
	Updated   int64             `json:"updated,omitempty"`
}

type OrderItem struct {
	Type        int32  `json:"type,omitempty"`
	Quantity    int64  `json:"quantity,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Currency    int32  `json:"currency,omitempty"`
	Parent      string `json:"parent,omitempty"`
	Description string `json:"description,omitempty"`
}

func (item *OrderItem) IsTypeReserved() bool { return item.Type == 0 }
func (item *OrderItem) IsTypeSku() bool      { return item.Type == 1 }
func (item *OrderItem) IsTypeDiscount() bool { return item.Type == 2 }
func (item *OrderItem) IsTypeTax() bool      { return item.Type == 3 }
func (item *OrderItem) IsTypeShipping() bool { return item.Type == 4 }
func CurrencyIsReserved(currency int32) bool { return currency == 0 }

func (item *OrderItem) GetType() int32 {
	return item.Type
}

func (order *Order) GetCurrency() int32 {
	return order.Currency
}

func (order *Order) GetChargeId() string {
	return order.ChargeId
}

func (order *Order) GetItems() []*OrderItem {
	return order.Items
}

func (order *Order) GetAmount() int64 {
	return order.Amount
}


func (o *Order) IsReturnable(amount int64) error {
	if o.Status != int32(Order_Paid) && o.Status != int32(Order_Fulfilled) && o.Status != int32(Order_Canceled) {
		return fmt.Errorf("Order is not paid or fulfilled.")
	}
	// if refund amount is bigger than the order amount return err
	if amount > o.GetAmount() {
		return fmt.Errorf("Refund amount is greater then order amount.")
	}
	return nil
}

func (o *Order) IsPayable() error {
	if o.Status != int32(Order_Created) {
		return fmt.Errorf("Order is not in created status.")
	}
	if time.Since(time.Unix(o.Created, 0)) > orderTTL {
		return fmt.Errorf("Order is too old for paying.")
	}
	if o.GetAmount() <= 0 {
		return fmt.Errorf("Order amount is Zero.")
	}
	return nil
}

/* type OrderItem_Type int32

const (
	OrderItem_reserved OrderItem_Type = 0
	OrderItem_sku      OrderItem_Type = 1
	OrderItem_discount OrderItem_Type = 2
	OrderItem_tax      OrderItem_Type = 3
	OrderItem_shipping OrderItem_Type = 4
) */

var OrderItem_Type_name = map[int32]string{
	0: "reserved",
	1: "sku",
	2: "discount",
	3: "tax",
	4: "shipping",
}
var OrderItem_Type_value = map[string]int32{
	"reserved": 0,
	"sku":      1,
	"discount": 2,
	"tax":      3,
	"shipping": 4,
}

type OrderStatus int32

const (
	Order_Created   OrderStatus = 0
	Order_Paid      OrderStatus = 1
	Order_Canceled  OrderStatus = 2
	Order_Fulfilled OrderStatus = 3
	Order_Returned  OrderStatus = 4
)

/* var OrderStatus_name = map[int32]string{
	0: "Created",
	1: "Paid",
	2: "Canceled",
	3: "Fulfilled",
	4: "Returned",
}
var OrderStatus_value = map[string]int32{
	"Created":   0,
	"Paid":      1,
	"Canceled":  2,
	"Fulfilled": 3,
	"Returned":  4,
} */

type Shipping struct {
	Name           string            `json:"name,omitempty"`
	Phone          string            `json:"phone,omitempty"`
	Address        *Shipping_Address `json:"address,omitempty"`
	Carrier        string            `json:"carrier,omitempty"`
	TrackingNumber string            `json:"trackingNumber,omitempty"`
}

type Shipping_Address struct {
	Line1      string `json:"line1,omitempty"`
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	Line2      string `json:"line2,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	State      string `json:"state,omitempty"`
}

type OrderList struct {
	Orders []*Order `json:"orders,omitempty"`
	Total  int32    `json:"total,omitempty"`
}

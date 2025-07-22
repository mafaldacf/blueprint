package digota

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

type Test struct {
	Value int
}

type OrderItem struct {
	Type        int32  `json:"type,omitempty" validate:"required,gte=1,lte=4"`
	Quantity    int64  `json:"quantity,omitempty" validate:"omitempty,gte=0"`
	Amount      int64  `json:"amount,omitempty"`
	Currency    int32  `json:"currency,omitempty" validate:"omitempty,gte=1,lte=128"`
	Parent      string `json:"parent,omitempty" validate:"omitempty,uuid4"`
	Description string `json:"description,omitempty"`
	Test        *Test  `json:"test ,omitempty"`
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

/* type OrderStatus int32

const (
	Order_Created   OrderStatus = 0
	Order_Paid      OrderStatus = 1
	Order_Canceled  OrderStatus = 2
	Order_Fulfilled OrderStatus = 3
	Order_Returned  OrderStatus = 4
) */

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

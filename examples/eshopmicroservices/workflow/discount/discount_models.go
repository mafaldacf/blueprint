package discount

type Coupon struct {
	Id          int     `bson:"Id"`
	ProductName string  `bson:"ProductName"`
	Description string  `bson:"Description"`
	Amount      float64 `bson:"Amount"`
}

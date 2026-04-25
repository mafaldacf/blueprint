package trainticket

type ConsignRequest struct {
	ID         string
	OrderID    string
	AccountID  string
	HandleDate string
	TargetDate string
	From       string
	To         string
	Consignee  string
	Phone      string
	Weight     float64
	IsWithin   bool
}

type ConsignRecord struct {
	ID         string  `bson:"ID"`
	OrderID    string  `bson:"OrderID"`
	AccountID  string  `bson:"AccountID"`
	HandleDate string  `bson:"HandleDate"`
	TargetDate string  `bson:"TargetDate"`
	FromPlace  string  `bson:"FromPlace"`
	ToPlace    string  `bson:"ToPlace"`
	Consignee  string  `bson:"Consignee"`
	Phone      string  `bson:"Phone"`
	Price      float64 `bson:"Price"`
	Weight     float64 `bson:"Weight"`
}

package train_ticket2

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
	ID         string
	OrderID    string
	AccountID  string
	HandleDate string
	TargetDate string
	FromPlace  string
	ToPlace    string
	Consignee  string
	Phone      string
	Price      float64
	Weight     float64
}

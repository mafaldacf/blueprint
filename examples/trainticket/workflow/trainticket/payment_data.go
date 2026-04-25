package trainticket

const PaymentType_P = "P"
const PaymentType_D = "D"
const PaymentType_O = "O"
const PaymentType_E = "E"

type Payment struct {
	ID          string `bson:"ID"`
	OrderID     string `bson:"OrderID"`
	UserID      string `bson:"UserID"`
	Price       string `bson:"Price"`
	PaymentType string `bson:"PaymentType"`
}

type Money struct {
	ID     string `bson:"ID"`
	UserID string `bson:"UserID"`
	Money  string `bson:"Money"`
}

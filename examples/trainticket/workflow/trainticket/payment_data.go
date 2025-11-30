package trainticket

type PaymentType string

const (
	PaymentType_P PaymentType = "P"
	PaymentType_D PaymentType = "D"
	PaymentType_O PaymentType = "O"
	PaymentType_E PaymentType = "E"
)

type Payment struct {
	ID          string
	OrderID     string
	UserID      string
	Price       string
	PaymentType PaymentType
}

type Money struct {
	ID     string
	UserID string
	Money  string
}

package trainticket

type InsideMoney struct {
	UserID string
	Money  string
	Type   MoneyType
}

const (
	INSIDE_MONEY_TYPE_ADD      = iota // 0
	INSIDE_MONEY_TYPE_DRAWBACK        // 1
)

type PaymentInfo struct {
	UserId  string
	OrderId string
	TripId  string
	Price   string
}

type MoneyType string

const (
	MoneyType_A MoneyType = "A"
)

type AccountInfo struct {
	UserId string
	Money  string
}

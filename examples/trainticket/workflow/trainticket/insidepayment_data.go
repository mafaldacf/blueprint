package trainticket

type InsideMoney struct {
	UserID string `bson:"UserID"`
	Money  string `bson:"Money"`
	Type   string `bson:"Type"`
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

const MoneyType_A = "A"
type AccountInfo struct {
	UserId string
	Money  string
}

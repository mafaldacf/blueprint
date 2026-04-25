package trainticket

const (
	ORDER_STATUS_NOT_PAID  = iota // 0
	ORDER_STATUS_PAID             // 1
	ORDER_STATUS_CHANGE           // 2
	ORDER_STATUS_COLLECTED        // 3
	ORDER_STATUS_USED             // 4
	ORDER_STATUS_CANCELED         // 5
)

const (
	ORDER_SEAT_CLASS_NONE        = iota // 0
	ORDER_SEAT_CLASS_BUSINESS           // 1
	ORDER_SEAT_CLASS_FIRSTCLASS         // 2
	ORDER_SEAT_CLASS_SECONDCLASS        // 3
	ORDER_SEAT_CLASS_HARDSEAT           // 4
	ORDER_SEAT_CLASS_SOFTSEAT           // 5
	ORDER_SEAT_CLASS_HARDBED            // 6
	ORDER_SEAT_CLASS_SOFTBED            // 7
	ORDER_SEAT_CLASS_HIGHSOFTBED        // 8
)

type Order struct {
	ID                     string `bson:"ID"`
	BoughtDate             string `bson:"BoughtDate"`
	TravelDate             string `bson:"TravelDate"`
	TravelTime             string `bson:"TravelTime"`
	AccountID              string `bson:"AccountID"`
	ContactsName           string `bson:"ContactsName"`
	DocumentType           int    `bson:"DocumentType"`
	ContactsDocumentNumber string `bson:"ContactsDocumentNumber"`
	TrainNumber            string `bson:"TrainNumber"`
	CoachNumber            int    `bson:"CoachNumber"`
	SeatClass              int    `bson:"SeatClass"`
	SeatNumber             string `bson:"SeatNumber"`
	FromStation            string `bson:"FromStation"`
	ToStation              string `bson:"ToStation"`
	Status                 int    `bson:"Status"`
	Price                  string `bson:"Price"`
}

type LeftTicketInfo struct {
	SoldTickets []Ticket
}

type Ticket struct {
	SeatNo       int
	StartStation string
	DestStation  string
}

type OrderSecurity struct {
	orderNumInLastOneHour int
	orderNumOfValidOrder  int
}

type SoldTicket struct {
	TravelDate      string
	TrainNumber     string
	NoSeat          int
	BusinessSeat    int
	FirstClassSeat  int
	SecondClassSeat int
	HardSeat        int
	SoftSeat        int
	HardBed         int
	SoftBed         int
	HighSoftBed     int
}

type OrderInfo struct {
	LoginId               string
	TravelDateStart       string
	TravelDateEnd         string
	BoughtDateStart       string
	BoughtDateEnd         string
	State                 int
	EnableTravelDateQuery bool
	EnableBoughtDateQuery bool
	EnableStateQuery      bool
}

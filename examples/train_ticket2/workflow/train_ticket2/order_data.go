package train_ticket2

const (
	ORDER_STATUS_NOT_PAID  = iota // 0
	ORDER_STATUS_PAID             // 1
	ORDER_STATUS_CHANGE           // 2
	ORDER_STATUS_COLLECTED        // 3
	ORDER_STATUS_USED             // 4
	ORDER_STATUS_CANCELED         // 5
)

type Order struct {
	ID                     string
	BoughtDate             string
	TravelDate             string
	TravelTime             string
	AccountID              string
	ContactsName           string
	DocumentType           int
	ContactsDocumentNumber string
	TrainNumber            string
	CoachNumber            int
	SeatClass              int
	SeatNumber             string
	FromStation            string
	ToStation              string
	Status                 int
	Price                  string
}

type LeftTicketInfo struct {
	SoldTickets []Ticket
}

type Ticket struct {
	SeatNo       int
	StartStation string
	DestStation  string
}

package train_ticket2

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
	Seatnumber             string
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

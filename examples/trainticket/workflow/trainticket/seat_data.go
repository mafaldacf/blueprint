package trainticket

type SeatRequest struct {
	TravelDate   string
	TrainNumber  string
	StartStation string
	DestStation  string
	SeatType     int
	TotalNum     int
	Stations     []string
}

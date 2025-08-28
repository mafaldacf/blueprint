package train_ticket2

type Travel struct {
	Trip          Trip
	StartPlace    string
	EndPlace      string
	DepartureTime string
}

type TravelResult struct {
	Status    bool
	Percent   float64
	TrainType TrainType
	Route     Route
	Prices    map[string]string
}
type TravelInfo struct {
	LoginID             string
	TripID              string
	TrainTypeName       string
	RouteID             string
	StartStationName    string
	TerminalStationName string
	StationsName        string
	StartTime           string
	EndTime             string
}

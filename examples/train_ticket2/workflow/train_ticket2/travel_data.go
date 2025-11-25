package train_ticket2

type Trip struct {
	ID                  string
	TripID              string
	TrainTypeName       string
	RouteID             string
	StartStationName    string
	StationsName        string
	TerminalStationName string
	StartTime           string
	EndTime             string
}

type AdminTrip struct {
	Trip      Trip
	TrainType TrainType
	Route     Route
}

type TripAllDetailInfo struct {
	TripID     string
	TravelDate string
	From       string
	To         string
}

type TripResponse struct {
	TripID               string
	TrainTypeName        string
	StartStation         string
	TerminalStation      string
	StartTime            string
	EndTime              string
	EconomyClass         int
	ConfortClass         int
	PriceForEconomyClass string
	PriceForConfortClass string
}

type TripAllDetail struct {
	trip         Trip
	tripResponse TripResponse
}

type TripInfo struct {
	startPlace    string
	endPlace      string
	departureTime string
}

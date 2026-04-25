package trainticket

type Trip struct {
	ID                  string `bson:"ID"`
	TripID              string `bson:"TripID"`
	TrainTypeName       string `bson:"TrainTypeName"`
	RouteID             string `bson:"RouteID"`
	StartStationName    string `bson:"StartStationName"`
	StationsName        string `bson:"StationsName"`
	TerminalStationName string `bson:"TerminalStationName"`
	StartTime           string `bson:"StartTime"`
	EndTime             string `bson:"EndTime"`
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

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

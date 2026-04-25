package trainticket

type Route struct {
	ID           string   `bson:"ID"`
	Stations     []string `bson:"Stations"`
	Distances    []int64  `bson:"Distances"`
	StartStation string   `bson:"StartStation"`
	EndStation   string   `bson:"EndStation"`
}

type RouteInfo struct {
	ID           string `bson:"ID"`
	StartStation string `bson:"StartStation"`
	EndStation   string `bson:"EndStation"`
	StationList  string `bson:"StationList"`
	DistanceList string `bson:"DistanceList"`
}

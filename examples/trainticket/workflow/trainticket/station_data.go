package trainticket

type Station struct {
	ID       string `bson:"ID"`
	Name     string `bson:"Name"`
	StayTime int64  `bson:"StayTime"`
}

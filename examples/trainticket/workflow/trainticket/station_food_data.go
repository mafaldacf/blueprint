package trainticket

type StationFoodStore struct {
	ID           string  `bson:"ID"`
	StationName  string  `bson:"StationName"`
	StoreName    string  `bson:"StoreName"`
	Telephone    string  `bson:"Telephone"`
	BusinessTime string  `bson:"BusinessTime"`
	DeliveryFee  float64 `bson:"DeliveryFee"`
	Foods        Food    `bson:"Foods"`
}

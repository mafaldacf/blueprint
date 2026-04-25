package trainticket

type TrainFood struct {
	ID     string `bson:"ID"`
	TripID string `bson:"TripID"`
	Foods  []Food `bson:"Foods"`
}

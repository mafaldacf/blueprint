package shipping

type Shipment struct {
	ID     string `bson:"ID"`
	Name   string `bson:"Name"`
	Status string `bson:"status"`
}

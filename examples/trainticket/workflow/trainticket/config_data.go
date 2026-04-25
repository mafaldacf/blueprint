package trainticket

type Config struct {
	Name        string `bson:"Name"`
	Value       string `bson:"Value"`
	Description string `bson:"Description"`
}

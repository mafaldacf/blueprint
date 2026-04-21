package catalog

type Product struct {
	Id          string `bson:"Id"`
	Name        string    `bson:"Name"`
	Category    []string  `bson:"Category"`
	Description string    `bson:"Description"`
	ImageFile   string    `bson:"ImageFile"`
	Price       float64   `bson:"Price"`
}

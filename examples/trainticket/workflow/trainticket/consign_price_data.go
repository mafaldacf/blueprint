package trainticket

type ConsignPrice struct {
	ID            string  `bson:"ID"`
	Index         int64   `bson:"Index"`
	InitialWeight float64 `bson:"InitialWeight"`
	InitialPrice  float64 `bson:"InitialPrice"`
	WithinPrice   float64 `bson:"WithinPrice"`
	BeyondPrice   float64 `bson:"BeyondPrice"`
}

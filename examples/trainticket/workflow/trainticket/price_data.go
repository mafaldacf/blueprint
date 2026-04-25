package trainticket

type PriceConfig struct {
	ID                  string  `bson:"ID"`
	TrainType           string  `bson:"TrainType"`
	RouteID             string  `bson:"RouteID"`
	BasicPriceRate      float64 `bson:"BasicPriceRate"`
	FirstClassPriceRate float64 `bson:"FirstClassPriceRate"`
}

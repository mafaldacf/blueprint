package trainticket

type PriceConfig struct {
	ID                  string
	TrainType           string
	RouteID             string
	BasicPriceRate      float64
	FirstClassPriceRate float64
}

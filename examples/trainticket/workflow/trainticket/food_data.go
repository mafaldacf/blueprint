package trainticket

type Food struct {
	Name  string  `bson:"Name"`
	Price float64 `bson:"Price"`
}

type FoodOrder struct {
	ID          string  `bson:"ID"`
	OrderID     string  `bson:"OrderID"`
	FoodType    int     `bson:"FoodType"`
	StationName string  `bson:"StationName"`
	StoreName   string  `bson:"StoreName"`
	FoodName    string  `bson:"FoodName"`
	Price       float64 `bson:"Price"`
}

package trainticket

type Food struct {
	Name  string
	Price float64
}

type FoodOrder struct {
	ID          string
	OrderID     string
	FoodType    int
	StationName string
	StoreName   string
	FoodName    string
	Price       float64
}

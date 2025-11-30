package trainticket

type Delivery struct {
	ID          string
	OrderID     string
	FoodName    string
	StoreName   string
	StationName string
}

func (d *Delivery) getId() string {
	return d.ID
}

func (d *Delivery) setId(id string) {
	d.ID = id
}

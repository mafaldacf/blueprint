package sockshop3

type cart struct {
    ID    string
    Items []Item
}

// A cart item is just an item ID and a quantity.  The catalogue service is responsible
// for managing the actual items.
type Item struct {
    ID        string  // Item ID will correspond to the ID used by the catalogue service
    Quantity  int     // The quantity of this item in the car
    UnitPrice float32 // The price of the item
}

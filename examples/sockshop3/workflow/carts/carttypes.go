package carts

type Cart struct {
    ID    string
    Items []Item
}

type Item struct {
    ID        string
    Quantity  int
    UnitPrice float32
}

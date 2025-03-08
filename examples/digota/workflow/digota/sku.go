package digota

type Sku struct {
	Id                string
	Name              string
	Price             uint64
	Currency          int32
	Active            bool
	Parent            string
	Metadata          map[string]string
	Attributes        map[string]string
	Image             string
	PackageDimensions *PackageDimensions
	Inventory         *Inventory
	Created           int64
	Updated           int64
}

type PackageDimensions struct {
	Height float64
	Length float64
	Weight float64
	Width  float64
}

type Inventory struct {
	Quantity int64
	Type     int32
}

/* type Inventory_Type int32

const (
	Inventory_Infinite Inventory_Type = 0
	Inventory_Finite   Inventory_Type = 1
) */

var Inventory_Type_name = map[int32]string{
	0: "Infinite",
	1: "Finite",
}
var Inventory_Type_value = map[string]int32{
	"Infinite": 0,
	"Finite":   1,
}

type SkuList struct {
	Orders []*Sku `json:"orders,omitempty"`
	Total  int32  `json:"total,omitempty"`
}

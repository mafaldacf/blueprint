package digota

type Sku struct {
	Id                string             `json:"id,omitempty" bson:"Id"`
	Name              string             `json:"name,omitempty" bson:"Name"`
	Price             uint64             `json:"price,omitempty" bson:"Price"`
	Currency          int32              `json:"currency,omitempty" bson:"Currency"`
	Active            bool               `json:"active,omitempty" bson:"Active"`
	Parent            string             `json:"parent,omitempty" bson:"Parent"`
	Metadata          map[string]string  `json:"metadata,omitempty" bson:"Metadata"`
	Attributes        map[string]string  `json:"attributes,omitempty" bson:"Attributes"`
	Image             string             `json:"image,omitempty" bson:"Image"`
	PackageDimensions *PackageDimensions `json:"packageDimensions,omitempty" bson:"PackageDimensions"`
	Inventory         *Inventory         `json:"inventory,omitempty" bson:"Inventory"`
	Created           int64              `json:"created,omitempty" bson:"Created"`
	Updated           int64              `json:"updated,omitempty" bson:"Updated"`
}

type PackageDimensions struct {
	Height float64 `json:"height,omitempty" bson:"Height"`
	Length float64 `json:"length,omitempty" bson:"Length"`
	Weight float64 `json:"weight,omitempty" bson:"Weight"`
	Width  float64 `json:"width,omitempty" bson:"Width"`
}

type Inventory struct {
	Quantity int64 `json:"quantity,omitempty" bson:"Quantity"`
	Type     int32 `json:"type,omitempty" bson:"Type"`
}

type Inventory_Type int32

const (
	Inventory_Infinite Inventory_Type = 0
	Inventory_Finite   Inventory_Type = 1
)

var Inventory_Type_name = map[int32]string{
	0: "Infinite",
	1: "Finite",
}
var Inventory_Type_value = map[string]int32{
	"Infinite": 0,
	"Finite":   1,
}

type SkuList struct {
	Orders []Sku `json:"orders,omitempty"`
	Total  int32 `json:"total,omitempty"`
}

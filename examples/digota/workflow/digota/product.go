package digota

type Product struct {
	Id          string            `json:"id,omitempty" bson:"Id"`
	Name        string            `json:"name,omitempty" bson:"Name"`
	Active      bool              `json:"active,omitempty" bson:"Active"`
	Attributes  []string          `json:"attributes,omitempty" bson:"Attributes"`
	Description string            `json:"description,omitempty" bson:"Description"`
	Images      []string          `json:"images,omitempty" bson:"Images"`
	Metadata    map[string]string `json:"metadata,omitempty" bson:"Metadata"`
	Shippable   bool              `json:"shippable,omitempty" bson:"Shippable"`
	ProductUrl  string            `json:"product_url,omitempty" bson:"ProductUrl"`
	Skus        []*Sku            `json:"skus,omitempty" bson:"Skus"`
	Created     int64             `json:"created,omitempty" bson:"Created"`
	Updated     int64             `json:"updated,omitempty" bson:"Updated"`
}

type ProductList struct {
	Products []Product
	Total    int32
}

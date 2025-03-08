package digota

type Product struct {
	Id          string            `json:"id,omitempty" bson:"_id"`
	Name        string            `json:"name,omitempty"`
	Active      bool              `json:"active,omitempty"`
	Attributes  []string          `json:"attributes,omitempty"`
	Description string            `json:"description,omitempty"`
	Images      []string          `json:"images,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Shippable   bool              `json:"shippable,omitempty"`
	Url         string            `json:"url,omitempty"`
	Skus        []*Sku            `json:"skus,omitempty"`
	Created     int64             `json:"created,omitempty"`
	Updated     int64             `json:"updated,omitempty"`
}

type ProductList struct {
	Products []*Product `json:"products,omitempty"`
	Total    int32      `json:"total,omitempty"`
}

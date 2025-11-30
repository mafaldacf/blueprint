package catalogue

type Sock struct {
	SockID      string   `json:"id" db:"sock_id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	ImageURL    []string `json:"imageUrl" db:"-"`
	ImageURL1   string   `json:"-" db:"image_url_1"`
	ImageURL2   string   `json:"-" db:"image_url_2"`
	Price       float32  `json:"price" db:"price"`
	Quantity    int      `json:"quantity" db:"quantity"`
	Tags        []string `json:"tag" db:"-"`
	TagString   string   `json:"-" db:"tag_name"`
}

type tag struct {
	TagID int    `db:"tag_id"`
	Name  string `db:"name"`
}

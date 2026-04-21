package catalogue

type Sock struct {
	SockID      string   `json:"id" db:"SockID"`
	Name        string   `json:"name" db:"Name"`
	Description string   `json:"description" db:"Description"`
	ImageURL    []string `json:"imageUrl" db:"-"`
	ImageURL1   string   `json:"-" db:"ImageURL1"`
	ImageURL2   string   `json:"-" db:"ImageURL2"`
	Price       float32  `json:"price" db:"Price"`
	Quantity    int      `json:"quantity" db:"Quantity"`
	Tags        []string `json:"tag" db:"-"`
	TagString   string   `json:"-" db:"tag_name"`
}

type tag struct {
	TagID int    `db:"TagID"`
	Name  string `db:"Name"`
}

package shopping_simple

type User struct {
	ID   int64
	Name string
}

type Media struct {
	Video []byte
}

type Post struct {
	User  User
	Media Media
	Text  string
}

var MyPost = Post{
	User: User{Name: "Alice"},
	Media: Media{},
	Text: "hello world",
}

type Product struct {
	ProductID    string
	Description  string
	PricePerUnit int
	Category     string
}

type Cart struct {
	CartID        string
	LastProductID string
	TotalQuantity int
	Products      []string
}

type CartProduct struct {
	CartID       string
	ProductID    string
	Quantity     int
	PricePerUnit int
}

type ProductQueueMessage struct {
	ProductID string
	Remove    bool
}

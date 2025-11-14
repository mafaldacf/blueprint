package sockshop3

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	Addresses Address
	Cards     Card
	UserID    string
	Salt      string
}
type Address struct {
	Street   string
	Number   string
	Country  string
	City     string
	PostCode string
	ID       string
}
type Card struct {
	LongNum string
	Expires string
	CCV     string
	ID      string
}

type dbCard struct {
	Card `bson:",inline"`
	ID   primitive.ObjectID `bson:"_id"`
}

package trainticket

type User struct {
	UserID       string `bson:"UserID"`
	Username     string `bson:"Username"`
	Password     string `bson:"Password"`
	Gender       int64  `bson:"Gender"`
	DocumentType int64  `bson:"DocumentType"`
	DocumentNum  string `bson:"DocumentNum"`
	Email        string `bson:"Email"`
}

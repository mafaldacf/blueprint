package trainticket

// DocumentType enum
const (
	NULL    int64 = iota
	ID_CARD
	PASSPORT
	OTHER
)

type Contact struct {
	ID             string `bson:"ID"`
	AccountID      string `bson:"AccountID"`
	Name           string `bson:"Name"`
	DocumentType   int    `bson:"DocumentType"`
	DocumentNumber string `bson:"DocumentNumber"`
	PhoneNumber    string `bson:"PhoneNumber"`
}

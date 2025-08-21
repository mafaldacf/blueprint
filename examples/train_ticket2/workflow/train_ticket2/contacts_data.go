package train_ticket2

// DocumentType enum
const (
	NULL int64 = iota
	ID_CARD
	PASSPORT
	OTHER
)

type Contact struct {
	ID             string
	AccountID      string
	Name           string
	DocumentType   int
	DocumentNumber string
	PhoneNumber    string
}

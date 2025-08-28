package train_ticket2

type InsideMoney struct {
	UserID string
	Money  string
	Type   int
}

const (
	INSIDE_MONEY_TYPE_ADD      = iota // 0
	INSIDE_MONEY_TYPE_DRAWBACK        // 1
)

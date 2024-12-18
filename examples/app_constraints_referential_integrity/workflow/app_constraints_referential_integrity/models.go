package app_constraints_referential_integrity

type Account struct {
	ReqID     string
	AccountID string
	Username  string // main user
	Timestamp int64
}

type AccountUsers struct { // secondary users
	AccountID string
	Usernames []string
}

type User struct {
	ReqID     string
	UserID    string
	Username  string
	Timestamp int64
}

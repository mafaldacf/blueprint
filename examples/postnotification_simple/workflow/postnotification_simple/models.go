package postnotification_simple

type Creator struct {
	Username string
}

type Post struct {
	ReqID     int64
	PostID    int64
	MediaID   int64
	Text      string
	Mentions  []string
	Timestamp int64
	Creator   Creator
}

type Message struct {
	ReqID     string
	PostID    string
	Timestamp string
}

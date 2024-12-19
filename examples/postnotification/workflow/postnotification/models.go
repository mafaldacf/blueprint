package postnotification

type Creator struct {
	Username string //`json:"username"`
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

type Media struct {
	MediaID int64
	Content string
}

type Message struct {
	ReqID     string
	PostID    string
	Timestamp string
}

type Timeline struct {
	ReqID  int64
	PostID int64
}

type Analytics struct {
	PostID int64
}

type TriggerAnalyticsMessage struct {
	PostID string
}

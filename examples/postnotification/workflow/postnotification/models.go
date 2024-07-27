package postnotification

type Creator struct {
	Username string //`json:"username"`
}

type Post struct {
	ReqID     int64  //`json:"reqid"`
	PostID    int64  //`json:"postid"`
	Text      string //`json:"text"`
	Mentions  []string
	Timestamp int64   //`json:"timestamp"`
	Creator   Creator //`json:"creator"`
}

type Message struct {
	ReqID     string //`json:"reqid"`
	PostID    string //`json:"postid"`
	Timestamp string //`json:"timestamp"`
}

type Analytics struct {
	PostID int64
}

type TriggerAnalyticsMessage struct {
	PostID string
}

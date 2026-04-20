package postnotification

type Creator struct {
	Username string
}

type Post struct {
	ReqID     int64    `bson:"ReqID"`
	PostID    int64    `bson:"PostID"`
	MediaID   int64    `bson:"MediaID"`
	Text      string   `bson:"Text"`
	Mentions  []string `bson:"Mentions"`
	Timestamp int64    `bson:"Timestamp"`
	Creator   Creator
}

type Message struct {
	ReqID     int64 `bson:"ReqID"`
	PostID    int64 `bson:"PostID"`
	Timestamp int64 `bson:"Timestamp"`
}

type Analytics struct {
	PostID int64 `bson:"PostID"`
}

type TriggerAnalyticsMessage struct {
	PostID string `bson:"PostID"`
}

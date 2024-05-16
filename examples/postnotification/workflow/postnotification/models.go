package postnotification

type Post struct {
	ReqID 		int64
	PostID 		int64
	Text 		string
	Timestamp 	int64
}

type Message struct {
	ReqID 		string
	PostID 		string
	Timestamp 	string
}

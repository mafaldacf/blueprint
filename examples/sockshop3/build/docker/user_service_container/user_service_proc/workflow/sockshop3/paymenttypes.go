package sockshop3

type Authorisation struct {
	Authorised bool   `json:"authorised"`
	Message    string `json:"message"`
}

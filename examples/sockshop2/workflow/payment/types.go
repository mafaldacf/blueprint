package payment

type Authorisation struct {
	Authorised bool   `json:"authorised"`
	Message    string `json:"message"`
}

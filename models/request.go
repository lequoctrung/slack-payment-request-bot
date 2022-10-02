package models

type Request struct {
	ID          uint    `json:"id"`
	Requester   string  `json:"requester"`
	Receiver    string  `json:"receiver"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
}

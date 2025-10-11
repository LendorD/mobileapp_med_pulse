package models

type Reception struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	// ... другие поля
}

type OneCReceptionsUpdate struct {
	CallID     int         `json:"call_id"`
	Receptions []Reception `json:"receptions"`
}

package models

type Requisition struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

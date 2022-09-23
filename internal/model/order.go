package model

// Order ...
type Order struct {
	OrderUID    string `json:"order_uid"`
	TrackNumber string `json:"track_number,omitempty"`
	Entry       string `json:"entry,omitempty"`
}

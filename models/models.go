package models

// Response for crypto currency
type Response struct {
	Id          string `json:"id,omitempty"`
	FullName    string `json:"fullName,omitempty"`
	Ask         string `json:"ask,omitempty"`
	Bid         string `json:"bid,omitempty"`
	Last        string `json:"last,omitempty"`
	Open        string `json:"open,omitempty"`
	Low         string `json:"low,omitempty"`
	High        string `json:"high,omitempty"`
	FeeCurrency string `json:"feeCurrency,omitempty"`
}

// Response for ErrorResponse to client
type ErrorResponse struct {
	StatusCode int    `json:"statuscode,omitempty"`
	Message    string `json:"message,omitempty"`
}

// Response for External Rest client
type RestResponse struct {
	StatusCode int    `json:"statuscode,omitempty"`
	Message    string `json:"message,omitempty"`
}

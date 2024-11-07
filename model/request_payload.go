package model

// RequestPayload represents the structure of the incoming request body
type RequestPayload struct {
	Name   string `json:"name" validate:"required"`
	Pan    string `json:"pan" validate:"required,pan"`
	Mobile string `json:"mobile" validate:"required,len=10,numeric"`
	Email  string `json:"email" validate:"required,email"`
}

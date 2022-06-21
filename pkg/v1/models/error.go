package models

type ErrorResponse struct {
	Status       string `json:"status,omitempty"`
	StatusCode   int    `json:"status_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

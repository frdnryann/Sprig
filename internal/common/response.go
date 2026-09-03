package common

// common.response
// Standard untuk menetapkan struktur response JSON
// yang dikirim ke client

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

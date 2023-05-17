package responses

type Response struct {
	StatusCode int         `json:"status_code"`
	Payload    interface{} `json:"payload"`
}

type ErrorResponse struct {
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

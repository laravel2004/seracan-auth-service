package response

type ErrorResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func NewErrorResponse(message string, error interface{}) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Error:   error,
	}
}

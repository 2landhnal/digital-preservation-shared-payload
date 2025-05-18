package util

type SuccessResponse struct {
	Message  string `json:"message"`
	Metadata any    `json:"metadata"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewErrorResponse(message string, err error) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Error:   err.Error(),
	}
}

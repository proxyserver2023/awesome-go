package errors

type Err interface {
	Error() string
	StatusCode() int
}

type APIError struct {
	// Status represents the HTTP status code
	Status           int         `json:"status"`
	ErrorCode        string      `json:"error_code"`
	Message          string      `json:"message"`
	DeveloperMessage string      `json:"developer_messsage,omitempty"`
	Details          interface{} `json:"details,omitempty"`
}

func (e APIError) Error() string {
	return e.Message
}

func (e APIError) StatusCode() int {
	return e.Status
}

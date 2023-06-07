package structures

type CustomError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return e.Message
}

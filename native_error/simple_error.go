package native_error

type SimpleError struct {
	Message    string
	StatusCode int
}

func (e *SimpleError) Error() string {
	return e.Message
}

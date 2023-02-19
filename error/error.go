package controllerError

type ControllerError struct {
	StatusCode int
	Message    string
}

func New(statusCode int, message string) ControllerError {
	return ControllerError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e ControllerError) Error() string {
	return e.Message
}

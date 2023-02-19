package controllerError

type ControllerError struct {
	StatusCode int
	Message    string
	Details    any
}

func New(statusCode int, message string, details ...any) ControllerError {
	return ControllerError{
		StatusCode: statusCode,
		Message:    message,
		Details: func() any {
			if len(details) > 0 {
				return details[0]
			}

			return nil
		}(),
	}
}

func (e ControllerError) Error() string {
	return e.Message
}

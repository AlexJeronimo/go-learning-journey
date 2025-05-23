package calculator

type DevideByZeroError struct {
	Message string
}

func (e *DevideByZeroError) Error() string {
	return e.Message
}

var ErrDivideByZero = &DevideByZeroError{Message: "cannot divide by zero"}

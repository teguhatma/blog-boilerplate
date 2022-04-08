package errors

type Code string

type FError interface {
	error
	Code() Code
	Message() string
	Cause() error
}

type baseError struct {
	cause   error
	code    Code
	message string
}

func (b baseError) Error() string {
	str := ""

	if b.code != "" {
		str += string(b.message)
	}

	return str
}

func (b baseError) Message() string {
	return b.message
}

func (b baseError) Code() Code {
	return b.code
}

func (b baseError) Cause() error {
	return b.cause
}

func NewWithCause(code Code, cause error, message string) FError {
	return &baseError{
		cause:   cause,
		code:    code,
		message: message,
	}
}

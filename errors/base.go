package errors

type Code int

type FError interface {
	error
	Code() Code
	Cause() error
}

type baseError struct {
	cause   error
	code    Code
	message string
}

func (b baseError) Code() Code {
	return b.code
}

func (b baseError) Error() string {
	return b.message
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

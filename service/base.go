package service

type Code int

type FError interface {
	error
	Code() Code
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

func NewWithCause(code Code, cause error, message string) FError {
	return &baseError{
		cause:   cause,
		code:    code,
		message: message,
	}
}

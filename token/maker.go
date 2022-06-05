package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for spesific username and duration
	CreateToken(string, time.Duration) (string, error)
	// VerifyToken check if the token valid or not
	VerifyToken(string) (*Payload, error)
}

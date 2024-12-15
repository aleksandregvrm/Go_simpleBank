package util

import "time"

type Maker interface {
	// Creates a token from the users credentials and give it a validity period
	CreateToken(username string, duration time.Duration) (string, error)

	// If token is valid or not
	VerifyToken(token string) (*Payload, error)
}

package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func IsPasswordMatch(candidatePassword, realPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(candidatePassword), []byte(realPassword))
	if err != nil {
		return false
	}
	return true
}

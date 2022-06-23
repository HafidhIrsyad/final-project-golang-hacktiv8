package helper

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPassHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

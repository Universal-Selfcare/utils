package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrMismatchedHashAndPassword = errors.New("credentials are incorrect")
	ErrPasswordTooLong           = errors.New("maximum password length is 72 characters")
)

func HashPassword(password string) (string, error) {
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		return "", ErrPasswordTooLong
	}

	if err != nil {
		return "", err
	}
	return string(hashedPassowrd), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

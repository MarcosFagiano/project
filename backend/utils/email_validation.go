package utils

import "net/mail"

func EmailValidation(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

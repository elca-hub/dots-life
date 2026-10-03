package model

import "net/mail"

type Email struct {
	value *mail.Address
}

func NewEmail(v string) (Email, error) {
	email, err := mail.ParseAddress(v)

	if err != nil {
		return Email{}, err
	}

	return Email{value: email}, nil
}

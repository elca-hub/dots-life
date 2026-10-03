package model

import "time"

type User struct {
	id        string
	name      Name
	email     Email
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(
	id string,
	name Name,
	email Email,
	createdAt time.Time,
	updatedAt time.Time,
) (*User, error) {
	return &User{
		id:        id,
		name:      name,
		email:     email,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (u *User) UpdateName(name Name) {
	u.name = name
}

func (u *User) UpdateEmail(email Email) {
	u.email = email
}

func (u *User) Name() Name {
	return u.name
}

func (u *User) Email() Email {
	return u.email
}

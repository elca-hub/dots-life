package model

import (
	"fmt"
	"unicode/utf8"
)

const (
	MAX_USER_NAME_LEN = 50
)

type Name struct {
	value string
}

func NewName(v string) (Name, error) {
	if utf8.RuneCountInString(v) == 0 || utf8.RuneCountInString(v) > MAX_USER_NAME_LEN {
		return Name{}, fmt.Errorf("Name: %sが条件を満たしていません: %d以上%d以下", v, 1, MAX_USER_NAME_LEN)
	}

	return Name{value: v}, nil
}

func (n Name) String() string {
	return n.value
}

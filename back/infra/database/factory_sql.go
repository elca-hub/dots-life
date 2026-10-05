package database

import (
	"back/domain/repo/db"
	"errors"
)

const (
	InstanceGormMySQL int = iota
)

type SQLInter interface {
	UserRepository() db.InterUserRepository
}

func NewDatabaseSQLFactory(instance int) (SQLInter, error) {
	switch instance {
	case InstanceGormMySQL:
		return NewGormHandler(NewMySQLConfig())
	default:
		return nil, errors.New("invalid instance")
	}
}

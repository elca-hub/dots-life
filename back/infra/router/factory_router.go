package router

import (
	"back/infra/database"
	"errors"
	"time"
)

type Server interface {
	Listen()
}

type Port int64

const (
	InstanceEcho int = iota
)

func NewWebServerFactory(
	instance int,
	port Port,
	ctxTimeout time.Duration,
	rds database.SQLInter,
) (Server, error) {
	switch instance {
	case InstanceEcho:
		// TODO: echoの実装
		return NewEchoEngine(port, ctxTimeout, rds), nil
	default:
		return nil, errors.New("invalid instance")
	}
}

package router

import (
	"back/adapter/logger"
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
	log logger.Logger,
	rds database.SQLInter,
) (Server, error) {
	switch instance {
	case InstanceEcho:
		return NewEchoEngine(port, ctxTimeout, log, rds), nil
	default:
		return nil, errors.New("invalid instance")
	}
}

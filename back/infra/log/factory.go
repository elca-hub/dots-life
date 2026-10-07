package log

import (
	"back/adapter/logger"
	"errors"
)

const (
	InstanceSlog int = iota
)

func NewLoggerFactory(instance int, mode string) (logger.Logger, error) {
	switch instance {
	case InstanceSlog:
		return NewSlogLogger(mode)
	default:
		return nil, errors.New("invalid logger instance")
	}
}

package infra

import (
	"back/adapter/logger"
	"back/infra/database"
	"back/infra/log"
	"back/infra/router"
	"back/mode"
	"fmt"
	"strconv"
	"time"
)

type HttpServerConfig struct {
	appName    string
	mode       string
	ctxTimeout time.Duration
	rds        database.SQLInter
	port       router.Port
	webServer  router.Server
	log        logger.Logger
}

func NewHttpServerConfig() *HttpServerConfig {
	return &HttpServerConfig{}
}

func (h *HttpServerConfig) Name(appName string) *HttpServerConfig {
	h.appName = appName

	return h
}

func (h *HttpServerConfig) Mode(m string) *HttpServerConfig {
	switch m {
	case mode.DevelopmentMode, mode.StagingMode, mode.ProductionMode:
		fmt.Printf("You set the MODE: %s\n", m)
		h.mode = m

		return h
	default:
		panic(fmt.Sprintf("You must set a MODE in your environments: %s, %s, %s", mode.DevelopmentMode, mode.StagingMode, mode.ProductionMode))
	}
}

func (h *HttpServerConfig) CtxTimeout(ctxTimeout time.Duration) *HttpServerConfig {
	h.ctxTimeout = ctxTimeout

	return h
}

func (h *HttpServerConfig) Rds(instance int) *HttpServerConfig {
	rds, err := database.NewDatabaseSQLFactory(instance)

	if err != nil {
		panic(err)
	}

	h.rds = rds

	return h
}

func (h *HttpServerConfig) Logger(instance int) *HttpServerConfig {
	log, err := log.NewLoggerFactory(instance, h.mode)

	if err != nil {
		panic(err)
	}

	h.log = log

	return h
}

func (h *HttpServerConfig) WebServerPort(port string) *HttpServerConfig {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		panic(err) // TODO: loggerの追加
	}
	h.port = router.Port(p)

	return h
}

func (h *HttpServerConfig) WebServer(instance int) *HttpServerConfig {
	s, err := router.NewWebServerFactory(
		instance,
		h.port,
		h.ctxTimeout,
		h.log,
		h.rds,
	)

	if err != nil {
		panic(err)
	}

	h.webServer = s

	return h
}

func (h *HttpServerConfig) Start() {
	h.webServer.Listen()
}

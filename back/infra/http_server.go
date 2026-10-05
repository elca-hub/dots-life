package infra

import (
	"back/infra/database"
	"back/infra/router"
	"strconv"
	"time"
)

type HttpServerConfig struct {
	appName    string
	ctxTimeout time.Duration
	rds        database.SQLInter
	port       router.Port
	webServer  router.Server
}

func NewHttpServerConfig() *HttpServerConfig {
	return &HttpServerConfig{}
}

func (h *HttpServerConfig) Name(appName string) *HttpServerConfig {
	h.appName = appName

	return h
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

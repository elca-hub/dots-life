package router

import (
	"back/adapter/logger"
	"back/infra/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type EchoEngine struct {
	router  *echo.Echo
	port    Port
	timeout time.Duration
	log     logger.Logger
	rds     database.SQLInter
}

func NewEchoEngine(
	port Port,
	timeout time.Duration,
	log logger.Logger,
	rds database.SQLInter,
) *EchoEngine {
	return &EchoEngine{
		echo.New(),
		port,
		timeout,
		log,
		rds,
	}
}

func (e *EchoEngine) Listen() {
	// middleware
	e.router.Use(middleware.Recover())
	e.router.Use(middleware.ContextTimeout(e.timeout))
	e.router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		HandleError: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				e.log.Errorf("request failed", "uri", v.URI, "status", v.Status, "error", v.Error)

				return nil
			}

			e.log.Infof("request", "uri", v.URI, "status", v.Status)

			return nil
		},
	}))

	e.setupRouter(e.router)

	sc := echo.StartConfig{Address: fmt.Sprintf(":%d", e.port)}
	if err := sc.Start(context.Background(), e.router); err != nil {
		e.log.Errorf("サーバが異常停止しました: %+v", err)
	}
}

func (e *EchoEngine) setupRouter(router *echo.Echo) {
	apiRouterGroup := router.Group("/api/v1")
	{
		apiRouterGroup.GET("/ping", e.healthCheckAction)
	}
}

func (e *EchoEngine) healthCheckAction(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]int{
		"status": http.StatusOK,
	})
}

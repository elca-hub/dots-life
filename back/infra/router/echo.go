package router

import (
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
	rds     database.SQLInter
}

func NewEchoEngine(
	port Port,
	timeout time.Duration,
	rds database.SQLInter,
) *EchoEngine {
	return &EchoEngine{
		echo.New(),
		port,
		timeout,
		rds,
	}
}

func (e *EchoEngine) Listen() {
	// middleware
	e.router.Use(middleware.Recover())
	e.router.Use(middleware.ContextTimeout(e.timeout))

	e.setupRouter(e.router)

	sc := echo.StartConfig{Address: fmt.Sprintf(":%d", e.port)}
	if err := sc.Start(context.Background(), e.router); err != nil {
		// TODO: loggerの追加
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

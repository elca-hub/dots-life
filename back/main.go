package main

import (
	"back/infra"
	"back/infra/database"
	"back/infra/log"
	"back/infra/router"
	"os"
	"time"
)

func main() {
	app := infra.NewHttpServerConfig().
		Name(os.Getenv("APP_NAME")).
		Mode(os.Getenv("MODE")).
		CtxTimeout(10 * time.Second).
		Logger(log.InstanceSlog).
		Rds(database.InstanceGormMySQL).
		WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceEcho)

	app.Start()
}

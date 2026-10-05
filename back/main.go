package main

import (
	"back/infra"
	"back/infra/database"
	"back/infra/router"
	"os"
	"time"
)

func main() {
	app := infra.NewHttpServerConfig().
		Name(os.Getenv("APP_NAME")).
		CtxTimeout(10 * time.Second).
		Rds(database.InstanceGormMySQL).
		WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceEcho)

	app.Start()
}

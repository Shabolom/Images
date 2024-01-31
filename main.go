package main

import (
	"Service_Photo/config"
	_ "Service_Photo/docs"
	migrate "Service_Photo/init"
	routes "Service_Photo/iternal/routse"
	tools "Service_Photo/tools/logs"

	log "github.com/sirupsen/logrus"
)

func main() {
	// @title    cmd Service
	// @version  1.0.0
	// @host     localhost:8888
	config.CheckFlagEnv()
	tools.InitLogger()

	err := config.InitPgSQL()
	if err != nil {
		log.WithField("component", "initialization").Panic(err)
	}

	migrate.Migrate()

	r := routes.SetupRouter()

	// запуск сервера
	if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Panic(err)
	}
}

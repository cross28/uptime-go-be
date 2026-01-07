package main

import (
	"context"
	"os"

	"crosssystems.co/uptime-go-be/application"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := application.NewApp()

	log.SetReportCaller(true)

	log.Infof("server starting: port %d | env %s", app.Config.Port, app.Config.Env)
	err := app.Start(context.TODO())

	if err != nil {
		log.Error(err.Error())
	}

	os.Exit(1)
}
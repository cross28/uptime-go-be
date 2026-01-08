package main

import (
	"context"
	"os"
	"os/signal"

	"crosssystems.co/uptime-go-be/application"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Warn("warning: .env file not detected")
	}
	app := application.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Infof("server starting: port %d | env %s", app.Config.Port, app.Config.Env)
	err := app.Start(ctx)
	if err != nil {
		log.Error(err.Error())
	}
}

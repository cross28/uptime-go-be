package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"

	"crosssystems.co/uptime-go-be/application"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(".env"); err != nil {
		slog.Warn("warning: .env file not detected")
	}

	app := application.NewApp()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	slog.Info("server starting: port %d | env %s", strconv.Itoa(app.Config.Port), app.Config.Env)
	err := app.Start(ctx)
	if err != nil {
		slog.Error(err.Error())
	}
}

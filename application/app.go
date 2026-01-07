package application

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Configuration struct {
	Port int
	Env string
}

type App struct {
	Config Configuration
	Router *chi.Mux
}

func NewApp() *App {
	var cfg Configuration

	flag.IntVar(&cfg.Port, "port", 8000, "api server port")
	flag.StringVar(&cfg.Env, "env", "dev", "server environment(development|qa|prod)")
	flag.Parse()
	
	return &App{
		Config: cfg,
		Router: RegisterRoutes(),
	}
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", a.Config.Port),
		Handler: a.Router,
		IdleTimeout: time.Minute,
		WriteTimeout: time.Second * 10,
		ReadTimeout: time.Second * 5,
	}

	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
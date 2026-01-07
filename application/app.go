package application

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/redis/go-redis/v9"
)

type Configuration struct {
	Port int
	Env string
}

type App struct {
	Config Configuration
	Router *chi.Mux
	RedisClient *redis.Client
}

func NewApp() *App {
	var cfg Configuration

	flag.IntVar(&cfg.Port, "port", 8000, "api server port")
	flag.StringVar(&cfg.Env, "env", "dev", "server environment(development|qa|prod)")
	flag.Parse()
	
	return &App{
		Config: cfg,
		Router: RegisterRoutes(),
		RedisClient: redis.NewClient(&redis.Options{}),
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

	err := a.RedisClient.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to start redis: %w", err)
	}

	defer func() {
		if err := a.RedisClient.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()
	
	ctx.Done()
	err = <-ch

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second * 10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
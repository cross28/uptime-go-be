package application

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Configuration struct {
	Port int
	Env  string
}

type App struct {
	Config      Configuration
	Router      *chi.Mux
	RedisClient *redis.Client
	PostgresDb  *pgxpool.Pool
}

func NewApp() *App {
	var cfg Configuration

	flag.IntVar(&cfg.Port, "port", 8000, "api server port")
	flag.StringVar(&cfg.Env, "env", "dev", "server environment(development|qa|prod)")
	flag.Parse()

	redisClient := NewRedisClient(RedisConfig{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	postgresPool, err := NewPostgresConnection(context.Background(), PostgresConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
	})
	if err != nil {
		os.Exit(1)
	}

	return &App{
		Config:      cfg,
		RedisClient: redisClient,
		PostgresDb:  postgresPool,
	}
}

func (a *App) Start(ctx context.Context) error {
	a.RegisterRoutes()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.Config.Port),
		Handler:      a.Router,
		IdleTimeout:  time.Minute,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
	}

	if err := a.RedisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to start redis: %w", err)
	}

	if err := a.PostgresDb.Ping(ctx); err != nil {
		return fmt.Errorf("failed to start postgres: %w", err)
	}

	defer func() {
		if err := a.RedisClient.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}

		a.PostgresDb.Close()
	}()

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}

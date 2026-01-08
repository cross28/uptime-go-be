package application

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Address  string
	Port     string
	Password string
}

func NewRedisClient(cfg RedisConfig) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
			Password: cfg.Password,
			// TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12}, // Needed for production
		},
	)
}

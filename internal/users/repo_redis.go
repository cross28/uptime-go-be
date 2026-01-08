package users

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisUserRepository struct {
	redis *redis.Client
}

func NewRedisUserRepository(redis *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{
		redis: redis,
	}
}

func UserIdKey(id string) string {
	return fmt.Sprintf("user:%s", id)
}

var ErrNotExist = errors.New("order does not exist")

func (r *RedisUserRepository) GetById(ctx context.Context, id string) (User, error) {
	key := UserIdKey(id)

	value, err := r.redis.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return User{}, ErrNotExist
	} else if err != nil {
		return User{}, fmt.Errorf("get order: %w", err)
	}

	var user User
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return User{}, fmt.Errorf("%w", err)
	}

	return user, nil
}

func (r *RedisUserRepository) Create(ctx context.Context, user *User) (string, error) {
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("failed to encode order: %w", err)
	}

	randomId := rand.Text()[:16]
	user.Id = randomId

	cmd := r.redis.SetNX(ctx, UserIdKey(user.Id), userJson, 0)
	return user.Id, cmd.Err()
}

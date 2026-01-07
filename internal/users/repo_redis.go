package users

import "github.com/redis/go-redis/v9"

type RedisUserRepository struct{
	redis *redis.Client
}

func NewRedisUserRepository(redis *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{redis: redis}
}

func (r *RedisUserRepository) GetById(id string) (*User, error) {
	return nil, nil
}

func (r *RedisUserRepository) Create(user *User) (string, error) {
	return "", nil
}
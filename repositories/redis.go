package repositories

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

type Redis struct {
	*redis.Client
}

// -------------------- Connection Logic ---------------------

func NewRedisClient(ctx context.Context) (*Redis, error) {
	connectOptions := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	client := redis.NewClient(&connectOptions)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{client}, nil

}

func (r *Redis) Close() {
	r.Client.Close()
}

// -------------------- Access Database ---------------------

func (r *Redis) AddToRedis(ctx context.Context, key string, value any, expiresAt time.Duration) error {
	err := r.Client.Set(ctx, key, value, expiresAt).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) GetFromRedis(ctx context.Context, key string) (string, error) {
	result, err := r.Client.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		fmt.Printf("%s does not exist within the cache pool", key)
		return "", err
	case err != nil:
		fmt.Println("Something went wrong with the redis server")
		return "", err
	default:
		return result, err
	}
}

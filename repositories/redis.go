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

func AddToRedis(ctx context.Context, c *redis.Client, key string, value any, expiresAt time.Duration) error {
	err := c.Set(ctx, key, value, expiresAt).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetFromRedis(ctx context.Context, c *redis.Client, key string) (string, error) {
	result, err := c.Get(ctx, key).Result()
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

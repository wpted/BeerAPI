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

// NewRedisClient takes input context and return a pointer to Redis struct.
// If connection failed, return nil and the given error
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
	fmt.Println("You've successfully connected to Redis.")
	return &Redis{client}, nil

}

// Close shuts the redis connection down
func (r *Redis) Close() {
	r.Client.Close()
}

// -------------------- Access Database ---------------------

// AddToRedis sets the given Key-Value pair within Redis, and return an empty error if succeeded.
// If process failed, return the given error.
func (r *Redis) AddToRedis(ctx context.Context, key string, value any, expiresAt time.Duration) error {
	err := r.Client.Set(ctx, key, value, expiresAt).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetFromRedis gets the Value from the given Key within Redis, returns the values string and an empty error if succeeded.
// If process failed, return an empty string and the given error.
func (r *Redis) GetFromRedis(ctx context.Context, key string) (string, error) {
	result, err := r.Client.Get(ctx, key).Result()
	switch {
	case err == redis.Nil:
		fmt.Printf("%s does not exist within the cache pool\n", key)
		return "", err
	case err != nil:
		fmt.Println("Something went wrong with the redis server")
		return "", err
	default:
		return result, err
	}
}

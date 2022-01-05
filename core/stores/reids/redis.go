package reids

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func DefaultRedis(ctx context.Context, addr string, pwd ...string) *redis.Client {
	password := ""
	if len(pwd) >= 1 {
		password = pwd[0]
	}
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}

func InitRedis(ctx context.Context, addr, pwd string, DB int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       DB,  // use default DB
	})
}

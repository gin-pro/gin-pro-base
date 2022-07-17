package cache

import (
	"github.com/go-redis/redis/v8"
)

func DefaultRedis(addr string, pwd ...string) *redis.Client {
	password := ""
	if len(pwd) >= 1 {
		password = pwd[0]
	}
	return InitRedis(addr, password, 0)
}

func InitRedis(addr, pwd string, DB int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       DB,  // use default DB
	})
}

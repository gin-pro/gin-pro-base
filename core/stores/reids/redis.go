package reids

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisC struct {
	cli *redis.Client
	ctx context.Context
}

func DefaultRedis(ctx context.Context, addr string, pwd ...string) *RedisC {
	password := ""
	if len(pwd) >= 1 && pwd[0] != "" {
		password = pwd[0]
	}
	return &RedisC{
		cli: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password, // no password set
			DB:       0,        // use default DB
		}),
		ctx: ctx,
	}
}

func InitRedis(ctx context.Context, addr, pwd string, DB int) *RedisC {
	return &RedisC{
		cli: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pwd, // no password set
			DB:       DB,  // use default DB
		}),
		ctx: ctx,
	}
}

func (c *RedisC) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.cli.Set(c.ctx, key, value, expiration)
}

func (c *RedisC) Get(key string) *redis.StringCmd {
	return c.cli.Get(c.ctx, key)
}

func (c *RedisC) SetEx(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.cli.SetEX(c.ctx, key, value, expiration)
}

func (c *RedisC) Del(key string) *redis.IntCmd {
	return c.cli.Del(c.ctx, key)
}

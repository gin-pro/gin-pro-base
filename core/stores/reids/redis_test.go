package reids

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestRedis(t *testing.T) {
	var ctx = context.Background()
	rdb := DefaultRedis(ctx, "localhost:6379")
	err := rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	val2, err := rdb.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

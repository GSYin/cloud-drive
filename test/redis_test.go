package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", time.Second*10).Err()
	if err != nil {
		t.Error(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

}

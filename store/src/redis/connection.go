package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

var lock = &sync.Mutex{}

type redisConnection struct {
	rdb *redis.Client
}

func (rc redisConnection) connect(options redis.Options) {
	rc.rdb = redis.NewClient(&options)

	ctx := context.TODO()
	_, err := rc.rdb.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}
}

func (rc redisConnection) connectByUrl(url string) {

	options, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	rc.rdb = redis.NewClient(options)

	ctx := context.TODO()
	_, pingErr := rc.rdb.Ping(ctx).Result()

	if pingErr != nil {
		panic(pingErr)
	}
}

var connection *redisConnection

func RedisConnection(options redis.Options) *redisConnection {

	if connection == nil {
		lock.Lock()
		defer lock.Unlock()
		if connection == nil {
			fmt.Println("Creating single instance now.")
			connection = &redisConnection{}
			connection.connect(options)
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return connection
}

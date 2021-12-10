package store

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
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

func (rc *redisConnection) connectByUrl(url string) {

	options, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	rc.rdb = redis.NewClient(options)

	ctx := context.TODO()
	resp, pingErr := rc.rdb.Ping(ctx).Result()

	log.Info("Redis ping responded with ", resp)

	if pingErr != nil {
		panic(pingErr)
	}
}

var connection *redisConnection

func RedisConnection(url string) *redisConnection {

	if connection == nil {
		lock.Lock()
		defer lock.Unlock()
		if connection == nil {
			log.Info("Creating redis instance")
			connection = &redisConnection{}
			connection.connectByUrl(url)
		}
	}

	return connection
}

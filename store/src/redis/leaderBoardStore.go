package store

import (
	"context"
	"fmt"
	interfaces "liquide/re/popularity-leaderboard-builder/store/interfaces"

	log "github.com/sirupsen/logrus"
)

type RedisLeaderboardStore struct {
	redis redisConnection
	interfaces.LeaderboardStore
}

func (s *RedisLeaderboardStore) IncrementScoreForAnItem(boardName string, item string, points float64) {

	ctx := context.TODO()

	// newScore, err := s.redis.rdb.Do(ctx, "ZINCRBY", boardName, points, item).Result()
	newScore, err := s.redis.rdb.ZIncrBy(ctx, boardName, points, item).Result()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Info(fmt.Sprintf("[%s] %s - %f", boardName, item, newScore))
}

func GetRedisLeaderboardStore(redisConnection *redisConnection) *RedisLeaderboardStore {

	return &RedisLeaderboardStore{
		redis: *redisConnection,
	}
}

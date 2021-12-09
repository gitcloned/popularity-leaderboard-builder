package store

import (
	"context"
	"fmt"
	interfaces "liquide/re/popularity-leaderboard-builder/store/interfaces"
)

type RedisLeaderboardStore struct {
	redis redisConnection
	interfaces.LeaderboardStore
}

func (s RedisLeaderboardStore) IncrementScoreForAnItem(boardName string, item string, points float64) {

	ctx := context.TODO()

	newScore, err := s.redis.rdb.Do(ctx, "ZINCRBY", boardName, points, item).Result()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(fmt.Sprintf("[%s] %s - %d", boardName, item, newScore))
}

package store

import (
	interfaces "liquide/re/popularity-leaderboard-builder/store/interfaces"
	redis "liquide/re/popularity-leaderboard-builder/store/src/redis"
)

func LeaderboardStoreProvider() (interfaces.LeaderboardStore, error) {

	// store := memory.InMemoryLeaderboardStore{
	// 	Board: make(map[string]float64),
	// }

	conn := redis.RedisConnection("redis://localhost:6379/0")
	return redis.GetRedisLeaderboardStore(conn), nil
}

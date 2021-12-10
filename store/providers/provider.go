package store

import (
	interfaces "liquide/re/popularity-leaderboard-builder/store/interfaces"
	memory "liquide/re/popularity-leaderboard-builder/store/src/memory"
)

func LeaderboardStoreProvider() (interfaces.LeaderboardStore, error) {

	store := memory.InMemoryLeaderboardStore{
		Board: make(map[string]float64),
	}

	return &store, nil
}

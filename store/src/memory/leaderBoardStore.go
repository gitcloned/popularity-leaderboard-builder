package store

import (
	"fmt"
	interfaces "liquide/re/popularity-leaderboard-builder/store/interfaces"
)

type InMemoryLeaderboardStore struct {
	interfaces.LeaderboardStore
}

func (s InMemoryLeaderboardStore) IncrementScoreForAnItem(boardName string, item string, points float64) {

	fmt.Println(fmt.Sprintf("[%s] %s - %d", boardName, item, points))
}

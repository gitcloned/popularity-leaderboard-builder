package store

import (
	"fmt"
	interfaces "liquide/re/popularity-leaderboard-builder/store/interfaces"

	log "github.com/sirupsen/logrus"
)

type InMemoryLeaderboardStore struct {
	Board map[string]float64

	interfaces.LeaderboardStore
}

func (s *InMemoryLeaderboardStore) IncrementScoreForAnItem(boardName string, item string, points float64) {

	i := s.Board[item]

	if i == 0 {
		s.Board[item] = points
	} else {
		s.Board[item] += points
	}

	log.Info(fmt.Sprintf("[%s] %s - %f", boardName, item, s.Board[item]))
}

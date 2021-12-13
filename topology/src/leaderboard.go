package topology

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	store "liquide/re/popularity-leaderboard-builder/store/interfaces"
)

type Leaderboard struct {
	Store store.LeaderboardStore
}

func (l *Leaderboard) RankItem(path string, itemName string, score float64, u *objects.UserAction) {

	// increment item score by the points earned through action
	l.Store.IncrementScoreForAnItem(path, itemName, score)
}

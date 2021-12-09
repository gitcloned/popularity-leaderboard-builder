package topology

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	store "liquide/re/popularity-leaderboard-builder/store/interfaces"
)

type Leaderboard struct {
	Name  string
	store store.LeaderboardStore
}

func (l Leaderboard) ProcessAction(u *objects.UserAction) {

	// increment item score by the points earned through action
	l.store.IncrementScoreForAnItem(l.Name, u.Item, u.Points)
}

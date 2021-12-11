package topology

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	store "liquide/re/popularity-leaderboard-builder/store/interfaces"
)

type Leaderboard struct {
	Name  string
	Store store.LeaderboardStore
}

func (l *Leaderboard) ProcessAction(u *objects.UserAction) {

	// increment item score by the points earned through action
	l.Store.IncrementScoreForAnItem(l.Name, u.Item.Id, u.Points)
}

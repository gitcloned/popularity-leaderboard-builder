package topology

import (
	"liquide/re/popularity-leaderboard-builder/objects"
)

type Rule interface {
	MatchUserAction(u *objects.UserAction) bool
}

package reader

import "liquide/re/popularity-leaderboard-builder/objects"

type UserActionReader interface {
	Read(chan objects.UserAction)
}

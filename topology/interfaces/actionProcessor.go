package topology

import "liquide/re/popularity-leaderboard-builder/objects"

type ActionProcessor interface {
	ProcessAction(*objects.UserAction)
}

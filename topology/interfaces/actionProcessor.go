package topology

import "liquide/re/popularity-leaderboard-builder/objects"

type ActionProcessor interface {
	ProcessAction(*objects.UserAction)
}
type ItemRanker interface {
	RankItem(path string, itemName string, score float64, userAction *objects.UserAction)
}

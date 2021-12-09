package topology

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
)

type Tree struct {
	root Node

	interfaces.ActionProcessor
}

func (t Tree) ProcessAction(u objects.UserAction) {
	t.root.ProcessAction(u)
}

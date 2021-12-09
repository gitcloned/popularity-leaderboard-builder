package topology

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
)

type Tree struct {
	Branches []Branch
	World    Node

	interfaces.ActionProcessor
}

func (t Tree) ProcessAction(u *objects.UserAction) {

	for _, branch := range t.Branches {

		branch.ProcessAction(u)
	}
}

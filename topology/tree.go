package topology

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	"sync"
)

var lock = &sync.Mutex{}

type Tree struct {
	Branches []Branch
	World    Node

	interfaces.ActionProcessor
}

func (t *Tree) ProcessAction(u *objects.UserAction) {

	for idx, _ := range t.Branches {

		t.Branches[idx].ProcessAction(u)
	}
}

package topology

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	src "liquide/re/popularity-leaderboard-builder/topology/src"
)

type Node struct {
	name        string
	leaderboard src.Leaderboard
	predicate   src.Predicate

	interfaces.ActionProcessor
}

func (n Node) ProcessAction(u *objects.UserAction) {

	n.leaderboard.ProcessAction(u)
}

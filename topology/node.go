package topology

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	src "liquide/re/popularity-leaderboard-builder/topology/src"
)

type Node struct {
	name        string
	leaderboard src.Leaderboard
	children    []Node
	predicate   src.Predicate

	interfaces.ActionProcessor
}

func (n Node) ProcessAction(u objects.UserAction) {

	if n.predicate.MatchUserAction(u) {
		n.leaderboard.ProcessAction(u)
	}

	for _, child := range n.children {
		child.ProcessAction(u)
	}
}

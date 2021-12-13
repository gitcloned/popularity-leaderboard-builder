package topology

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
)

type Garden struct {
	Trees []Tree

	interfaces.ActionProcessor
}

func (g *Garden) ProcessAction(u *objects.UserAction) {

	for idx, _ := range g.Trees {

		g.Trees[idx].ProcessAction(u)
	}
}

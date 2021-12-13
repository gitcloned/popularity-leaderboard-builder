package topology

import (
	"fmt"
	objects "liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	src "liquide/re/popularity-leaderboard-builder/topology/src"
)

type Node struct {
	Name        string // field name : field value
	leaderboard src.Leaderboard
	predicate   src.Predicate

	branches []Branch // any sub branch from this node

	interfaces.ItemRanker
}

func (n *Node) RankItem(path string, itemName string, score float64, u *objects.UserAction) {

	nodePath := fmt.Sprintf("%s%s", path, n.Name)

	n.leaderboard.RankItem(nodePath, itemName, score, u)

	for idx, _ := range n.branches {

		n.branches[idx].RankItem(fmt.Sprintf("%s%s", nodePath, "-"), itemName, score, u)
	}
}

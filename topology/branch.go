package topology

import (
	objects "liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	topology "liquide/re/popularity-leaderboard-builder/topology/src"
	"reflect"
)

type Branch struct {
	Path  string // path of this branch
	Field string // user action field to refer for the branch value

	Branches []Branch        // sub branches
	nodes    map[string]Node // dub nodes for distinct value of the branch field received

	interfaces.ActionProcessor
}

func (b Branch) node(u *objects.UserAction) Node {

	// value from user action for this branch
	// TODO: Handle if field does not exists, or value is not string or nil
	branchValue := reflect.ValueOf(*u).FieldByName(b.Field).String()

	node, exists := b.nodes[branchValue]

	if exists {
		return node
	} else {
		node = Node{
			name: b.Path + "." + branchValue,
			leaderboard: topology.Leaderboard{
				Name: b.Path + "." + branchValue,
			},
		}
		b.nodes[branchValue] = node
		return node
	}
}

func (b Branch) ProcessAction(u *objects.UserAction) {

	b.node(u).ProcessAction(u)

	for _, branch := range b.Branches {

		branch.ProcessAction(u)
	}
}

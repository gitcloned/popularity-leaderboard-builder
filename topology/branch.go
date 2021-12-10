package topology

import (
	"fmt"
	objects "liquide/re/popularity-leaderboard-builder/objects"
	store "liquide/re/popularity-leaderboard-builder/store/providers"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	topology "liquide/re/popularity-leaderboard-builder/topology/src"
	"reflect"

	log "github.com/sirupsen/logrus"
)

type Branch struct {
	Path  string // path of this branch
	Field string // user action field to refer for the branch value

	Branches []Branch        // sub branches
	nodes    map[string]Node // dub nodes for distinct value of the branch field received

	initialized bool

	interfaces.ActionProcessor
}

func (b *Branch) init() {

	if !b.initialized {

		// log.Info("Initializing branch: ", b.Path)

		b.nodes = make(map[string]Node)
		b.initialized = true
	}
}

func (b *Branch) node(u *objects.UserAction) *Node {

	// value from user action for this branch
	// TODO: Handle if field does not exists, or value is not string or nil
	branchValue := reflect.ValueOf(*u).FieldByName(b.Field).String()

	node, exists := b.nodes[branchValue]

	if exists {
		return &node
	} else {
		lbStore, err := store.LeaderboardStoreProvider()

		if err != nil {
			panic(err)
		}

		node = Node{
			name: b.Path + "." + branchValue,
			leaderboard: topology.Leaderboard{
				Name:  b.Path + "." + branchValue,
				Store: lbStore,
			},
		}

		log.Info(fmt.Sprintf("Creating branch at path '%s' with name '%s'", b.Path, branchValue))
		b.nodes[branchValue] = node
		return &node
	}
}

func (b *Branch) ProcessAction(u *objects.UserAction) {

	b.init()

	b.node(u).ProcessAction(u)

	for idx, _ := range b.Branches {

		b.Branches[idx].ProcessAction(u)
	}
}

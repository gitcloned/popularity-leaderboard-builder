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
	Name  string // path of this branch
	Field string // user action field to refer for the branch value

	Branches []Branch        // sub branches
	nodes    map[string]Node // dub nodes for distinct value of the branch field received

	initialized bool

	interfaces.ItemRanker
}

func (b *Branch) init() {

	if !b.initialized {

		// log.Info("Initializing branch: ", b.Path)

		b.nodes = make(map[string]Node)
		b.initialized = true
	}
}

func (b *Branch) node(u *objects.UserAction, branches []Branch) *Node {

	// value from user action for this branch
	// TODO: Handle if field does not exists, or value is not string or nil
	branchValue := reflect.ValueOf(*u).FieldByName(b.Field).String()

	path := fmt.Sprintf("%s:%s", b.Name, branchValue)

	node, exists := b.nodes[path]

	if exists {
		return &node
	} else {
		lbStore, err := store.LeaderboardStoreProvider()

		if err != nil {
			panic(err)
		}

		node = Node{
			Name: path,
			leaderboard: topology.Leaderboard{
				Store: lbStore,
			},
			branches: branches,
		}

		log.Info(fmt.Sprintf("Creating branch node '%s'", path))
		b.nodes[path] = node
		return &node
	}
}

func (b *Branch) RankItem(path string, itemName string, score float64, u *objects.UserAction) {

	b.init()

	b.node(u, b.Branches).RankItem(path, itemName, score, u)
}

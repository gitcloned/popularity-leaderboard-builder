package topology

import (
	"fmt"
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	"reflect"
	"sync"
)

var lock = &sync.Mutex{}

type Tree struct {
	Name          string
	Branches      []Branch
	ItemFieldName string

	interfaces.ActionProcessor
}

func (t *Tree) ProcessAction(u *objects.UserAction) {

	// extract the field name to rank by this tree
	fieldName := reflect.ValueOf(u.Item).FieldByName(t.ItemFieldName).String()
	value := u.Points

	// pass to each of the branch for this tree
	for idx, _ := range t.Branches {

		t.Branches[idx].RankItem(fmt.Sprintf("%s-", t.Name), fieldName, value, u)
	}
}

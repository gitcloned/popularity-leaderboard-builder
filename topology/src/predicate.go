package topology

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
)

type Predicate struct {
	rules []interfaces.Rule
}

func (p Predicate) MatchUserAction(u *objects.UserAction) bool {

	matched := true

	for _, rule := range p.rules {

		matched = matched && rule.MatchUserAction(u)

		if !matched {
			break
		}
	}

	return matched
}

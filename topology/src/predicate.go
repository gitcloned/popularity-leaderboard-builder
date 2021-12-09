package topology

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	rules "liquide/re/popularity-leaderboard-builder/topology/src/rules"
)

type Predicate struct {
	rules []rules.Rule
}

func (p Predicate) MatchUserAction(u objects.UserAction) bool {

	matched := true

	for _, rule := range p.rules {

		matched = matched && rule.MatchUserAction(u)

		if !matched {
			break
		}
	}

	return matched
}

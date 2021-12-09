package topology

import (
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	"reflect"
	"strings"
)

type StringOperator interface {
	interfaces.Operator
}

type StringEqualToOperator struct {
	StringOperator
}

type StringContainsOperator struct {
	StringOperator
}

func (o StringEqualToOperator) Match(lhs reflect.Value, rhs StringOperandValue) bool {
	return lhs.String() == rhs.value
}

func (o StringContainsOperator) Match(lhs reflect.Value, rhs StringOperandValue) bool {
	return strings.Contains(lhs.String(), rhs.value)
}

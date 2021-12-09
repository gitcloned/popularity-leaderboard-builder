package topology

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	"reflect"
)

type IntOperandValue struct {
	value int
	interfaces.OperandValue
}

func (v IntOperandValue) Value() int {
	return v.value
}

type IntValueRule struct {
	fieldName string
	operator  IntOperator
	value     int
	negate    bool

	interfaces.Rule
}

func (r IntValueRule) MatchUserAction(u objects.UserAction) bool {

	return r.operator.Match(reflect.ValueOf(&u), IntOperandValue{
		value: r.value,
	})
}

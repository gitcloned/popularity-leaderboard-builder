package topology

import (
	"liquide/re/popularity-leaderboard-builder/objects"
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	"reflect"
)

type StringOperandValue struct {
	value string
	interfaces.OperandValue
}

func (v StringOperandValue) Value() string {
	return v.value
}

type StringValueRule struct {
	fieldName string
	operator  StringOperator
	value     string
	negate    bool

	interfaces.Rule
}

func (r StringValueRule) MatchUserAction(u objects.UserAction) bool {

	return r.operator.Match(reflect.ValueOf(&u), StringOperandValue{
		value: r.value,
	})
}

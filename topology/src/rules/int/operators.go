package topology

import (
	interfaces "liquide/re/popularity-leaderboard-builder/topology/interfaces"
	"reflect"
)

type IntOperator interface {
	interfaces.Operator
}

type IntEqualToOperator struct {
	IntOperator
}

type IntGreaterThanOperator struct {
	IntOperator
}

type IntLessThanOperator struct {
	IntOperator
}

type IntGreaterThanEqualToOperator struct {
	IntOperator
}

type IntLessThanEqualToOperator struct {
	IntOperator
}

func (o IntEqualToOperator) Match(lhs reflect.Value, rhs IntOperandValue) bool {
	return int(lhs.Int()) == rhs.value
}

func (o IntGreaterThanOperator) Match(lhs reflect.Value, rhs IntOperandValue) bool {
	return int(lhs.Int()) > rhs.value
}

func (o IntGreaterThanEqualToOperator) Match(lhs reflect.Value, rhs IntOperandValue) bool {
	return int(lhs.Int()) >= rhs.value
}

func (o IntLessThanOperator) Match(lhs reflect.Value, rhs IntOperandValue) bool {
	return int(lhs.Int()) < rhs.value
}

func (o IntLessThanEqualToOperator) Match(lhs reflect.Value, rhs IntOperandValue) bool {
	return int(lhs.Int()) <= rhs.value
}

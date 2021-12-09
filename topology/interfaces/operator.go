package topology

import "reflect"

type OperandValue interface {
}

type Operator interface {
	Match(lhs reflect.Value, rhs OperandValue) bool
}

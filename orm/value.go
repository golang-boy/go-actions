package orm

import "fmt"

type Value struct {
	value any
}

func NewValue(value any) *Value {
	return &Value{value: value}
}

func (v *Value) expr() {}

func (v *Value) String() string {
	return fmt.Sprintf("%v", v.value)
}

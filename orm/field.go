package orm

type Field struct {
	name string
}

func (f *Field) String() string {
	return f.name
}

// F("id").Eq(1).And(F("name").Like("abc%"))
func F(name string) *Field {
	return &Field{name: name}
}

func (f *Field) Eq(value any) *Condition {
	return &Condition{
		left:  f,
		op:    OpEq,
		right: NewValue(value),
	}
}

func (f *Field) expr() {}

func (f *Field) GT(value any) *Condition {
	return &Condition{
		left:  f,
		op:    OpGt,
		right: NewValue(value),
	}
}

func (f *Field) LT(value any) *Condition {
	return &Condition{
		left:  f,
		op:    OpLt,
		right: NewValue(value),
	}
}

package orm

type Operator string

const (
	OpEq    Operator = "="
	OpNe    Operator = "!="
	OpGt    Operator = ">"
	OpLt    Operator = "<"
	OpGe    Operator = ">="
	OpLe    Operator = "<="
	OpIn    Operator = "in"
	OpNotIn Operator = "not in"
	OpLike  Operator = "like"
	// 其他
	OpNOT = "NOT"
	OpAND = "AND"
	OpOR  = "OR"
)

func (o Operator) String() string {
	return string(o)
}

type Condition struct {
	left  Expression
	op    Operator
	right Expression
}

func (c *Condition) expr() {}

func (c *Condition) And(right *Condition) *Condition {
	return &Condition{
		left:  c,
		op:    OpAND,
		right: right,
	}
}

func (c *Condition) Or(right *Condition) *Condition {
	return &Condition{
		left:  c,
		op:    OpOR,
		right: right,
	}
}

func Not(c *Condition) *Condition {
	return &Condition{
		left:  nil,
		op:    OpNOT,
		right: c,
	}
}

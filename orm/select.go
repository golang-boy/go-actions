package orm

import (
	"context"
	"database/sql"
	"errors"
	"go-actions/orm/internal/errs"
	"go-actions/orm/meta"
	"strings"
)

var ErrInvalidExpression = errors.New("invalid expression")

type Selector[T any] struct {
	sb    *strings.Builder
	table string
	where []*Condition
	args  []any

	model *meta.Model
	db    *DB
}

func NewSelector[T any](db *DB) *Selector[T] {
	return &Selector[T]{
		db: db,
		sb: new(strings.Builder),
	}
}

func (s *Selector[T]) Build() (*Query, error) {
	var (
		t   T
		err error
	)

	//
	s.model, err = s.db.r.Get(&t)
	if err != nil {
		return nil, err
	}

	s.sb.WriteString("SELECT * FROM ")

	if s.table == "" {
		s.sb.WriteByte('`')
		s.sb.WriteString(s.model.TableName)
		s.sb.WriteByte('`')
	} else {
		s.sb.WriteString(s.table)
	}

	if len(s.where) > 0 {
		s.sb.WriteString(" WHERE ")

		// 组合
		p := s.where[0]
		for i := 1; i < len(s.where); i++ {
			p = p.And(s.where[i])
		}

		err := s.buildExpression(p)
		if err != nil {
			return nil, err
		}
	}
	s.sb.WriteByte(';')
	return &Query{
		SQL:  s.sb.String(),
		Args: s.args,
	}, nil
}

func (s *Selector[T]) buildExpression(e Expression) error {
	if e == nil {
		return nil
	}
	switch expr := e.(type) {
	case *Condition:
		_, leftok := expr.left.(*Condition)
		if leftok {
			s.sb.WriteByte('(')
		}
		err := s.buildExpression(expr.left)
		if err != nil {
			return err
		}
		if leftok {
			s.sb.WriteByte(')')
		}
		s.sb.WriteByte(' ')
		s.sb.WriteString(expr.op.String())
		s.sb.WriteByte(' ')

		_, rightok := expr.right.(*Condition)
		if rightok {
			s.sb.WriteByte('(')
		}
		err = s.buildExpression(expr.right)
		if err != nil {
			return err
		}
		if rightok {
			s.sb.WriteByte(')')
		}
	case *Field:
		fd, ok := s.model.FieldMap[expr.name]
		if !ok {
			return errs.NewErrUnknownField(expr.name)
		}
		s.sb.WriteByte('`')
		s.sb.WriteString(fd.ColName)
		s.sb.WriteByte('`')
	case *Value:
		s.sb.WriteByte('?')
		s.addArg(expr.value)
	default:
		return ErrInvalidExpression
	}
	return nil
}

func (s *Selector[T]) addArg(val any) {
	if s.args == nil {
		s.args = make([]any, 0, 4)
	}
	s.args = append(s.args, val)
}

func (s *Selector[T]) Where(where ...*Condition) *Selector[T] {
	s.where = where
	return s
}

func (s *Selector[T]) From(table string) *Selector[T] {
	s.table = table
	return s
}

func (s *Selector[T]) Get(ctx context.Context) (*T, error) {
	var db *sql.DB = s.db.db
	q, err := s.Build()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, q.SQL, q.Args...)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, ErrNoRows
	}

	tp := new(T)
	val := s.db.Creator(s.model, tp)
	err = val.SetColumns(rows)

	return tp, err
}

// func (s *Selector[T]) List(ctx context.Context) ([]*T, error) {
// 	var db *sql.DB = s.db.db
// 	q, err := s.Build()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// rows, err := db.QueryContext(ctx, q.SQL, q.Args...)

// 	return nil, nil
// }

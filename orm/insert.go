package orm

import (
	"context"
	"database/sql"
)

type Inserter[T any] struct {
	db *sql.DB
}

func NewInserter[T any]() *Inserter[T] {
	return &Inserter[T]{}
}

func (i *Inserter[T]) Build() (*Query, error) {
	return nil, nil
}

func (i *Inserter[T]) Values(vals ...*T) *Inserter[T] {
	return nil
}

func (i *Inserter[T]) Exec(ctx context.Context) (sql.Result, error) {
	q, err := i.Build()
	if err != nil {
		return nil, err
	}
	return i.db.ExecContext(ctx, q.SQL, q.Args...)
}
